syntax = "proto3";
package {{ .Package }};
option go_package = "./{{ .Path }}";

// {{ .Name }} TODO.
message {{ .Name }} {
	// Id TODO.
  uint64 id = 1;
	// Name TODO.
  string name = 2;
}
