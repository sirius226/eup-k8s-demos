#!/bin/bash

set -e

usage="./bin/docker/container.sh <app>"
app=${1?$usage}

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
image="sirius226/eup-demo-${app}:1.0"

docker build -t ${image} --build-arg APP_NAME=${app} --rm .
docker push ${image}
