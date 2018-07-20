#!/usr/bin/env bash

health=`ps aux | grep heep-proxy | grep -v "grep"`

if [ -z "$health" ]; then
    echo unhealth >> /helloworld/logs/health.log
    exit 1
fi

echo health >> /helloworld/logs/health.log
exit 0