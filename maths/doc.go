// Package maths contains typed functions (e.g. Min, Max) for all built-in
// integer and floating point types. Package math in the standard library only
// provides those functions for float64. To regenerate, run `go generate` on this
// package.
package maths

//go:generate -command genmaths go run generator.go
//go:generate genmaths -clean
//go:generate genmaths -type int
//go:generate genmaths -type int8
//go:generate genmaths -type int16
//go:generate genmaths -type int32
//go:generate genmaths -type int64
//go:generate genmaths -type uint
//go:generate genmaths -type uint8
//go:generate genmaths -type uint16
//go:generate genmaths -type uint32
//go:generate genmaths -type uint64
//go:generate genmaths -type float32
//go:generate genmaths -type float64
//go:generate go fmt
