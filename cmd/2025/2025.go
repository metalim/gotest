package main

import "fmt"

func main() {
	fmt.Printf("45 * 45 = %d\n", 45*45)
	var sum int
	for i := 1; i < 10; i++ {
		sum += i * i * i
	}
	fmt.Printf("sum((1..9) ** 3) = %d\n", sum)
}
