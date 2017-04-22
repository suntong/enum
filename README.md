
# enum

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/enum?status.svg)](http://godoc.org/github.com/suntong/enum)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/enum)](https://goreportcard.com/report/github.com/suntong/enum)
[![travis Status](https://travis-ci.org/suntong/enum.svg?branch=master)](https://travis-ci.org/suntong/enum)

## TOC
- [API](#api)
  - [> Enum_test.go](#-enum_testgo)

Go Enum and its string representation lib

# API

#### > Enum_test.go
```go
////////////////////////////////////////////////////////////////////////////
// Porgram: Enum_test.go
// Purpose: Go Enum and its string representation lib demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Egon https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ
////////////////////////////////////////////////////////////////////////////

package enum_test

import (
	"fmt"

	"github.com/suntong/enum"
)

var (
	Alpha = enum.Ciota("Alpha")
	Beta  = enum.Ciota("Beta")
)

type Example struct {
	enum.Enum
}

// for standalone test, change package to main and the next func def to,
// func main() {
func Example_output() {
	fmt.Printf("%+v\n", Example{Alpha})
	fmt.Printf("%+v\n", Example{Beta})
	fmt.Println("=======")
	fmt.Printf("%d\t%d\n", Alpha, Alpha+1)
	fmt.Printf("%+v\t%+v\n", Example{Beta - 1}, Example{Alpha + 1})
	fmt.Println("=======")
	if a, ok := enum.Get("Alpha"); ok {
		fmt.Printf("%d\n", a)
	}
	if b, ok := enum.Get("Beta"); ok {
		fmt.Printf("%d: %+v\n", b, Example{b})
	}

	// Output:
	// Alpha
	// Beta
	// =======
	// 0	1
	// Alpha	Beta
	// =======
	// 0
	// 1: Beta
}
```

All patches welcome.
