# M3

## user (mel)
* create
  * input: username
  * output: uuid
* view
  * input: uuid
  * output: username
* search
  * input: query
  * output: uuids

### stub
* uuid, username

## auth (mel)
* register
  * input: username, password
  * output: uuid
* login
  * input: username, password
  * output: uuid
* verify
  * input: cookie
  * output: uuid
* logout
  * input: uuid
  * output:

### stub
* uuid, hash(password), status

## follow (matheus)
* follow
  * input: uuid1, uuid2
  * output:
* unfollow
  * input: uuid1, uuid2
  * output:
* view
  * input: uuid
  * output: uuids

### stub
* uuid, uuids

## post (matheus)
* create
  * input: uuid, data
  * output: puid
* view
  * input: puid
  * output: data
* list
  * input: uuid
  * output: puids

### stub
* uuid, puid, data

## feed (miguel)
* view
  * input: uuid
  * output: puids

### stub
* uuid, puids


## Web

### Prerequisites

#### Node.js v11.x
```bash
$ sudo apt-get install -y curl
$ curl -sL https://deb.nodesource.com/setup_11.x | sudo -E bash -
$ sudo apt-get install -y nodejs
```

### Build
```bash
$ git clone git@github.com:mds796/CSGY9223-Final.git
$ cd CSGY9223-Final/static/
$ npm install
```

### How to run
Once you do, run `./run` from the root of the repository.
That command will execute `npm run build:static` from the `static` directory, execute `go build`, and finally run the Web Server. This allows the web server to serve updates to the static files.

If you prefer to run the commands yourself, run the following from the root of the repository:
1. cd static
1. npm run build:static
1. cd ../
1. go build
1. ./CSGY9223-Final web start&

To stop the web server:
1. ./CSGY9223-Final web stop
