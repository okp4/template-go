#--- Build stage
FROM golang:1.18.3-stretch AS go-builder

WORKDIR /src

COPY . /src/

RUN make build CGO_ENABLED=0

#--- Image stage
FROM alpine:3.16.1

COPY --from=go-builder /src/target/dist/template-go /usr/bin/template-go

WORKDIR /opt

ENTRYPOINT ["/usr/bin/template-go"]
