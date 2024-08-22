ARG UBUNTU_VERSION
FROM ubuntu:${UBUNTU_VERSION:-20.04} AS base

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt update && \
    apt install -y curl rsyslog tini fuse libcurl4-gnutls-dev libxml2 && \
    mkdir -p /var/lib/rsyslog && \
    apt clean && rm -rf /var/lib/apt/lists/*

COPY src/rsyslog.conf /etc/rsyslog.conf
COPY src/fuse.conf /etc/fuse.conf

FROM ubuntu:${UBUNTU_VERSION:-20.04} AS compile

ENV DEBIAN_FRONTEND="noninteractive"

RUN <<EOF
    apt update -y
    apt install -y curl make automake build-essential gcc openssl libssl-dev fuse libfuse-dev unzip libtool \
          pkg-config libcurl4-gnutls-dev libxml2 libxml2-dev
    curl -L https://github.com/s3fs-fuse/s3fs-fuse/archive/refs/tags/v1.94.zip -o s3fs-fuse.zip
    unzip s3fs-fuse.zip
    cd s3fs-fuse-1.94
    ./autogen.sh
    ./configure
    make
    make install
EOF

FROM golang:1.22-alpine AS dev

RUN apk update && apk add git

COPY src /src/

RUN cd /src/ && CGO_ENABLED=0 GOOS=linux go build -o /docker-s3fs

FROM base

COPY --from=dev /docker-s3fs /
COPY --from=compile /usr/local/bin/s3fs /usr/local/bin/s3fs