#!/usr/bin/env bash

./CSGY9223-Final raftkv restart -N node0 -H localhost -P 7000 &

# Wait for an election
sleep 5

./CSGY9223-Final raftkv restart -N node1 -H localhost -P 7010 --joinHost localhost --joinPort 7000 &
./CSGY9223-Final raftkv restart -N node2 -H localhost -P 7020 --joinHost localhost --joinPort 7000 &