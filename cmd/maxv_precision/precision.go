package main

import (
	"fmt"
	"math"
	"math/big"
)

func calculateAndPrint(precision uint) {
	fmt.Printf("\nРезультаты для точности %d бит:\n", precision)

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

	fmt.Printf("Гамма-фактор: %.10e\n", gamma)
	fmt.Printf("v/c: %.100f\n", v_c)
	fmt.Printf("Скорость протона: %.20f м/с\n", v)
	fmt.Printf("Разница между скоростью света и скоростью протона: %.10e м/с\n", diff)

	log10_diff, _ := diff.Float64()
	fmt.Printf("log10 разницы: %.10f\n", math.Log10(log10_diff))
}

func main() {
	precisions := []uint{64, 128, 256, 512, 1000}
	for _, prec := range precisions {
		calculateAndPrint(prec)
	}
}
