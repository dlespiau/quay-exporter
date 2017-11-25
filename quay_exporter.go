package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {

}

func main() {
	quayToken := flag.String("quay-token", "", "Quay.io OAuth 2 Bearer Token")
	flag.Parse()

	registry := NewRegistry()
	if *quayToken != "" {
		registry.SetBearerToken(*quayToken)
	}

	repositories, err := registry.ListRepositories("weaveworks")
	if err != nil {
		log.Fatal(err)
	}

	for _, repository := range repositories {
		fmt.Printf("%s/%s\n", repository.Namespace, repository.Name)
	}

	info, err := registry.GetRepository("weaveworks/ui-server")
	if err != nil {
		log.Fatal(err)
	}
	latest, ok := info.Tags["latest"]
	if ok {
		fmt.Println(latest.ManifestDigest)
	}
}
