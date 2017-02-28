// Package set contains generated set types and associated methods.
//
// A set type is currently generated for Go's built-in string type and each of
// its built-in numeric types.
//
// Run the following command to regenerate:
//
//      go generate github.com/frou/stdext/set
package set

//go:generate -command set_type go run generator.go -type

//go:generate set_type string
//go:generate set_type uint8
//go:generate set_type uint16
//go:generate set_type uint32
//go:generate set_type uint64
//go:generate set_type int8
//go:generate set_type int16
//go:generate set_type int32
//go:generate set_type int64
//go:generate set_type float32
//go:generate set_type float64
//go:generate set_type complex64
//go:generate set_type complex128
