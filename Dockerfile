FROM nexus.halykbank.nb:9503/golang:1.18-alpine-crt AS builder
WORKDIR /new
RUN export GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o main.go ./main.go

EXPOSE 80

