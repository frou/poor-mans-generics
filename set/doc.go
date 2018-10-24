// Package set contains generated set types and their associated methods. A set
// type is currently generated for each of Go's built-in numeric types as well
// as its built-in string type.
//
// Like Go's map type, these set types should not be written to (i.e. the
// Add/Remove/Clear methods) if there are concurrent readers/writers.
//
// To regenerate, run `go generate` on this package.
package set

//go:generate -command genset go run generator.go
//go:generate genset -clean
//go:generate genset -type int
//go:generate genset -type int8
//go:generate genset -type int16
//go:generate genset -type int32
//go:generate genset -type int64
//go:generate genset -type uint
//go:generate genset -type uint8
//go:generate genset -type uint16
//go:generate genset -type uint32
//go:generate genset -type uint64
//go:generate genset -type float32
//go:generate genset -type float64
//go:generate genset -type complex64
//go:generate genset -type complex128
//go:generate genset -type string
//go:generate go fmt
