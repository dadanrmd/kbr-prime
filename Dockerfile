FROM golang:1.16-alpine AS builder
ARG APP_DIR=kbrprime-be

RUN apk update && apk add tzdata

WORKDIR /

COPY ./$APP_DIR ./app

WORKDIR /app

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o kbrprime-be

###############################
FROM golang:1.16-alpine
ARG APP_DIR=kbrprime-be
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/kbrprime-be /app/kbrprime-be

# put env production to /app/params/.env
WORKDIR /app
RUN mkdir params
COPY ./$APP_DIR/migrations/sql/ migrations/sql
RUN mkdir -p /app/logs

# this port should be the same with on env
# EXPOSE 38019

ENTRYPOINT ["/app/kbrprime-be"]
