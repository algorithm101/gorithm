#!/bin/bash

export PKGS=$(go list ./... | grep -v /vendor | grep -v /test)
go test -p ${CPUS} ${PKGS} -coverprofile=coverage.out
go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $3 }'