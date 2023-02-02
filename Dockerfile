FROM golang:latest
RUN apt-get update
WORKDIR /go/src/app
COPY . .
CMD go run examples/fullapp/cmd/main.go

