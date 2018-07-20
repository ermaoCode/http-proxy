#!/usr/bin/env bash

# start jvessel probe
/jvessel/satellite.sh

http-proxy

for (( ; ; ))
do
    sleep 5
    echo  `date` "hello world"
    echo  `date` "hello world" >> /helloworld/logs/helloworld.log
done