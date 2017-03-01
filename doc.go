// Package set contains generated set types and their associated methods. A set
// type is currently generated for each of Go's built-in numeric types as well
// as its built-in string type. To regenerate, run `go generate` on this
// package.
package set

//go:generate -command set_type go run generator.go
//go:generate set_type -clean
//go:generate set_type -type uint8
//go:generate set_type -type uint16
//go:generate set_type -type uint32
//go:generate set_type -type uint64
//go:generate set_type -type int8
//go:generate set_type -type int16
//go:generate set_type -type int32
//go:generate set_type -type int64
//go:generate set_type -type float32
//go:generate set_type -type float64
//go:generate set_type -type complex64
//go:generate set_type -type complex128
//go:generate set_type -type string
//go:generate go fmt
