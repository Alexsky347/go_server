FROM golang:1.21-alpine3.18

WORKDIR /app
COPY ../ .
RUN go mod download

COPY ./pkg/config/config.sample.yaml ./pkg/config/config.yaml
RUN go build -o /app/go-ddd ./cmd/api

EXPOSE 8888
ENTRYPOINT ["/app/go-ddd"]
