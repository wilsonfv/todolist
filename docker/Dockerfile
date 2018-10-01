FROM golang:1.9.2

WORKDIR /go/src/github.com/wilsonfv/todolist

COPY app app
COPY Gopkg.toml Gopkg.lock ./

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

CMD ["go", "run", "-v app/app_server.go"]