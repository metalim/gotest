package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// Устанавливаем точность
	precision := 1000

	// Константы
	c := new(big.Float).SetPrec(uint(precision)).SetFloat64(299792458)
	m_p := new(big.Float).SetPrec(uint(precision)).SetFloat64(1.67262192369e-27)
	E_universe := new(big.Float).SetPrec(uint(precision)).SetFloat64(1e68)

	// Промежуточные вычисления
	mc2 := new(big.Float).Mul(m_p, new(big.Float).Mul(c, c))

	// Шаг 1: Расчет гамма-фактора
	gamma := new(big.Float).Add(new(big.Float).Quo(E_universe, mc2), big.NewFloat(1))

	// Шаг 2: Расчет 1 - v^2/c^2 напрямую
	one := big.NewFloat(1)
	one_minus_v2_c2 := new(big.Float).Quo(one, new(big.Float).Mul(gamma, gamma))

	// Шаг 3: Расчет v/c
	v_c := new(big.Float).Sqrt(new(big.Float).Sub(one, one_minus_v2_c2))

	// Шаг 4: Расчет v
	v := new(big.Float).Mul(v_c, c)

	// Шаг 5: Расчет разницы между c и v
	diff := new(big.Float).Sub(c, v)

	// Вывод результатов
	fmt.Printf("Гамма-фактор: %.10e\n", gamma)
	fmt.Printf("v/c: %.100f\n", v_c)
	fmt.Printf("Скорость протона: %.20f м/с\n", v)
	fmt.Printf("Разница между скоростью света и скоростью протона: %.10e м/с\n", diff)

	// Вычисление log10 разницы
	log10_diff, _ := diff.Float64()
	fmt.Printf("log10 разницы: %.10f\n", math.Log10(log10_diff))
}
