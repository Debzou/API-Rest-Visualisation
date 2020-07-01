#!/bin/bash

# remove dependency
rm -rf vendor
rm go.mod
rm go.sum
# add dependency
go mod init main
go mod vendor
