.PHONY: clean build

build:
	go build -o bin/remo cmd/remo.go

test:
	go test

clean:
	-rm -rf bin/*
