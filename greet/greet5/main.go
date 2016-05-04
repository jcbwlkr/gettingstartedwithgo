package main

import (
	"flag"
	"fmt"
	"math/rand"
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

	tmpls := []string{
		"Howdy there, %s",
		"Salutations %s. How are you today?",
		"Avast it be the scoundrel %s",
		"Oh hai %s (^_^)",
		"Hello, %s",
	}
	rand.Seed(time.Now().Unix())
	t := tmpls[rand.Intn(len(tmpls))]

	msg := fmt.Sprintf(t, name)

	if yell {
		msg = strings.ToUpper(msg)
	}

	fmt.Println(msg)
}
