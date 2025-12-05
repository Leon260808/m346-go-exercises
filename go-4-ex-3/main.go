package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	var solutions []float64
	D := computeDiscrimnant(a, b, c)
	switch {
	case D > 0:
		solutions = append(solutions, (-b+math.Sqrt(D))/(2*a))
		solutions = append(solutions, (-b-math.Sqrt(D))/(2*a))

	case D == 0:
		solutions = append(solutions, -b/(2*a))

	case D < 0:
	}
	return solutions
}

func computeDiscrimnant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1))
	fmt.Println(computeQuadraticFormula(2, 4, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))
}
