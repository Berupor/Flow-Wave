FROM golang:1.20-alpine

WORKDIR /
COPY .env .env

WORKDIR /app

COPY etl_service/go.mod ./go.mod
RUN go mod download

COPY etl_service/ .

RUN go build -o app

CMD ["./app"]
