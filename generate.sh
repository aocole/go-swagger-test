#!/bin/bash
rm -rf client      cmd         models      restapi
swagger generate server -f swagger.yml
swagger generate client -f swagger.yml
go test -v ./...
