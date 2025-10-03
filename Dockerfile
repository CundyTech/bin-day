# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o bin-day main.go

# Run stage
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/bin-day .
EXPOSE 8080
ENTRYPOINT ["/app/bin-day"]
