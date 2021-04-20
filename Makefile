# Makefile for 1Pass application.
#
# @author TSS

all: test.s build run clean

test:
	go test -v ./...

test.s:
	go test ./...

build:
	go build -o ./bin/1pass

run:
	./bin/1pass

clean:
	rm -rf ./bin/
