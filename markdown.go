package main

import (
	"fmt"

	"github.com/russross/blackfriday"
)

func main() {
	in := `
# Title
## Subtitle

This is my content
`
	out := blackfriday.MarkdownCommon([]byte(in))

	fmt.Println(string(out))
}
