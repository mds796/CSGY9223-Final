#!/usr/bin/env bash

CMD=$1
NODE=$2

if [[ $CMD == "start" ]]; then
    if [[ $NODE == "all" ]]; then
        # Start node0 in standalone mode
        ./CSGY9223-Final raftkv restart -N node0 -H localhost -P 7000 &

        # Wait for an election
        sleep 3

        # Start node1 and node2 and join cluster at localhost:7000
        ./CSGY9223-Final raftkv restart -N node1 -H localhost -P 7010 --joinHost localhost --joinPort 7000 &
        ./CSGY9223-Final raftkv restart -N node2 -H localhost -P 7020 --joinHost localhost --joinPort 7000 &

        # Wait raft nodes to join and update configuration
        sleep 2
    elif [[ $NODE == "node0" ]]; then
        ./CSGY9223-Final raftkv restart -N node0 -H localhost -P 7000 &
    elif [[ $NODE == "node1" ]]; then
        ./CSGY9223-Final raftkv restart -N node1 -H localhost -P 7010 --joinHost localhost --joinPort 7000 &
    elif [[ $NODE == "node2" ]]; then
        ./CSGY9223-Final raftkv restart -N node2 -H localhost -P 7020 --joinHost localhost --joinPort 7000 &
    else
        echo "Choose which node to start: node0, node1, node2 or all"
    fi
elif [[ $CMD == "restart" ]]; then
    if [[ $NODE == "node0" ]]; then
        # Restart node0 and join cluster at localhost:7010
        ./CSGY9223-Final raftkv restart -N node0 -H localhost -P 7000 --joinHost localhost --joinPort 7010 &
    elif [[ $NODE == "node1" ]]; then
        ./CSGY9223-Final raftkv restart -N node1 -H localhost -P 7010 --joinHost localhost --joinPort 7000 &
    elif [[ $NODE == "node2" ]]; then
        ./CSGY9223-Final raftkv restart -N node2 -H localhost -P 7020 --joinHost localhost --joinPort 7000 &
    else
        echo "Choose which node to restart: node0, node1, node2"
    fi
elif [[ $CMD == "stop" ]]; then
    if [[ $NODE == "all" ]]; then
        ./CSGY9223-Final raftkv stop -N node0
        ./CSGY9223-Final raftkv stop -N node1
        ./CSGY9223-Final raftkv stop -N node2
    elif [[ $NODE == "node0" ]]; then
        ./CSGY9223-Final raftkv stop -N node0
    elif [[ $NODE == "node1" ]]; then
        ./CSGY9223-Final raftkv stop -N node1
    elif [[ $NODE == "node2" ]]; then
        ./CSGY9223-Final raftkv stop -N node2
    else
        echo "Choose which node to stop: node0, node1, node2 or all"
    fi
else
    echo "Unknown command '$CMD'"
fi
