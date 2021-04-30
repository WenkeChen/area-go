FROM golang:alpine

MAINTAINER WencoChen

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
