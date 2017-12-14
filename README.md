# quay.io exporter

`quay-exporter` is a daemon exposing information about your quay.io
repositories as Prometheus metrics. Those metrics can be then used to monitor
the number and severity of vulnerabilities present in the docker images
published in that service.

## Install Instructions

### Compiling from sources

To run the daemon locally, use:

```shell
$ go get github.com/dlespiau/quay-exporter
$ quay-exporter weaveworks
```

`quay-exporter` can access private repositories when provided with an OAUTH 2
bearer token using the `-quay-token` command line parameter.

### Using the Docker image

Using `quay-expoter` from the published Docker image is one command away:

```shell
docker run -p 8080:8080 quay.io/damien.lespiau/quay-exporter weaveworks
```

## Visualize Metrics

To view the available metrics, point your browser at `http://localhost:8080/metrics/`:

```
quay_vulnerabilities{organization="weaveworks",os="debian:9",repository="build-golang",severity="critical"} 7
```

The latest tag of `weaveworks/build-golang` is running a Debian 9 image with
7 known critical vulnerabilities. Fortunately, `build-golang` is only used for
building containers images, not running services! Also rebuilding the image
will update the packages in the base image, which will fix the known
vulnerabilities.

## Troubleshooting

One can find more information about what the daemon is doing by increasing the log level:

```shell
$ quay-exporter -log-level debug  weaveworks
```
