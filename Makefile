# Makefile for 1Pass application.
#
# @author TSS

app = 1pass-app
bin = bin
binary = 1pass
core = 1pass-core
out = out
parse = 1pass-parse
repo = github.com/mashmb/1pass/1pass-app
term = 1pass-term
test-all = test-core test-parse test-term test-app
test-all-simple = test-core.s test-parse.s test-term.s test-app.s
version = 0.0.0

all: build

build: $(test-all-simple)
	echo "--- Building $(binary) ---"
	cd $(app) && go build -o ../$(bin)/$(binary)

clean:
	echo "--- Cleaning ---"
	rm -rf ./$(bin)/
	rm -rf ./$(out)/

release: $(test-all-simple)
	echo "--- Building release $(binary) [$(version)] ---"
	cd $(app) && env GOOS=linux GOARCH=amd64 go build -o ../$(out)/$(binary) $(repo)
	cd $(out) && tar -czvf "$(binary)_$(version)_Linux_x86_64.tar.gz" $(binary)
	cd $(out) && rm -rf $(binary)

run:
	echo "--- Running $(binary) ---"
	./$(bin)/$(binary)

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
