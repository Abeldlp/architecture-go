build:
	@go build -o bin/fact

run: build
	./bin/fact
test:
	bo test -v ./...
