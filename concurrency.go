package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency!")

	go printEvery("In goroutine", 500*time.Millisecond)
	printEvery("In main", 250*time.Millisecond)
}

func printEvery(s string, d time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(d)
	}
}
