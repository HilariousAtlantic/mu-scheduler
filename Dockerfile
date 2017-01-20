FROM golang:latest

ADD . /go/src/github.com/hilariousatlantic/mu-scheduler

RUN go get "github.com/lib/pq"
RUN go get "github.com/labstack/echo"

WORKDIR /go/src/github.com/hilariousatlantic/mu-scheduler

RUN go build -o start_server server/*.go

CMD ["./start_server"]

EXPOSE 8000
