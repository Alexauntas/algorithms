package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("no args")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fibonacciNumbers(n))
}

func fibonacciNumbers(n int) []int {
	result := make([]int, 0, n)
	for i := range n {
		result = append(result, step(i, nil))
	}
	return result
}

func fibonacciNumbersWithCache(n int) []int {
	result := make([]int, 0, n)
	cache := make(cache, n)
	for i := range n {
		result = append(result, step(i, cache))
	}
	return result
}

func step(n int, c cache) int {
	if n <= 1 {
		return 1
	}
	if c == nil {
		return step(n-1, nil) + step(n-2, nil)
	}
	value, exists := c.get(n)
	if exists {
		return value
	}
	value = step(n-1, c) + step(n-2, c)
	c.set(n, value)
	return value
}

type cache map[int]int

func (c cache) get(n int) (int, bool) {
	value, exists := c[n]
	return value, exists
}

func (c cache) set(n int, value int) {
	c[n] = value
	return
}
