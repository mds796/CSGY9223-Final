#!/usr/bin/env bash

if [[ "$OSTYPE" == "darwin"* ]]; then
    pgrep CSGY9223-Final | xargs kill -9 && rm -fr .*.pid
    pgrep npm | xargs kill -9
    pgrep polymer | xargs kill -9
else
    pgrep CSGY9223-Final | xargs -r kill -9 && rm -fr .*.pid
    pgrep npm | xargs -r kill -9
    pgrep polymer | xargs -r kill -9
fi
