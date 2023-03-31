default: test lint

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/spf13/cobra-cli@latest
	go install github.com/cespare/reflex@latest

test:
	go test -race -v ./... -gcflags=all=-l

watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v ./...'

clean:
	rm -rf dist

build:
	go build -o dist/go-cli main.go

proto:
	protoc --go_out=. pb/*.proto

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

msg:
	xgettext -C --add-comments=TRANSLATORS: --force-po -kT -kN:1,2 -kX:2,1c -kXN:2,3,1c -o doc/message.pot */*.go

msginit:
	msginit -l zh_CN --no-translator -i doc/message.pot -o doc/zh_CN.po
