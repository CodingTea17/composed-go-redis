FROM golang:alpine

RUN apk add git
RUN go get -u github.com/go-redis/redis

WORKDIR /app

COPY . /app/
RUN go build -o main .
CMD ["/app/main"]

