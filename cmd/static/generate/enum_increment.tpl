package {{ .Package }}

// {{ .Name }} TODO.
type {{ .Name }} int32

const (
	// One TODO.
	One {{ .Name }} = iota
	// Tow TODO.
	Two
	// Three TODO.
	Three
)

var (
	{{ .Name }}Names = map[int32]string{
		0: "one",
		1: "two",
		2: "three",
		// TODO
	}
	{{ .Name }}Values = map[string]int32{
		"one":   0,
		"two":   1,
		"three": 2,
		// TODO
	}
)
