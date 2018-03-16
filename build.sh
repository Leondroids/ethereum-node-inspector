#!/usr/bin/env bash

containername=ethereum-node-inspector

echo "Test and Build Binary"
docker-compose  run --rm unit
docker build -t $containername .
