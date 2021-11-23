.PHONY: clean
clean:
	go clean -i .

.PHONY: lint
lint: clean
	golint .

remove:
	rm -rf example-ci

.PHONY: test
test: lint
	go test ./tests

.PHONY: nodemon
nodemon:
	nodemon --exec go run example-ci --signal SIGTERM

.PHONY: run
run:
	go run example-ci

.PHONY: build
build: remove
	go build -o example-ci
