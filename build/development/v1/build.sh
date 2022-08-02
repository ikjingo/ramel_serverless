#!/usr/bin/env bash

GOOS=linux go build -o main $(cat api_list.txt | grep $1 | awk '{print $2}')
zip deployment.zip main