FROM golang:1.14-alpine

WORKDIR /wear-store

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .