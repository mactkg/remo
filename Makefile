.PHONY: clean build

build:
	go build -o bin/remo cmd/remo.go

test:
	go test

clean:
	-rm -rf bin/*

install:
	cp bin/remo /usr/local/bin/

install-home-local: build
	cp bin/remo ${HOME}/local/bin
