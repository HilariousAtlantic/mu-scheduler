FROM golang:latest

ENV project $GOPATH/src/github.com/hilariousatlantic/mu-scheduler/
WORKDIR ${project}

# COPY manifest.json lock.json ${project}

CMD ./docker/go/build.sh
