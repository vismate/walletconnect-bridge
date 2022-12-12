FROM golang:1.19 AS build

RUN useradd -u 1000 dimo

WORKDIR /go/src/github.com/DIMO-Network/walletconnect-bridge/
COPY . /go/src/github.com/DIMO-Network/walletconnect-bridge/

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOFLAGS=-mod=vendor
ENV GOARCH=amd64

RUN go mod vendor
RUN go mod tidy
RUN make build-linux

FROM busybox AS package

LABEL maintainer="DIMO <hello@dimo.zone>"

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /go/src/github.com/DIMO-Network/walletconnect-bridge/walletconnect-bridge /app

USER dimo

EXPOSE 8080
RUN ls /app
ENTRYPOINT ["/app/walletconnect-bridge"]
