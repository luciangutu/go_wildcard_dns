FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o wildcard-dns-server main.go


FROM alpine:latest

RUN apk --no-cache add bind-tools
WORKDIR /app
COPY --from=builder /app/wildcard-dns-server .
EXPOSE 53/udp

ENTRYPOINT ["./wildcard-dns-server", "-ip"]
