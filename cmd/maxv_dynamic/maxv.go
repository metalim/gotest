package main

import (
	"fmt"
	"math"
	"math/big"
)

func calculateWithPrecision(precision uint) (*big.Float, *big.Float, *big.Float) {
	c := new(big.Float).SetPrec(precision).SetFloat64(299792458)
	m_p := new(big.Float).SetPrec(precision).SetFloat64(1.67262192369e-27)
	E_universe := new(big.Float).SetPrec(precision).SetFloat64(1e68)

	mc2 := new(big.Float).Mul(m_p, new(big.Float).Mul(c, c))
	gamma := new(big.Float).Add(new(big.Float).Quo(E_universe, mc2), big.NewFloat(1))

	one := big.NewFloat(1)
	one_minus_v2_c2 := new(big.Float).Quo(one, new(big.Float).Mul(gamma, gamma))

	v_c := new(big.Float).Sqrt(new(big.Float).Sub(one, one_minus_v2_c2))

	v := new(big.Float).Mul(v_c, c)

	diff := new(big.Float).Sub(c, v)

	return v, v_c, diff
}

func adaptivePrecisionCalculation() {
	initialPrecision := uint(64)
	maxPrecision := uint(2048)
	stepPrecision := uint(64)

	var significantResult bool
	var finalPrecision uint

	for precision := initialPrecision; precision <= maxPrecision; precision += stepPrecision {
		v, v_c, diff := calculateWithPrecision(precision)

		fmt.Printf("Precision: %d bits\n", precision)
		fmt.Printf("v: %.20f м/с\n", v)
		fmt.Printf("v/c: %.50f\n", v_c)
		fmt.Printf("Разница между скоростью света и скоростью протона: %.50e м/с\n", diff)

		if diff.Sign() > 0 {
			log10_diff, _ := diff.Float64()
			if log10_diff != 0 {
				fmt.Printf("log10 разницы: %.10f\n", math.Log10(log10_diff))
				significantResult = true
				finalPrecision = precision
				break
			}
		}

		fmt.Println()
	}

	if significantResult {
		fmt.Printf("\nПолучен значимый результат при точности %d бит\n", finalPrecision)
	} else {
		fmt.Printf("\nНе удалось получить значимый результат даже при максимальной точности %d бит\n", maxPrecision)
	}
}

func main() {
	adaptivePrecisionCalculation()
}
