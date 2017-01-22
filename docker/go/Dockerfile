FROM golang:latest

ADD ../../ /go/src/github.com/hilariousatlantic/mu-scheduler

RUN go get "github.com/lib/pq"
RUN go get "github.com/labstack/echo"

WORKDIR /go/src/github.com/hilariousatlantic/mu-scheduler

CMD ["/run.sh"]

RUN chmod +x run.sh
