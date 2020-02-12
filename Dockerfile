FROM golang:1.13.7

# Some developer tools
RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/go-delve/delve/cmd/dlv

RUN mkdir /app
WORKDIR /app

# Install dependencies only when these files change
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src/html ./html
COPY src/main.go ./
EXPOSE 40000 3002

ENTRYPOINT CompileDaemon \
    -log-prefix=false \
    -directory=/app \
    -build-dir=. \
    -color=true \
    -graceful-kill=true \
    -build="go build -gcflags=-N -gcflags=-l -o /app/demo ." \
    -command="./demo"

# Use this command for running delve
#    -command="/go/bin/dlv --listen=:40000 --headless=true --accept-multiclient --api-version=2 exec /app/demo"



