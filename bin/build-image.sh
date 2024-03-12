#!/bin/sh

export DOCKER_BUILDKIT=0
docker build . --tag dummyserver:v1.0 -f packages/dummyserver/Dockerfile