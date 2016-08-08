#!/bin/bash

go get github.com/tools/godep

# Download the code
cd $GOPATH/src/github.com/cheyang

git clone https://github.com/cheyang/numa-utils.git

#create version go file
cd $GOPATH/src/github.com/cheyang/numa-utils/cmd
godep go build -v -ldflags="-s" -o $GOPATH/bin/gonumactl

STATUS=${?}

if [[ ${STATUS} -ne 0 ]]; then
  echo "Failed in building gonumactl"
  exit 1
fi

