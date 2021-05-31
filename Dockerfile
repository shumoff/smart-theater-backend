FROM golang:1.16.4-buster

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /code

RUN apt-get update && apt-get -yqq install nano

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o smart ./src

WORKDIR /dist

RUN cp /code/smart .
