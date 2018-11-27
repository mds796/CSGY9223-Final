# M3

## Prerequisites

### Go 1.11
Follow these instructions: https://medium.com/@RidhamTarpara/install-go-1-11-on-ubuntu-18-04-16-04-lts-8c098c503c5f

### Node.js v11.x
```bash
$ sudo apt-get install -y curl
$ curl -sL https://deb.nodesource.com/setup_11.x | sudo -E bash -
$ sudo apt-get install -y nodejs
```

### gRPC
```bash
$ go get -u google.golang.org/grpc
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
$ unzip protoc-3.6.1-linux-x86_64.zip -d protoc-3.6.1-linux-x86_64/ # Update the environment variable PATH to include the path to the protoc binary file.
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ export PATH=$PATH:$GOPATH/bin
```

## Build
```bash
$ git clone git@github.com:mds796/CSGY9223-Final.git ~/go
$ cd CSGY9223-Final/static/
$ npm install
$ go get github.com/google/uuid github.com/pkg/errors github.com/spf13/cobra
$ go get -u google.golang.org/grpc
```

## Run

### Start
```bash
$ cd CSGY9223-Final/static/
$ npm run build:static
$ cd ../
$ go build
$ ./CSGY9223-Final web start&
```

Alternatively, you can just do:
```bash
$ cd CSGY9223-Final/
$ ./run
```

This command will execute all of the above steps. It allows the web server to serve updates to the static files.

### Stop
```bash
$ cd CSGY9223-Final/
$ ./CSGY9223-Final web stop
```
