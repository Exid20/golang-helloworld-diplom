FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y file
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
RUN ldd main || true
RUN file main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
