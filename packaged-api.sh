#!/bin/bash

# remove dependency
rm -rf vendor
rm go.mod
rm go.sum
# add dependency
go mod init api
go mod vendor