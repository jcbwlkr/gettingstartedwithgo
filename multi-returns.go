package main

import "fmt"

func main() {
	n := 16
	d := 3
	q, r := divide(n, d)

	fmt.Printf("%d divided by %d is %d with remainder %d", n, d, q, r)
}

func divide(num, div int) (int, int) {
	return int(num / div), num % div
}
