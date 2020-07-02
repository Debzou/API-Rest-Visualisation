#!/bin/bash
# update dependancy

# remove dependency
rm -rf vendor
rm go.mod
rm go.sum
# add dependency
go mod init github.com/Debzou/REST-API-GO
go mod vendor
