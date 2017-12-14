FROM alpine:3.7
LABEL maintainer="damien@weave.works"
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY ./quay-exporter /usr/bin/quay-exporter
EXPOSE 8080
ENTRYPOINT ["/usr/bin/quay-exporter"]
CMD ["-help"]
