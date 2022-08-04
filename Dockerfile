### Builder
FROM golang:1.18.5 as builder
RUN apt-get update && apt-get install -y ca-certificates

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main cmd/eas_slack_bot.go


### Make executable image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app .
ENTRYPOINT ["/main"]
