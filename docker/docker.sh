#!/usr/bin/env bash

docker build --tag lbry/sentinel:$TRAVIS_BRANCH ./
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
docker push lbry/sentinel:$TRAVIS_BRANCH