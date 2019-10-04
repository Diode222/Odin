FROM golang:1.13
MAINTAINER Diode "diodebupt@163.com"
WORKDIR $GOPATH/src/github.com/Diode222/Odin
ADD . $GOPATH/src/github.com/Diode222/Odin
ENV GO111MODULE on
RUN go mod download && go build main.go
EXPOSE 36666
ENTRYPOINT  ["./main"]
