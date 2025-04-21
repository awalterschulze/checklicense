FROM golang:1.24

WORKDIR /usr/src

COPY main.go .

RUN go build main.go

ENTRYPOINT ["./main"]
