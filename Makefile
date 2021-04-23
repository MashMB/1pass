# Makefile for 1Pass application.
#
# @author TSS

app = 1pass-app
binary = 1pass
core = 1pass-core
parse = 1pass-parse
term = 1pass-term
test-all = test-core test-parse test-term test-app
test-all-simple = test-core.s test-parse.s test-term.s test-app.s

all: build run clean

build: $(test-all-simple)
	echo "--- Building $(binary) ---"
	cd $(app) && go build -o ../bin/$(binary)

clean:
	echo "--- Cleaning ---"
	rm -rf ./bin/

run:
	echo "--- Running $(binary) ---"
	./bin/$(binary)

test: $(test-all)

test-app:
	echo "--- Testing $(app) ---"
	cd $(app) && go test -v ./...

test-app.s:
	echo "--- Testing $(app) ---"
	cd $(app) && go test ./...

test-core:
	echo "--- Testing $(core) ---"
	cd $(core) && go test -v ./...

test-core.s:
	echo "--- Testing $(core) ---"
	cd $(core) && go test ./...

test-parse:
	echo "--- Testing $(parse) ---"
	cd $(parse) && go test -v ./...

test-parse.s:
	echo "--- Testing $(parse) ---"
	cd $(parse) && go test ./...

test-term:
	echo "--- Testing $(term) ---"
	cd $(term) && go test -v ./...

test-term.s:
	echo "--- Testing $(term) ---"
	cd $(term) && go test ./...
