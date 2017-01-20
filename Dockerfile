FROM golang:latest

ADD . /go/src/github.com/hilariousatlantic/mu-scheduler

RUN go get "github.com/lib/pq"
RUN go get "github.com/labstack/echo"

RUN go install github.com/hilariousatlantic/mu-scheduler/server

ENTRYPOINT /go/bin/server

EXPOSE 3000
