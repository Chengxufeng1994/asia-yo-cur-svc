#  Build Stage
FROM golang:1.19-alpine3.16 AS builder

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Run stage
FROM alpine:3.16

WORKDIR /usr/app

COPY --from=builder /usr/app/main .
COPY currencies.json .
COPY app.development.env .
COPY start.sh .

EXPOSE 8000

ENTRYPOINT ["/usr/app/start.sh"]
CMD ["/usr/app/main"]