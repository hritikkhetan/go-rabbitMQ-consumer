# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY *.* ./

RUN go mod download

RUN go build -o /go-rabbitmq-consumer

EXPOSE 5672

CMD [ "/go-rabbitmq-consumer" ]