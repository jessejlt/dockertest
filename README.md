dockertest
---

An attempt to vendor dependencies w/o using something like [godep](https://github.com/tools/godep).

There is a simple http server that response to route `/bar` on port 8080. It's functionality is meaningless and is only in place as a testbed for dependencies, in this case, [log15](https://github.com/inconshreveable/log15)

Use
---

* Build the container
	- `make build`
* Run the app in a development scenario
	- `make dev`
* Execute the unit tests
	- `make test`
* Vendor a dependency
	- `make deps` then just `go get <dependency>`
