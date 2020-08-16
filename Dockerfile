FROM golang:1.13-alpine
RUN apk add git gcc g++
ENV CGO_ENABLED 1
ENV GOPATH /go
ENV CC gcc

WORKDIR /go/src/api
COPY src/api .

RUN go get -u github.com/cosmtrek/air
ENTRYPOINT  air run main.go
EXPOSE 8080