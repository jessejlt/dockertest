.PHONY: build dev test clean

build:
	docker build -t dockertest .

dev: build test
	docker run -a STDOUT -a STDERR -c 1 -m 256m -p 8080:8080 --restart=no dockertest

test:
	docker run -a STDOUT -a STDERR -c 1 -m 256m --restart=no --entrypoint="/bin/go-test" -v $$(pwd):/gopath/src/app -e GOPATH=/gopath/src/app/vendor dockertest

deps:
	docker run -v $$(pwd)/vendor:/gopath/src/app/vendor -e GOPATH=/gopath/src/app/vendor --entrypoint /bin/bash -i -t dockertest
