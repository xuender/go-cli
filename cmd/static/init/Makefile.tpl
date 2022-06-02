default: lint test

lint:
	golangci-lint run

test:
	go test ./... -gcflags=all=-l

clean:
	rm -rf dist

build:
	go build -o dist/{{ .Name }} main.go
