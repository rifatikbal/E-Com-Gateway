#! /bin/bash

docker-compose up -d

set -e
set -o pipefail

until $(curl --output /dev/null --silent --fail http://0.0.0.0:8500/v1/kv); do
    echo 'waiting for consul'
    sleep 6
done

curl --request PUT --data-binary @config.example.yml http://localhost:8500/v1/kv/gateway

GO111MODULE=on CGO_ENABLED=0 go build -v .

export GATEWAY_CONSUL_URL="127.0.0.1:8500"
export GATEWAY_CONSUL_PATH="gateway"

# start application
./E-Com-Gateway migrate
