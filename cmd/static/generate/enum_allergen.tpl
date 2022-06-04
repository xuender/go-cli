package {{ .Package }}

// {{ .Name }} TODO.
type {{ .Name }} int32

const (
	// One TODO.
	One {{ .Name }} = 1 << iota
	// Tow TODO.
	Two
	// Three TODO.
	Three
)

var (
	{{ .Name }}Names = map[int32]string{
		1 << 0: "one",
		1 << 1: "two",
		1 << 2: "three",
		// TODO
	}
	{{ .Name }}Values = map[string]int32{
		"one":   1 << 0,
		"two":   1 << 1,
		"three": 1 << 2,
		// TODO
	}
)

// {{ .Name }}Merge merage.
func {{ .Name }}Merge(values ...{{ .Name }}) {{ .Name }} {
	var merge {{ .Name }}

	for _, value := range values {
		merge |= value
	}

	return merge
}

// {{ .Name }}Split split.
func {{ .Name }}Split(value {{ .Name }}) []{{ .Name }} {
	slice := []{{ .Name }}{}
	i32 := int32(value)

	for key := range {{ .Name }}Names {
		if i32&key > 0 {
			slice = append(slice, {{ .Name }}(key))
		}
	}

	return slice
}
