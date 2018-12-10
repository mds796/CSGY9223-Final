# M3

## Prerequisites

### go1.11
Please follow [these](https://medium.com/@RidhamTarpara/install-go-1-11-on-ubuntu-18-04-16-04-lts-8c098c503c5f) instructions.

### Node.js v11.x
```bash
sudo apt-get install -y curl
curl -sL https://deb.nodesource.com/setup_11.x | sudo -E bash -
sudo apt-get install -y nodejs
```

### Miscellaneous Dependencies
```bash
sudo apt-get install -y git-core
go get github.com/spf13/cobra
go get github.com/pkg/errors
go get github.com/etcd-io/etcd
go get github.com/mattn/goreman
go get github.com/gogo/protobuf/proto
go get github.com/google/uuid
rm -rf $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace
```

### gRPC
```bash
go get -u google.golang.org/grpc
```

### Protocol Buffers v3
```bash
cd Downloads/
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip protoc-3.6.1-linux-x86_64.zip -d protoc-3.6.1-linux-x86_64/
export PATH=$PATH:~/Downloads/protoc-3.6.1-linux-x86_64
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
```

## Build
```bash
mkdir ~/go/src/github.com/mds796
git clone git@github.com:mds796/CSGY9223-Final.git ~/go/src/github.com/mds796/CSGY9223-Final
cd ~/go/src/github.com/mds796/CSGY9223-Final/static/
npm install
```

## Run

### Start
```bash
$ cd CSGY9223-Final/static/
$ npm run build:static
$ cd ../
$ go build
$ ./CSGY9223-Final web start&
$ ./CSGY9223-Final user start&
$ ./CSGY9223-Final auth start&
$ ./CSGY9223-Final post start&
$ ./CSGY9223-Final follow start&
$ ./CSGY9223-Final feed start&
$ ./raft.sh&
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
$ ./CSGY9223-Final user stop
$ ./CSGY9223-Final auth stop
$ ./CSGY9223-Final post stop
$ ./CSGY9223-Final follow stop
$ ./CSGY9223-Final feed stop
$ pgrep etcd | xargs kill -9
```

Alternatively, you can just do:
```bash
$ cd CSGY9223-Final/
$ ./stop.sh
```
