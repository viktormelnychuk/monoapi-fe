FROM golang

RUN mkdir -p /go/src/github.com/viktormelnychuk/monoapi

ADD . /go/src/github.com/viktormelnychuk/monoapi

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/viktormelnychuk/monoapi/monoapi/cmd  -watch github.com/viktormelnychuk/monoapi/monoapi
