package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем новый big.Float без явного указания precision
	f := new(big.Float)
	fmt.Printf("Дефолтная precision для нового big.Float: %d бит\n", f.Prec())

	// Создаем big.Float из float64
	f64 := big.NewFloat(1.23)
	fmt.Printf("Precision для big.Float, созданного из float64: %d бит\n", f64.Prec())

	// Создаем big.Float с явно указанной precision
	f128 := new(big.Float).SetPrec(128)
	fmt.Printf("Precision для big.Float с явно указанной precision: %d бит\n", f128.Prec())

	// Выполняем операцию между f64 и f128
	result := new(big.Float).Add(f64, f128)
	fmt.Printf("Precision результата операции между float64 и float128: %d бит\n", result.Prec())
}
