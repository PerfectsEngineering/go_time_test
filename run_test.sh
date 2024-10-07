#!/bin/sh

#hwclock -s

go build -o time cmd/time/main.go

go test ./...

ls -al /tmp

tail -f time_test.go
