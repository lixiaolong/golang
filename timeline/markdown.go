package main

import (
	"html"
	//"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func Markdown(s string) string {
	s = html.EscapeString(s)
	output := blackfriday.MarkdownBasic([]byte(s))
	return string(output)
}
