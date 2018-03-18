#!/usr/bin/env bash


export version=0.0.3
export containername=ethereum-node-inspector

echo "Test and Build Binary"
docker-compose -f docker-compose.yml run --rm unit
echo "Build image"
docker build -t $containername:$version .

echo "Pushing image to docker hub"
docker tag $containername:$version leondroid/$containername:$version
docker push leondroid/$containername:$version