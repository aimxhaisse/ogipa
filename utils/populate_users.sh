#!/usr/bin/env bash

host=localhost
port=8080

for user in m1ch3l marc3l r0bert g4rsK00L
do
    curl -X PUT -d '{"login":"$user"}' "http://$host:$port/users/new"
done

