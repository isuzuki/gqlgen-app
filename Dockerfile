FROM golang:1.16-alpine3.13

RUN go get -u github.com/cosmtrek/air

WORKDIR /go/src/app

COPY ../src .

RUN go get -d -v ./... && \
    go install -v ./...

CMD ["air"]