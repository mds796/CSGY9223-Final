#!/bin/bash

# install NPM
# run: cd static && npm install && npm install -g polymer-cli

# The following error might occur:
#    $ panic: /debug/requests is already registered. You may have two independent copies of golang.org/x/net/trace in your binary, trying to maintain separate state. This may involve a vendored copy of golang.org/x/net/trace.
#
# This happens because both gRPC and etcd vendor "net/trace". In this case, run the following command to remove etcd's version of "net/trace":
#    $ rm -rf $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace
#
# Reference: https://github.com/etcd-io/etcd/issues/9357#issuecomment-377560659

function awaitServer()
{
    PORT=$1
    NAME=$2

    while :; do # success
        curl -s http://localhost:${PORT}/ >> /dev/null
        EXIT_CODE=$?

        if [[ ${EXIT_CODE} -eq 0 ]]; then # connection reset by peer
            break
        else
            if [[ ${EXIT_CODE} -eq 56 ]]; then # connection reset by peer
                break
            else
                echo "Waiting for $NAME server to start (exit code: $EXIT_CODE)..." && sleep 1
            fi
        fi
    done

    echo "$NAME has started."
}

./raft.sh start all

cd static && npm run build:static && npm run start -- --port=8000 &
awaitServer 8000 "static asset"

go build

./CSGY9223-Final user restart -H 0.0.0.0 -P 8081 &
awaitServer 8081 "user"

./CSGY9223-Final auth restart -H 0.0.0.0 -P 8082 &
awaitServer 8082 "auth"

./CSGY9223-Final post restart -H 0.0.0.0 -P 8083 &
awaitServer 8083 "post"

./CSGY9223-Final follow restart -H 0.0.0.0 -P 8084 &
awaitServer 8084 "follow"

./CSGY9223-Final feed restart -H 0.0.0.0 -P 8085 &
awaitServer 8085 "feed"

./CSGY9223-Final web restart -H 0.0.0.0 &
awaitServer 8080 "web"

./warmup.sh
