package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var yell bool
	flag.BoolVar(&yell, "yell", false, "Provide -yell to make greet louder")
	flag.Parse()

	name := "Jacob"
	if arg := flag.Arg(0); arg != "" {
		name = arg
	}

	msg := fmt.Sprintf("Hello, %s", name)

	if yell {
		msg = strings.ToUpper(msg)
	}

	fmt.Println(msg)
}
