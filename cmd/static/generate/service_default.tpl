package {{ .Package }}

// {{ .Name }} TODO.
type {{ .Name }} struct{}

// New{{ .Name }} TODO.
func New{{ .Name }}() *{{ .Name }}{
	return &{{ .Name }}{}
}

// String TODO.
func (p *{{ .Name }}) String() string {
	return "{{ .Name }}"
}
