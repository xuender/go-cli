VERSION = $(shell git describe --tags)
BUILD_TIME = $(shell date +%F' '%T)

default: lint-fix test

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/spf13/cobra-cli@latest
	go install github.com/cespare/reflex@latest

test:
	go test -race -v ./... -gcflags=all=-l -cover

watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -v ./...'

clean:
	rm -rf dist

build:
	go build \
	-ldflags "-X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/go-cli main.go

install:
	go install \
	-ldflags "-X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \

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

msgfmt:
	msgfmt doc/zh_CN.po -o locales/zh_CN.mo

watchtest:
	go run main.go w ./a.sh
