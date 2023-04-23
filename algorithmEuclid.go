package main

import "fmt"

// Алгоритм E (Алгоритм Евклида).
// Даны два целых положительных числа m и n.
// Требуется найти их наибольший общий делитель, т. е. наибольшее целое положительное число, которое нацело делит оба числа m и n.
func main() {
	m := 119
	n := 544

	result := algorithmEuclid(m, n)
	fmt.Println(result)
}

func algorithmEuclid(m, n int) (result int) {
	// E1: ищем остаток
	reminder := m % n
	// E2: сравнение остатка с 0, если reminder == 0, то n искомое число
	if reminder == 0 {
		result = n
	}
	// E3: замещение - m <- n, n <- r
	if reminder != 0 {
		m = n
		n = reminder
		fmt.Println("m:", m, "n:", n)
		result = algorithmEuclid(m, n)
	}

	return
}
