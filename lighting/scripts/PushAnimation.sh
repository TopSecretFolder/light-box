#!/bin/bash

ENDPOINT=$1
JSON=$2

curl -X POST "$ENDPOINT/animation/push" \
  -H "Accept: */*" \
  -H "Cache-Control: no-cache" \
  -H "Content-Type: application/json" \
  -d "$(cat $JSON)"
