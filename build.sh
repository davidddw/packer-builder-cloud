#!/bin/bash

export GOPATH=$(pwd)

cd ${GOPATH}/src/github.com/mitchellh/gox
go build
cd -

ln -s ${GOPATH}/src/github.com/mitchellh/gox/gox /usr/bin/gox
sh src/github.com/mitchellh/packer/scripts/build.sh 
rm -fr /usr/bin/gox
