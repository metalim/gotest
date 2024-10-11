package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	n := 100_000
	primes := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			primes++
		}
	}
	duration := time.Since(start)
	fmt.Printf("Primes: %d\n", primes)
	fmt.Printf("Go Time: %f seconds\n", duration.Seconds())
}
