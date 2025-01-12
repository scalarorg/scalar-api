FROM golang:1.23.3-alpine AS builder

ARG VERSION="HEAD"

RUN apk add --no-cache \
    make \
    git \
    build-base \
    linux-headers \
    libc-dev \
    pkgconfig \
    alpine-sdk \
    libsodium-dev \
    libsodium-static \
    openssh

# Build
WORKDIR /go/src/github.com/scalarorg/xchains-api
# Cache dependencies
COPY go.mod go.sum /go/src/github.com/scalarorg/xchains-api/

# Copy the rest of the files
COPY ./ /go/src/github.com/scalarorg/xchains-api/

RUN LDFLAGS='-extldflags "-static" -v' \
    CGO_ENABLED=1 \
    CGO_LDFLAGS="$CGO_LDFLAGS -lstdc++ -lm -lsodium" \
    LINK_STATICALLY=true \
    BUILD_TAGS="muslc" \
    make build

# Final minimal image with binary only
FROM alpine:3.16 AS run

RUN addgroup --gid 1138 -S xchains-api && adduser --uid 1138 -S xchains-api -G xchains-api
RUN apk add bash curl jq

# Label should match your github repo
LABEL org.opencontainers.image.source="https://github.com/scalarorg/xchains-api:${VERSION}"

# Copy over binaries from the build-env
COPY --from=builder /go/src/github.com/scalarorg/xchains-api/build/xchains-api /bin/xchains-api
COPY --from=builder /go/src/github.com/scalarorg/xchains-api/contrib/images/xchains-api/entrypoint.sh /bin/entrypoint.sh

WORKDIR /home/xchains-api
RUN chown -R xchains-api /home/xchains-api
USER xchains-api

ENTRYPOINT ["/bin/entrypoint.sh"]
CMD []
STOPSIGNAL SIGTERM
