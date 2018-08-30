FROM golang:alpine

RUN apk add git
RUN go get github.com/mediocregopher/radix.v2

WORKDIR /app

COPY . /app/
RUN go build -o main .
CMD ["/app/main"]

