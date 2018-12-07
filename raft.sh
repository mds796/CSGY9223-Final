#!/usr/bin/env bash

readonly GO_PATH=$(go env GOPATH)

cd ${GO_PATH}/src/go.etcd.io/etcd && go build && ./build && ${GO_PATH}/bin/goreman start