FROM golang:latest

WORKDIR /go/src/github.com/cloudydong/testgo

COPY . .

EXPOSE 8090

CMD ["go","run","main.go"]