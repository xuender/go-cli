# {{ .Name }}

[![Action][action-svg]][action-url]
[![Report Card][goreport-svg]][goreport-url]
[![godoc][godoc-svg]][godoc-url]
[![License][license-svg]][license-url]

## Install

```shell
go install {{ .Package }}@latest
```

## License

© {{ .User }}, {{ .Year }}~time.Now

[{{ .License }} LICENSE](https://{{ .Package }}/blob/master/LICENSE)

[action-url]: https://{{ .Package }}/actions
[action-svg]: https://{{ .Package }}/workflows/Go/badge.svg

[goreport-url]: https://goreportcard.com/report/{{ .Package }}
[goreport-svg]: https://goreportcard.com/badge/{{ .Package }}

[godoc-url]: https://godoc.org/{{ .Package }}
[godoc-svg]: https://godoc.org/{{ .Package }}?status.svg

[license-url]: https://{{ .Package }}/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-{{ .License }}-blue.svg
