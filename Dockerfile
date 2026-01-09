FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# 1. We don't need 'go mod download' anymore!
# We just copy everything, including the new 'vendor' folder
COPY . .

# 2. Build using the '-mod=vendor' flag
# This tells Go to ignore the internet and use the local folder
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main ./cmd/api

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]