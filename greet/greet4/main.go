package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	var yell bool
	flag.BoolVar(&yell, "yell", false, "Provide -yell to make greet louder")
	flag.Parse()

	name := "Jacob"
	if arg := flag.Arg(0); arg != "" {
		name = arg
	}

	var msg string

	switch time.Now().Second() {
	case 0, 1:
		msg = fmt.Sprintf("Howdy there, %s", name)
	case 2, 3:
		msg = fmt.Sprintf("Salutations %s. How are you today?", name)
	case 4, 5:
		msg = fmt.Sprintf("Avast it be the scoundrel %s", name)
	case 6, 7:
		msg = fmt.Sprintf("Oh hai %s (^_^)", name)
	default:
		msg = fmt.Sprintf("Hello, %s", name)
	}

	if yell {
		msg = strings.ToUpper(msg)
	}

	fmt.Println(msg)
}
