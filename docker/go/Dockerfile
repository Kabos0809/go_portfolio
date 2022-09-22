FROM golang:1.19.1-alpine3.16

RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
RUN mkdir -p /app/go
WORKDIR /app/go

CMD ["air", "-c", ".air.toml"]