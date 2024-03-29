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
	MsgtypeSys       = wxType.IotaAt("系统消息", 10000)
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
		fmt.Printf("收到一条%s(type %d)\n", wxType.String(wxt), wxt)
	}
	fmt.Println("=======")
	for _, wxt := range []int{10000, 51} {
		wxtstr, ok := wxType.Exist(wxt)
		if ok {
			fmt.Printf("收到一条%s(type %d)\n", wxtstr, wxt)
		} else {
			fmt.Printf("收到一条未知消息(type %d):%s\n", wxt, wxtstr)
		}
	}
	fmt.Println("=======")
	for _, wxt := range []int{10000, 51} {
		wxtstr, ok := wxType.Exist(wxt)
		if !ok {
			wxtstr = "未知消息"
		}
		fmt.Printf("收到一条%s(type %d)\n", wxtstr, wxt)
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
	// 收到一条文本消息(type 1)
	// 收到一条图片消息(type 3)
	// 收到一条语音消息(type 34)
	// 收到一条名片消息(type 42)
	// 收到一条视频消息(type 43)
	// 收到一条表情消息(type 47)
	// 收到一条地理位置消息(type 48)
	// 收到一条APP消息(type 49)
	// 收到一条VOIP消息(type 50)
	// =======
	// 收到一条系统消息(type 10000)
	// 收到一条未知消息(type 51):
	// =======
	// 收到一条系统消息(type 10000)
	// 收到一条未知消息(type 51)
}
