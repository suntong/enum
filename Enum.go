////////////////////////////////////////////////////////////////////////////
// Package: Enum
// Purpose: Go Enum and its string representation lib
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Egon https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ
////////////////////////////////////////////////////////////////////////////

package enum

var enums []string

// The Enum is the type to hold enumerables
type Enum int

// String turns Enum to its string form
func (e Enum) String() string {
	return enums[int(e)]
}

// Ciota converts string to enumerable, similar to Go's iota
func Ciota(s string) Enum {
	enums = append(enums, s)
	return Enum(len(enums) - 1)
}

// Get lookup the given string for interal enumerable
func Get(s string) (Enum, bool) {
	for ii, vv := range enums {
		if vv == s {
			return Enum(ii), true
		}
	}
	return -1, false
}
