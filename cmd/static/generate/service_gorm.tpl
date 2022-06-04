package {{ .Package }}

import "gorm.io/gorm"

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

// Method TODO.
func (p *{{ .Name }}) Method(txdb *gorm.DB) {
	// TODO method
}
