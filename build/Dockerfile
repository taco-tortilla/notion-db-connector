FROM golang:1.23.2-alpine3.20

RUN apk update && apk add git

RUN mkdir -p /app/notion-db-connector

WORKDIR /app/notion-db-connector

RUN go install github.com/air-verse/air@latest

COPY ./notion-db-connector .

CMD ["air"]