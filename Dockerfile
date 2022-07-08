FROM golang:1.18.3-alpine3.16 as builder

WORKDIR /go/src
COPY . .

CMD ["go", "run", "main.go"]
