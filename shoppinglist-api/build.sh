#!/bin/bash

TAG="0.1"
if [ -f VERSION ]; then
    echo "Current Verdsion $TAG"
else
    echo "VERSION file not found create a file called VERSION in the working directory"
    exit 1
fi

TAG=$(cat VERSION)
TAG=$(echo $TAG | awk -F. '{printf "%d.%d", $1, $2+1}')

docker build -t kdkanishka/shoppinglist-api:$TAG .
docker push kdkanishka/shoppinglist-api:$TAG

echo $TAG > VERSION