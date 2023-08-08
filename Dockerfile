#--- Build stage
FROM golang:1.20-bullseye AS go-builder

WORKDIR /src

COPY . /src/

RUN make build CGO_ENABLED=0

#--- Image stage
FROM alpine:3.18.3

COPY --from=go-builder /src/target/dist/template-go /usr/bin/template-go

WORKDIR /opt

ENTRYPOINT ["/usr/bin/template-go"]
