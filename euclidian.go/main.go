package main

import "fmt"

// Algorithm E (Euclidean algorithm).
// given 2 positive integers m & n.
// It is required to find their greatest common divisor, that is,
// the largest positive integer that divides both numbers m and n.
func main() {
	m := 119
	n := 544

	result := algorithmEuclid(m, n)
	fmt.Println(result)
}

func algorithmEuclid(m, n int) (result int) {
	if m < n {
		t := m
		m = n
		n = t
	}
	// E1: find reminder
	reminder := m % n
	// E2: comparison of the remainder with 0, if reminder == 0, then n is the desired number
	if reminder == 0 {
		result = n
	}
	// E3: substitution - m <- n, n <- r
	if reminder != 0 {
		m = n
		n = reminder
		fmt.Println("m:", m, "n:", n)
		result = algorithmEuclid(m, n)
	}

	return
}
