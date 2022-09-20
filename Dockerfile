FROM golang:1.18-alpine

MAINTAINER fischer

WORKDIR /go/src/SportAnalyse
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . .
RUN go get -d -v ./...

RUN go install -v ./...

CMD["SignInBoard"]
