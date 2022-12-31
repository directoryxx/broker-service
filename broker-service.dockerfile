# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o ./build/brokerApp ./internal/app

RUN chmod +x /app/build/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/build/brokerApp /app

COPY ./.env /.env

CMD [ "/app/brokerApp" ]