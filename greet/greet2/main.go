package main

import (
	"flag"
	"fmt"
)

func main() {
	var yell bool
	flag.BoolVar(&yell, "yell", false, "Provide -yell to make greet louder")
	flag.Parse()

	if yell {
		fmt.Println("HELLO, JACOB")
	} else {
		fmt.Println("Hello, Jacob")
	}
}
