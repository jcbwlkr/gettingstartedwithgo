package main

import (
	"fmt"
	"sync" // HL
	"time"
)

func main() {
	wg := &sync.WaitGroup{} // HL
	wg.Add(2)               // HL
	go printEvery("Foo", 500*time.Millisecond, wg)
	go printEvery("Bar", 250*time.Millisecond, wg)
	wg.Wait() // HL
}

func printEvery(s string, d time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // HL

	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(d)
	}
}
