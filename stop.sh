#!/usr/bin/env bash

pgrep CSGY9223-Final | xargs -r kill -9 && rm -fr .*.pid
pgrep etcd | xargs -r kill -9
pgrep polymer | xargs -r kill -9