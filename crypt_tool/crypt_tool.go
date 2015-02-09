package main

import (
	"github.com/lixiaolong/golang/bytesbuffer"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
)

func main() {
	var inTE, outTE *walk.TextEdit
	mw := MainWindow{
		Title:   "crypt_tool",
		MinSize: Size{580, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{
						AssignTo: &outTE,
						ReadOnly: true,
						Text: `str -> hex: 1234 -> 31323334
		hex -> str: 31323334 -> 1234
		< - > : 交换两边
		clear : 清除内容`},
				},
			},
			HSplitter{
				MaxSize: Size{50, 30},
				Children: []Widget{
					PushButton{
						Text: "str -> hex",
						OnClicked: func() {
							str := bytesbuffer.NewBufferString(inTE.Text())
							outTE.SetText(strings.TrimSpace(str.ByteString(" ")))
						},
					},
					PushButton{
						Text: "hex -> str",
						OnClicked: func() {
							str := bytesbuffer.NewBufferString("")
							str.WriteByteString(inTE.Text())
							outTE.SetText(str.Buffer.String())
						},
					},
					PushButton{
						Text: "< - >",
						OnClicked: func() {
							str := inTE.Text()
							inTE.SetText(outTE.Text())
							outTE.SetText(str)
						},
					},
					PushButton{
						Text: "clear",
						OnClicked: func() {
							inTE.SetText("")
							outTE.SetText("")
						},
					},
				},
			},
		},
	}
	mw.Run()
}
