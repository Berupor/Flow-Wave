FROM golang:1.20-alpine

WORKDIR /
COPY .env .env

WORKDIR /app

COPY data_processing_service/go.mod ./go.mod
RUN go mod download

COPY data_processing_service/ .

RUN go build -o app

CMD ["./app"]
