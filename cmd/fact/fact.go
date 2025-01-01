package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(fact(2025))
}

func fact(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
