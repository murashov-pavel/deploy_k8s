FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY cmd/sandbox ./cmd/sandbox
RUN go mod download
RUN go build -o sandbox ./cmd/sandbox
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sandbox .
EXPOSE 8080
RUN chmod +x sandbox
CMD ["./sandbox"]

