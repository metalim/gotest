package main

import (
	"fmt"
	"math"
)

const (
	G               = 6.67430e-11
	c               = 299792458
	M_sun           = 1.989e30
	M_a_star        = 4.1e6 * M_sun
	R_schwarzschild = 2 * G * M_a_star / c / c
	R_isco          = 3 * R_schwarzschild
)

func main() {
	fmt.Printf("M_sun = %e kg\n", M_sun)
	fmt.Printf("M_a_star = %e kg = %.1f M_sun\n", M_a_star, M_a_star/M_sun)
	fmt.Printf("R_schwarzschild = %e m = %.1f km\n", R_schwarzschild, R_schwarzschild/1000)
	fmt.Printf("R_isco = %e m = %.1f km\n", R_isco, R_isco/1000)

	T_isco := 2 * math.Pi * math.Sqrt(R_isco*R_isco*R_isco/G/M_a_star)
	fmt.Printf("T_isco = %.1f s = %.1f min\n", T_isco, T_isco/60)
}
