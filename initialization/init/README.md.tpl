# {{ .Name }}

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![Lines of code][lines-svg]][lines-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

## 🚀 Install

```shell
go install {{ .Package }}@latest
```

## 💡 Usage

TODO

## 👤 Contributors

![Contributors][contributors-svg]

## 📝 License

© {{ .User }}, {{ .Year }}~time.Now

[{{ .License }} LICENSE][license-url]

[action-url]: https://{{ .Package }}/actions
[action-svg]: https://{{ .Package }}/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/{{ .Package }}
[goreport-svg]: https://goreportcard.com/badge/{{ .Package }}

[godoc-url]: https://godoc.org/{{ .Package }}
[godoc-svg]: https://godoc.org/{{ .Package }}?status.svg

[license-url]: https://{{ .Package }}/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-{{ .License }}-blue.svg

[contributors-svg]: https://contrib.rocks/image?repo={{ .Package | noweb }}

[lines-svg]: https://sloc.xyz/{{ .Package | url }}
[lines-url]: https://github.com/boyter/scc
