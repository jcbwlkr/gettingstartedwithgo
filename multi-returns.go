package main

import "fmt"

func main() {
	n := 16
	d := 3
	q, r := divide(n, d) // HL

	fmt.Printf("%d divided by %d is %d with remainder %d", n, d, q, r)
}

// divide accepts a number and divisor and returns the quotient and remainder
func divide(num, div int) (int, int) { // HL
	quot := int(num / div)
	rem := num % div

	return quot, rem // HL
}
