#!/usr/bin/env sh

set -xe

# make tags
etags -Q --declarations **/*.go

# build
go build
