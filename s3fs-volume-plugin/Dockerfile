ARG UBUNTU_VERSION
FROM --platform=${TARGETPLATFORM:-linux/amd64} ubuntu:${UBUNTU_VERSION:-20.04} AS base

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt update && \
    apt install -y s3fs curl rsyslog tini && \
    mkdir -p /var/lib/rsyslog && \
    apt clean && rm -rf /var/lib/apt/lists/*

COPY s3fs-volume-plugin/rsyslog.conf /etc/rsyslog.conf
COPY s3fs-volume-plugin/fuse.conf /etc/fuse.conf

FROM --platform=${TARGETPLATFORM:-linux/amd64} golang:1.22-alpine AS dev

RUN apk update && apk add git

COPY s3fs-volume-plugin/ /go/src/github.com/sovarto/docker-volume-plugins/s3fs-volume-plugin
COPY mounted-volume/ /go/src/github.com/sovarto/docker-volume-plugins/mounted-volume

RUN cd /go/src/github.com/sovarto/docker-volume-plugins/s3fs-volume-plugin && \
    CGO_ENABLED=0 GOOS=linux go build -o /s3fs-volume-plugin

FROM base

COPY --from=dev /s3fs-volume-plugin /
