package fibonacci

import (
	"log"
)

// Fibonacci function return the value of fibonacci number for given index
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// PrintFibonacci print fibonacci number in sequence
func PrintFibonacci(n int64) {
	var (
		i, t1, t2, tmp int64
	)

	t1 = 0
	t2 = 1

	for i = 0; i < n; i++ {
		log.Println(t1)
		tmp = t1 + t2
		t1 = t2
		t2 = tmp
	}
}
