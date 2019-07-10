FROM golang:1.8

ADD main.go /go/src/main.go
ADD mmo-plugin.yaml /plugin/mmo-plugin.yaml

RUN go build -o /hooks/main /go/src/main.go
ENTRYPOINT ["/hooks/main"]