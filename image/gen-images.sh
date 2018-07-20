#!/usr/bin/env bash

set -e
cd `dirname $0`

image_name="http-proxy"

go build ../http-proxy.go

version=`date "+%Y%m%d"`
docker build -t ${image_name}:$version .
docker tag ${image_name}:$version ${image_name}:latest