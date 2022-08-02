#!/usr/bin/env bash

aws lambda --region ap-northeast-2 update-function-code --function-name $(cat api_list.txt | grep $1 | awk '{print $3}') --zip-file fileb://deployment.zip