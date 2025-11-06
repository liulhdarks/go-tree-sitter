package powershell

// #cgo CFLAGS: -std=c11 -fPIC
// #include "parser.h"
// TSLanguage *tree_sitter_powershell();
import "C"

import "unsafe"

// Get the tree-sitter Language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_powershell())
}
