package stepsprimes

import (
	"math"
)

func StepsPrimes(step, start, end int) []int{
	var numbers []int
	for i := start; i < end; i++ {
		if isPrime(start, i) && isPrime(start+step, i) && len(numbers) < 2{
			numbers = append(numbers, start, start+step)
		}
			start += 1
	}
	return numbers
}

func isPrime(num, i int) bool{
	for i := 2; i <= int(math.Floor(float64(num) / 2)); i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}
