FROM golang:1.24

WORKDIR /app

RUN go install github.com/google/wire/cmd/wire@v0.6.0 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2 && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENTRYPOINT ["tail", "-f", "/dev/null"]