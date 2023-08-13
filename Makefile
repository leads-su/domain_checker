PROGRAM_NAME = domain_checker
.DEFAULT_GOAL := build

#.PHONY:build

build:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o ./${PROGRAM_NAME}
