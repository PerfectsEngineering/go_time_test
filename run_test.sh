#!/bin/sh

#hwclock -s

go build -o time cmd/time/main.go

go test ./...

# For debugging the container, uncomment the below section
# before starting the container. So you can docker exec into the 
# container before it exits:

#ls -al /tmp
#
#tail -f time_test.go
