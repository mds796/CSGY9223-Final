#!/usr/bin/env bash

if [[ "$OSTYPE" == "darwin"* ]]; then
    pgrep etcd | xargs echo | cut -d ' ' -f1 | xargs kill -9
else
    pgrep etcd | xargs -r echo | cut -d ' ' -f1 | xargs -r kill -9
fi