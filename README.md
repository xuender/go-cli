# go-cli

[![GoCI](https://github.com/xuender/go-cli/workflows/Go/badge.svg)](https://github.com/xuender/go-cli/actions)
[![codecov](https://codecov.io/gh/xuender/go-cli/branch/main/graph/badge.svg?token=8CTpNIHxYT)](https://codecov.io/gh/xuender/go-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-cli)](https://goreportcard.com/report/github.com/xuender/go-cli)

CLI tool for Golang.

## Install

```shell
go install github.com/xuender/go-cli@latest
```

## init

Initialize the Golang project and create a default configuration file.

```shell
git clone url
cd dir
go-cli init
```

### github

Initialize the github configuration files.

```shell
git-cli init github
```

### gitee

Initialize the gitee configuration files.

```shell
git-cli init gitee
```

Use template.

```shell
go-cli init newName
```

## generate

Generate source code including commands, tests, examples, structures, protobuf, etc.

### cmd

Generate command support cobra and flag.

```shell
go-cli g c cmdName
go-cli g c cmdName -t flag
go-cli g c cmdName -t cobra
```

### struct

Generate struct and new function.

```shell
go-cli g s pkg/Book
```

### interface

Generate interface and comments.

```shell
go-cli g i pkg/Book
```

### test

Generate unit tests for exposed functions in file or directory.

```shell
go-cli g t pkg/book.go
```

### example

Generate test examples for exposed functions in file or directory.

```shell
go-cli g e pkg/book.go
```

### proto

Generate protobuf and comments.

```shell
go-cli g p pb/Book
go-cli g p pb/BookType -t enum
```

## template

Initialize template.

```shell
go-cli template 
```

Edit `~/.config/go-cli/*/*.tpl`

### New Initialization Template

```shell
mkdir ~/.config/go-cli/newName
vi ~/.config/go-cli/newName/xxx.tpl
# initialize by newName
go-cli init newName
```

## License

© xuender, 2022~time.Now

[MIT License](https://github.com/xuender/go-cli/blob/master/LICENSE)
