FROM golang:1.11.2

RUN go get github.com/pilu/fresh

RUN mkdir /app

WORKDIR /app

ADD go.mod \
    go.sum \
    *.go \
    index.html \
    ./

RUN go build ./...

CMD ["fresh"]




