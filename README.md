# M3

## Prerequisites

### Node.js v11.x
```bash
$ sudo apt-get install -y curl
$ curl -sL https://deb.nodesource.com/setup_11.x | sudo -E bash -
$ sudo apt-get install -y nodejs
```

## Build
```bash
$ git clone git@github.com:mds796/CSGY9223-Final.git
$ cd CSGY9223-Final/static/
$ npm install
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
