#build stage
FROM golang:alpine AS builder

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o mainfile

#final stage
FROM alpine:latest

EXPOSE 8080

CMD ["./mainfile"]
