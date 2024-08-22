FROM golang:1.22-alpine AS build

RUN <<EOF
    apk update
    apk --no-cache add \
        ca-certificates \
        build-base \
        git \
        alpine-sdk \
        curl \
        libcurl \
        automake \
        autoconf \
        libxml2-dev \
        mailcap \
        fuse-dev \
        curl-dev \
        unzip

    curl -L https://github.com/s3fs-fuse/s3fs-fuse/archive/refs/tags/v1.94.zip -o s3fs-fuse.zip
    unzip s3fs-fuse.zip
    cd s3fs-fuse-1.94
    ./autogen.sh
    ./configure
    make -j
    make install
EOF

COPY src /src/

RUN cd /src/ && CGO_ENABLED=0 GOOS=linux go build -o /docker-s3fs

FROM alpine:latest

LABEL org.opencontainers.image.title="aekis/docker-s3fs"
LABEL org.opencontainers.image.description="Provides a Docker Volume Driver for Mounting Amazon S3 Buckets using S3FS"
LABEL org.opencontainers.image.authors="Axel Mendoza <axel@aekis.dev>"
LABEL org.opencontainers.image.url="https://github.com/aekis-dev/docker-s3fs"
LABEL org.opencontainers.image.documentation="https://github.com/aekis-dev/docker-s3fs/README.md"
LABEL org.opencontainers.image.source="https://github.com/aekis-dev/docker-s3fs/src/Dockerfile"

RUN <<EOF
    apk update
    apk --no-cache add \
        fuse \
        libcurl \
        libxml2 \
        libgcc \
        libstdc++ \
        mailcap \
        ca-certificates \
        rsyslog \
        tini
    deluser xfs
    mkdir -p /var/lib/rsyslog
    rm -rf /var/cache/apk/*
EOF

COPY --from=build /docker-s3fs /
COPY --from=build /usr/local/bin/s3fs /usr/local/bin/s3fs

COPY src/rsyslog.conf /etc/rsyslog.conf
COPY src/fuse.conf /etc/fuse.conf
