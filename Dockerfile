FROM golang:1.18.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git nano

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

CMD go run .

