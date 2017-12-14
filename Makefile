REGISTRY      = quay.io
NAMESPACE    = damien.lespiau
VERSION      = $(shell git symbolic-ref --short HEAD)-$(shell git rev-parse --short HEAD)
IMAGE_BASE   = $(REGISTRY)/$(NAMESPACE)/quay-exporter
IMAGE        = $(IMAGE_BASE):$(VERSION)
IMAGE_LATEST = $(IMAGE_BASE):latest

all: image

clean:
	go clean

FORCE:

quay-exporter: FORCE
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./

.image.done: Dockerfile quay-exporter
	docker build -t $(IMAGE) -t $(IMAGE_LATEST) .
	touch $@

image: .image.done

push-image: image
	docker push $(IMAGE)
	docker push $(IMAGE_LATEST)

.PHONY: all clean image
