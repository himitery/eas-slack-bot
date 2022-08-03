### Builder
FROM golang:1.18.5 as builder

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
COPY --from=builder /app .
CMD ["/main"]
