# FROM alpine:latest

# USER root

# # To be able to download `ca-certificates` with `apk add` command
# COPY my-root-ca.crt /root/my-root-ca.crt
# RUN cat /root/my-root-ca.crt >> /etc/ssl/certs/ca-certificates.crt

# # Add again root CA with `update-ca-certificates` tool
# RUN apk --no-cache add ca-certificates \
#     && rm -rf /var/cache/apk/*
# COPY my-root-ca.crt /usr/local/share/ca-certificates
# RUN update-ca-certificates

# RUN apk --no-cache add curl

# Builder
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production
FROM alpine:latest
RUN apk --no-cache update && apk add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

EXPOSE 8080
CMD ["./main"]