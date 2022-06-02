syntax = "proto3";
package {{ .Package }};
option go_package = "./{{ .Path }}";

// {{ .Name }} TODO.
enum {{ .Name }} {
	// One TODO.
  one = 0;
	// Two TODO.
  two = 1;
	// Three TODO.
  three = 1;
}
