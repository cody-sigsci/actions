#!/bin/bash

echo "Running gosec in CWD"
curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $GOROOT/bin latest
go get -d ./...
$GOROOT/bin/gosec --include=G101 -fmt=json -out=results.json ./...
go run /gosec-report/main.go