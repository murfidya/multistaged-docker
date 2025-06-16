# 1st Stage: Build
FROM golang:latest AS builder

WORKDIR /app
COPY app.go .

RUN go build -o app app.go

# 2nd Stage: Run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]
