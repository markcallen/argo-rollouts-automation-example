FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY app/ .
RUN go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["/app/server"]
