#!/usr/bin/env sh

[ "$1" = "debug" ] && set -x
set -e

PORT_TO_RUN_ON=:9000

if [ "$1" = "debug" ]; then
    GOTRACEBACK=crash APP_PORT=$PORT_TO_RUN_ON ./go-todo
else
    APP_PORT=$PORT_TO_RUN_ON ./go-todo
fi
