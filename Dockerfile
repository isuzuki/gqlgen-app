FROM golang:1.16-alpine3.13

WORKDIR /go/src/app

COPY ../src .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]