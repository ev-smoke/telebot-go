FROM --platform=$BUILDPLATFORM quay.io/projectquay/golang:1.23 AS builder
ARG PLATFORM
ARG ARCH

WORKDIR /go/src/app
COPY . .
ARG TARGETARCH
RUN make build TARGETARCH=$TARGETARCH

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/telebot-go .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
ENTRYPOINT ["/telebot-go" , "go"]