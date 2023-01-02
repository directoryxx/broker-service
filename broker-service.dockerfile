# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

RUN apk add build-base librdkafka-dev pkgconf

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=1 go build -tags musl -o ./build/brokerApp ./internal/app

RUN chmod +x /app/build/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/build/brokerApp /app

COPY ./.env /.env

CMD [ "/app/brokerApp" ]