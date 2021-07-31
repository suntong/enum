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
