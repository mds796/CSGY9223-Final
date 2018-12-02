#!/usr/bin/env bash

pgrep CSGY9223-Final | xargs kill -9 && rm -fr .*.pid
pgrep etcd | xargs kill -9