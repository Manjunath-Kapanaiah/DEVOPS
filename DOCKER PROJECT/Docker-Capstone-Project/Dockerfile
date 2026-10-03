# Stage 1 - Builder

FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o registry-tracker ./cmd

# Stage 2 - Runtime

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/registry-tracker .
EXPOSE 8081
CMD ["./registry-tracker"]
