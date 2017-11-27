package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dlespiau/quay-exporter/pkg/models"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var quayVulnerabilities = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "quay_vulnerabilities",
	Help: "Number of vulnerabilities reported by quay.io.",
}, []string{"organization", "repository", "os", "severity"})

func init() {
	prometheus.MustRegister(quayVulnerabilities)
}

type operation int

const (
	opScan operation = iota
)

type command struct {
	op operation
}

type severity string

const (
	severityUnknown    severity = "unknown"
	severityNegligible          = "negligible" // Debian calls this "unimportant"
	severityLow                 = "low"
	severityMedium              = "medium"
	severityHigh                = "high"
	severityCritical            = "critical"
)

var strToSeverity = map[string]severity{
	"unknown":    severityUnknown,
	"negligible": severityNegligible,
	"low":        severityLow,
	"medium":     severityMedium,
	"critical":   severityCritical,
}

func severityFromString(input string) severity {
	input = strings.ToLower(input)
	if severity, ok := strToSeverity[input]; ok {
		return severity
	}
	return severityUnknown
}

// Convert a Clair severity to a severity.
// For reference: https://github.com/coreos/clair/blob/v2.0.1/database/severity.go#L31
func severityFromClairSeverity(input string) severity {
	if input == "Defcon1" {
		input = "Critical"
	}
	return severityFromString(input)
}

// See https://www.first.org/cvss/specification-document table 14.
// Parsing the CVSSv2 attack vectors may be a better way to derive the severity
// of the vulnerability!
func severityFromCVSSv2Score(score float64) severity {
	if score >= 9.0 {
		return severityCritical
	}
	if score >= 7.0 {
		return severityHigh
	}
	if score >= 4.0 {
		return severityMedium
	}
	if score >= 0.1 {
		return severityLow
	}
	return severityNegligible
}

func severityFromVulnerabilityItem(vulnerabilities *models.ImageSecurityDataLayerFeaturesItemsVulnerabilitiesItems) severity {
	if vulnerabilities == nil {
		return severityUnknown
	}

	// Clair may have figured out the severity for us.
	severity := severityFromClairSeverity(vulnerabilities.Severity)
	if severity != severityUnknown {
		return severity
	}

	// If we don't have a curated severity, we derive it from the CVSSv2 score.
	if vulnerabilities.Metadata == nil ||
		vulnerabilities.Metadata.NVD == nil ||
		vulnerabilities.Metadata.NVD.CVSSv2 == nil {
		return severityUnknown
	}
	score := vulnerabilities.Metadata.NVD.CVSSv2.Score
	return severityFromCVSSv2Score(score)
}

func scanRepository(registry *Registry, namespace, repository string) {
	log := logrus.WithFields(logrus.Fields{
		"namespace":  namespace,
		"repository": repository,
	})
	repositorySpec := namespace + "/" + repository

	tags, err := registry.GetTags(repositorySpec, 1, 1)
	if err != nil {
		log.Warn(err)
		return
	}
	if len(tags.Tags) == 0 {
		log.Warn("couldn't find any tag")
		return
	}

	log.Debugf("fetching vulnerabilities for tag '%s'", tags.Tags[0].Name)

	vulnerabilities, err := registry.GetVulnerabilities(repositorySpec, tags.Tags[0].DockerImageID)
	if err != nil {
		log.Warn(err)
		return
	}
	// The image may not have been scanned yet.
	if vulnerabilities.Status != "scanned" {
		return
	}

	// Collect the number of vulnerabilities for each (namespace,severity).
	features := vulnerabilities.Data.Layer.Features
	vulnerabilitiesMap := make(map[string]float64)
	for _, feature := range features {
		for _, vulnerability := range feature.Vulnerabilities {
			severity := severityFromVulnerabilityItem(vulnerability)
			key := string(severity) + "#" + vulnerability.NamespaceName
			vulnerabilitiesMap[key]++
		}
	}

	// Expose the corresponding prom metrics.
	for key, numVulnerabilities := range vulnerabilitiesMap {
		parts := strings.SplitN(key, "#", 2)
		severity := parts[0]
		os := parts[1]

		log.Debugf("found %d %s vulnerabilities for %s", int(numVulnerabilities), severity, os)

		quayVulnerabilities.With(prometheus.Labels{
			"organization": namespace,
			"repository":   repository,
			"os":           os,
			"severity":     severity,
		}).Set(numVulnerabilities)
	}
}

func scanVulnerabilities(registry *Registry, namespace string) {
	log := logrus.WithFields(logrus.Fields{
		"namespace": namespace,
	})

	log.Info("polling quay.io for vulnerabilities")

	repositories, err := registry.ListRepositories(namespace)
	if err != nil {
		log.Warn(err)
		return
	}

	for _, repository := range repositories {
		scanRepository(registry, repository.Namespace, repository.Name)
		// Be gentle with the service
		time.Sleep(2 * time.Second)
	}
}

type updateVulnerabilities struct {
	commands  <-chan command
	registry  *Registry
	namespace string
}

func taskUpdateVulnerabilities(args updateVulnerabilities, close <-chan bool) {
	for {
		select {
		case cmd := <-args.commands:
			switch cmd.op {
			case opScan:
				scanVulnerabilities(args.registry, args.namespace)
			}
		case <-close:
			return
		}
	}
}

type scanKicker struct {
	commands chan<- command
	interval time.Duration
}

func taskScanKicker(args scanKicker, close <-chan bool) {
	scan := command{op: opScan}

	for {
		args.commands <- scan
		select {
		case <-time.After(args.interval):
			continue
		case <-close:
			return
		}
	}
}

func setLogLevel(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("error parsing log level: %v", err)
	}
	logrus.SetLevel(level)
	return nil
}

func main() {
	quayToken := flag.String("quay-token", "", "quay.io OAuth 2 bearer token")
	quayPollInterval := flag.Duration("quay-poll-interval", 30*time.Minute, "quay.io poll interval")
	logLevel := flag.String("log-level", "info", "verbosity of log output - one of 'debug', 'info' (default), 'warning', 'error', 'fatal'")
	flag.Parse()

	if err := setLogLevel(*logLevel); err != nil {
		logrus.Fatal(err)
	}

	if flag.NArg() != 1 {
		logrus.Fatal("please specify the namespace to scan")
	}
	quayNamespace := flag.Arg(0)

	registry := NewRegistry()
	registry.SetBearerToken(*quayToken)

	var wg sync.WaitGroup

	// The main task that will fetch quay vulnerabilities data and expose prom
	// metrics out of it.
	commands := make(chan command)
	closeUpdate := make(chan bool)

	wg.Add(1)
	go func() {
		taskUpdateVulnerabilities(updateVulnerabilities{
			commands:  commands,
			registry:  registry,
			namespace: quayNamespace,
		}, closeUpdate)
		wg.Done()
	}()

	// A task kicking the main task at regular intervals to trigger a scan.
	closeKicker := make(chan bool)

	wg.Add(1)
	go func() {
		taskScanKicker(scanKicker{
			commands: commands,
			interval: *quayPollInterval,
		}, closeKicker)
		wg.Done()
	}()

	// Serve prom metrics.
	// TODO(damien): Make that into a nicely cancellable task
	http.Handle("/metrics", promhttp.Handler())
	logrus.Fatal(http.ListenAndServe(":8080", nil))

	close(closeUpdate)
	close(closeKicker)
	wg.Wait()
}
