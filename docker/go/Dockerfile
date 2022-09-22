FROM golang:1.19.1-alpine3.16

RUN mkdir -p /go/app && cd /go/app
COPY ./backend /go/app
WORKDIR /go/app

RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0




ENV GO111MODULE=on

CMD ["air", "-c", ".air.toml"]