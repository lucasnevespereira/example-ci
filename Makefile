clean:
	rm example-ci

run:
	go run example-ci

test:
	go test ./tests

nodemon:
	nodemon --exec go run example-ci --signal SIGTERM

build:
	go build -o example-ci
