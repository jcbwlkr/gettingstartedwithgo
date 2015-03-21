package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	input := "The rain in Spain falls mainly on the plains"

	re, err := regexp.Compile("[rm[ain") // HL
	if err != nil {
		log.Fatalln(err)
	}

	matches := re.FindAllString(input, -1)

	fmt.Println(matches)
}
