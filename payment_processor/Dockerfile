FROM golang:1.20 AS builder

WORKDIR /go/src
RUN apt-get update && apt-get install build-essential librdkafka-dev -y
COPY . .
RUN go build -o /go/bin/processor cmd/main.go 

CMD ["/go/bin/processor"]
