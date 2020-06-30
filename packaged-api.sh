#!/bin/bash

# remove dependency
rm go.mod
rm go.sum
# add dependency
go mod init api
go mod vendor