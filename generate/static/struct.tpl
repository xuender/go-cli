
// {{ .Name }} is a struct.
type {{ .Name }} struct{}

// New{{ .Name }} creates a new instance of {{ .Name }}.
func New{{ .Name }}() *{{ .Name }} {
	return &{{ .Name }}{}
}

// String returns a string representation of the {{ .Name }}.
func (p *{{ .Name }}) String() string {
	return "{{ .Name }}"
}
