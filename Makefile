build:
	go build -o bin/ .

benchmark:
	go test -bench=. ./...

test:
	go test ./...
