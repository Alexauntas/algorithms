package main

import (
	"testing"
)

func BenchmarkFibonacciNumbers(b *testing.B) {
	benchmarks := []struct {
		name    string
		fibFunc func(int) []int
		n       int
	}{
		{"Slow_N10", fibonacciNumbers, 10},
		{"Slow_N30", fibonacciNumbers, 30},
		{"Slow_N50", fibonacciNumbers, 50},

		{"Fast_N10", fibonacciNumbersWithCache, 10},
		{"Fast_N30", fibonacciNumbersWithCache, 30},
		{"Fast_N50", fibonacciNumbersWithCache, 50},
		{"Fast_N1000", fibonacciNumbersWithCache, 1000},
	}

	for i := range benchmarks {
		b.Run(benchmarks[i].name, func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				benchmarks[i].fibFunc(benchmarks[i].n)
			}
		})
	}
}
