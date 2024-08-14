FROM golang:1.20-alpine AS builder_base_local
RUN apk --no-cache add ca-certificates
COPY devCert.crt /usr/local/share/ca-certificates/devCert.crt
RUN update-ca-certificates

FROM golang:1.20-alpine AS builder_base_other
RUN apk --no-cache add ca-certificates

FROM builder_base_local as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main src/main.go

# FROM golang:1.20-alpine AS builder
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN apk --no-cache add ca-certificates
# RUN update-ca-certificates
# RUN go mod download
# RUN go build -o main src/main.go

FROM golang:1.20-alpine as run
RUN apk --no-cache add ca-certificates
RUN update-ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
CMD ["./main"]