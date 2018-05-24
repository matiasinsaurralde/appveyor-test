#!/bin/bash
export GOPATH=/c/gopath
cd $GOPATH/src/github.com/matiasinsaurralde/appveyor-test/dotnet
go get
# cd $GOPATH/src/github.com/matiasinsaurralde/go-dotnet/dotnet
go test -v