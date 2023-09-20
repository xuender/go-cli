# go-cli

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![godoc][godoc-svg]][godoc-url]
[![Lines of code][lines-svg]][lines-url]
[![License][license-svg]][license-url]

CLI tool for Golang.

## 🚀 Install

```shell
go install github.com/xuender/go-cli@latest
```

## 💡 Usage

### init

Initialize the Golang project and create a default configuration file.

```shell
git clone url
cd dir
go-cli init
```

#### github

Initialize the github configuration files.

```shell
git-cli init github
```

#### gitee

Initialize the gitee configuration files.

```shell
git-cli init gitee
```

Use template.

```shell
go-cli init newName
```

### generate

Generate source code including commands, tests, examples, structures, protobuf, etc.

#### cmd

Generate command support cobra and flag.

```shell
go-cli g c cmdName
go-cli g c cmdName -t flag
go-cli g c cmdName -t cobra
```

#### struct

Generate struct and new function.

```shell
go-cli g s pkg/Book
```

#### interface

Generate interface and comments.

```shell
go-cli g i pkg/Book
```

#### test

Generate unit tests for exposed functions in file or directory.

```shell
go-cli g t pkg/book.go
```

#### example

Generate test examples for exposed functions in file or directory.

```shell
go-cli g e pkg/book.go
```

#### proto

Generate protobuf and comments.

```shell
go-cli g p pb/Book
go-cli g p pb/BookType -t enum
```

### struct

Struct related.

#### new

Create a new struct function by other struct.

```shell
go-cli s n book/book.go pb/book.pb.go
```

```go
// NewBookByPbBook creates a new Book of pb.Book.
func NewBookByPbBook(elem *pb.Book) *Book {
  return &Book{
    ID:    elem.ID,
    Title: elem.Title,
  }
}
```

### convert

Convert struct to other structs.

```shell
go-cli s c book/book.go pb/book.pb.go
```

```go
// FromPbBook from pb.Book.
func (p *Book) FromPbBook(elem *pb.Book) *Book {
  p.ID = elem.ID
  p.Title = elem.Title

  return p
}

// ToPbBook to pb.Book.
func (p *Book) ToPbBook() *pb.Book {
  return &pb.Book{
    ID:    p.ID,
    Title: p.Title,
  }
}
```

### template

Initialize template.

```shell
go-cli template 
```

Edit `~/.config/go-cli/*/*.tpl`

#### New Initialization Template

```shell
mkdir ~/.config/go-cli/newName
vi ~/.config/go-cli/newName/xxx.tpl
# initialize by newName
go-cli init newName
```

## 👤 Contributors

![Contributors][contributors-svg]

## 📝 License

© ender, 2023~time.Now

[MIT LICENSE][license-url]

[action-url]: https://github.com/xuender/go-cli/actions
[action-svg]: https://github.com/xuender/go-cli/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/github.com/xuender/go-cli
[goreport-svg]: https://goreportcard.com/badge/github.com/xuender/go-cli

[godoc-url]: https://godoc.org/github.com/xuender/go-cli
[godoc-svg]: https://godoc.org/github.com/xuender/go-cli?status.svg

[license-url]: https://github.com/xuender/go-cli/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg

[contributors-svg]: https://contrib.rocks/image?repo=xuender/go-cli
[lines-svg]: https://sloc.xyz/github/xuender/go-cli
[lines-url]: https://github.com/boyter/scc
