
# enum

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/enum?status.svg)](http://godoc.org/github.com/suntong/enum)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/enum)](https://goreportcard.com/report/github.com/suntong/enum)
[![travis Status](https://travis-ci.org/suntong/enum.svg?branch=master)](https://travis-ci.org/suntong/enum)

## TOC
- [API](#api)
  - [> test/Enum_test.go](#-testenum_testgo)
- [Ref: API, v1](#ref-api-v1)
  - [> Enum_v1_test.go](#-enum_v1_testgo)
- [Author(s) & Contributor(s)](#author(s)-&-contributor(s))

Go Enum and its string representation lib

## API

#### > test/Enum_test.go
```go
////////////////////////////////////////////////////////////////////////////
// Porgram: Enum_test.go
// Purpose: Go Enum and its string representation lib demo
// Authors: Tong Sun (c) 2017-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package test

import (
	"fmt"

	"github.com/suntong/enum"
)

var (
	example = enum.NewEnum()
	Alpha   = example.Iota("Alpha")
	Beta    = example.Iota("Beta")

	weekday = enum.NewEnum()
	Sunday  = weekday.Iota("Sunday")
	Monday  = weekday.Iota("Monday")

	wkday = enum.NewEnum()
	Suday = wkday.Iota("Su")
	Moday = wkday.Iota("Mo")
	Tuday = wkday.Iota("Tu")
	Weday = wkday.Iota("We")
	Thday = wkday.Iota("Th")
	Frday = wkday.Iota("Fr")
	Saday = wkday.Iota("Sa")
)

// for standalone test, change package to main and the next func def to,
// func main() {
func Example_output() {
	fmt.Printf("%s\n", example.String(Alpha))
	fmt.Printf("%s\n", example.String(Beta))
	fmt.Println("=======")
	fmt.Printf("%d\t%d\n", Alpha, Alpha+1)
	fmt.Printf("%s\t%s\n", example.String(Beta-1), example.String(Alpha+1))
	fmt.Println("=======")
	if a, ok := example.Get("Alpha"); ok {
		fmt.Printf("%d: %s\n", a, example.String(a))
	}
	if b, ok := example.Get("Beta"); ok {
		fmt.Printf("%d: %s\n", b, example.String(b))
	}

	fmt.Printf("%d:%s\n", Sunday, weekday.String(Sunday))
	fmt.Printf("%d:%s\n", Monday, weekday.String(Monday))
	fmt.Println("=======")

	fmt.Printf("%s: %d\n", wkday.String(Suday), Suday)
	fmt.Printf("%s: %d\n", wkday.String(Moday), Moday)
	fmt.Printf("%s: %d\n", wkday.String(Moday+Thday), Frday)
	fmt.Printf("%s: %d\n", wkday.String(Saday), Weday*2)
	// Output:
	// Alpha
	// Beta
	// =======
	// 0	1
	// Alpha	Beta
	// =======
	// 0: Alpha
	// 1: Beta
	// 0:Sunday
	// 1:Monday
	// =======
	// Su: 0
	// Mo: 1
	// Fr: 5
	// Sa: 6
}
```

## Ref: API, v1

#### > Enum_v1_test.go
```go
////////////////////////////////////////////////////////////////////////////
// Porgram: Enum_test.go
// Purpose: Go Enum and its string representation lib demo
// Authors: Tong Sun (c) 2017-2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package enum_test

import (
	"fmt"

	"gopkg.in/suntong/enum.v1"
)

var (
	// Note this is a v1 example, i.e.,
	// the enum variables are declared,
	// instead of defined with `enum.NewEnum()`
	example enum.Enum
	Alpha   = example.Iota("Alpha")
	Beta    = example.Iota("Beta")

	weekday enum.Enum
	Sunday  = weekday.Iota("Sunday")
	Monday  = weekday.Iota("Monday")

	wkday enum.Enum
	Suday = wkday.Iota("Su")
	Moday = wkday.Iota("Mo")
	Tuday = wkday.Iota("Tu")
	Weday = wkday.Iota("We")
	Thday = wkday.Iota("Th")
	Frday = wkday.Iota("Fr")
	Saday = wkday.Iota("Sa")
)

// for standalone test, change package to main and the next func def to,
// func main() {
func Example_output() {
	fmt.Printf("%s\n", example.String(Alpha))
	fmt.Printf("%s\n", example.String(Beta))
	fmt.Println("=======")
	fmt.Printf("%d\t%d\n", Alpha, Alpha+1)
	fmt.Printf("%s\t%s\n", example.String(Beta-1), example.String(Alpha+1))
	fmt.Println("=======")
	if a, ok := example.Get("Alpha"); ok {
		fmt.Printf("%d: %s\n", a, example.String(a))
	}
	if b, ok := example.Get("Beta"); ok {
		fmt.Printf("%d: %s\n", b, example.String(b))
	}

	fmt.Printf("%d:%s\n", Sunday, weekday.String(Sunday))
	fmt.Printf("%d:%s\n", Monday, weekday.String(Monday))
	fmt.Println("=======")

	fmt.Printf("%s: %d\n", wkday.String(Suday), Suday)
	fmt.Printf("%s: %d\n", wkday.String(Moday), Moday)
	fmt.Printf("%s: %d\n", wkday.String(Moday+Thday), Frday)
	fmt.Printf("%s: %d\n", wkday.String(Saday), Weday*2)
	// Output:
	// Alpha
	// Beta
	// =======
	// 0	1
	// Alpha	Beta
	// =======
	// 0: Alpha
	// 1: Beta
	// 0:Sunday
	// 1:Monday
	// =======
	// Su: 0
	// Mo: 1
	// Fr: 5
	// Sa: 6
}
```

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome.
