# go-cli

[![GoCI](https://github.com/xuender/go-cli/workflows/Go/badge.svg)](https://github.com/xuender/go-cli/actions)
[![codecov](https://codecov.io/gh/xuender/go-cli/branch/main/graph/badge.svg?token=8CTpNIHxYT)](https://codecov.io/gh/xuender/go-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-cli)](https://goreportcard.com/report/github.com/xuender/go-cli)

CLI tool for Golang.

## install

```shell
go install github.com/xuender/go-cli@latest
```

## init

init golang project

```shell
git clone url
cd dir
go-cli init
```

## generate

### cmd

```shell
go-cli g c cmdName
go-cli g c cmdName -t flag
go-cli g c cmdName -t cobra
```

### struct

```shell
go-cli g s dir/Book
```

### test

```shell
go-cli g t dir/book.go
```

### example

```shell
go-cli g e dir/book.go
```

## License

© xuender, 2022~time.Now

[MIT License](https://github.com/xuender/go-cli/blob/master/LICENSE)
