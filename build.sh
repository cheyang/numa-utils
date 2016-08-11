#!/bin/bash

go get github.com/tools/godep

mkdir -p $GOPATH/build

export PATH=$GOPATH/bin:$PATH

# Download the code
cd $GOPATH/src/github.com/cheyang

git clone https://github.com/cheyang/numa-utils.git

#create local cli
cd $GOPATH/src/github.com/cheyang/numa-utils/cmd/cli
godep go build -v -ldflags="-s" -o $GOPATH/build/gonumactl

STATUS=${?}

if [[ ${STATUS} -ne 0 ]]; then
  echo "Failed in building gonumactl"
  exit 1
fi

#create api server
cd $GOPATH/src/github.com/cheyang/numa-utils/cmd/service
godep go build -v -ldflags="-s" -o $GOPATH/build/numa-service

STATUS=${?}

if [[ ${STATUS} -ne 0 ]]; then
  echo "Failed in building numa-service"
  exit 1
fi