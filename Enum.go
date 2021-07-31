////////////////////////////////////////////////////////////////////////////
// Package: Enum
// Purpose: Go Enum and its string representation lib
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ
////////////////////////////////////////////////////////////////////////////

package enum

// Enum holds the enumerables
type Enum struct {
	enums map[int]string
	iota  int
}

// type Enum int

func NewEnum() Enum {
	return Enum{make(map[int]string), 0}
}

// String turns Enum to its string form
func (e Enum) String(v int) string {
	return e.enums[v]
}

// Exist returns Enum's string form if exist
func (e Enum) Exist(v int) (string, bool) {
	elem, ok := e.enums[v]
	return elem, ok
}

// Iota converts string to enumerable, similar to Go's iota
func (e *Enum) Iota(s string) int {
	r := e.iota
	e.iota++
	e.enums[r] = s
	return r
}

// IotaAt converts string to the enumerable given as the at value
func (e *Enum) IotaAt(s string, at int) int {
	e.iota = at
	return e.Iota(s)
}

// Get lookup the given string for interal enumerable
func (e Enum) Get(s string) (int, bool) {
	for ii, vv := range e.enums {
		if vv == s {
			return ii, true
		}
	}
	return -1, false
}
