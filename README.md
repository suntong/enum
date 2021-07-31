
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

	wxType           = enum.NewEnum()
	MsgtypeText      = wxType.IotaAt("文本消息", 1)
	MsgtypeImage     = wxType.IotaAt("图片消息", 3)
	MsgtypeVoice     = wxType.IotaAt("语音消息", 34)
	MsgtypeSharecard = wxType.IotaAt("名片消息", 42)
	MsgtypeVideo     = wxType.IotaAt("视频消息", 43)
	MsgtypeEmoticon  = wxType.IotaAt("表情消息", 47)
	MsgtypeLocation  = wxType.IotaAt("地理位置消息", 48)
	MsgtypeApp       = wxType.IotaAt("APP消息", 49)
	MsgtypeVoipmsg   = wxType.IotaAt("VOIP消息", 50)
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

	fmt.Println("=======")
	for _, wxt := range []int{
		MsgtypeText, MsgtypeImage, MsgtypeVoice, MsgtypeSharecard, MsgtypeVideo,
		MsgtypeEmoticon, MsgtypeLocation, MsgtypeApp, MsgtypeVoipmsg} {
		fmt.Printf("收到%s(type %d)\n", wxType.String(wxt), wxt)
	}
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
	// =======
	// 收到文本消息(type 1)
	// 收到图片消息(type 3)
	// 收到语音消息(type 34)
	// 收到名片消息(type 42)
	// 收到视频消息(type 43)
	// 收到表情消息(type 47)
	// 收到地理位置消息(type 48)
	// 收到APP消息(type 49)
	// 收到VOIP消息(type 50)
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
