FROM golang:1.21 AS builder

WORKDIR /app
COPY . .
RUN go build -o myapp

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 8080
CMD ["./ci-cd"]
