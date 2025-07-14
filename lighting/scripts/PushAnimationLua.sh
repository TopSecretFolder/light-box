#!/bin/bash

ENDPOINT=$1
LUA=$2

curl -X POST "$ENDPOINT/animation/push/lua" \
  -H "Accept: */*" \
  -H "Cache-Control: no-cache" \
  -H "Content-Type: text/plain" \
  -d "$(cat $LUA)"
