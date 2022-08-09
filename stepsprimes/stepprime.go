package stepsprimes

import (
	"math"
)

func Gap(g, m, n int) []int {
	var numbers []int
	for i := m; i < n; i++ {
		if isPrime(m, i) && isPrime(m+g, i) && len(numbers) < 2 {
			numbers = append(numbers, m, m+g)
		}
		m += 1
	}
	return numbers
}

func isPrime(num, i int) bool {
	for i := 2; i <= int(math.Floor(float64(num)/2)); i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}
