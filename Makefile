lint:
	golint .

remove:
	rm -rf example-ci

test:
	go test ./tests


nodemon:
	nodemon --exec go run example-ci --signal SIGTERM


run:
	go run example-ci


build: remove
	go build -o example-ci
