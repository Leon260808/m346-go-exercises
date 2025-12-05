package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}
func main() {
	fmt.Println(computeHypotenuse(3, 4))
	// Ergebnis = 5
	s := ShortSides{a: 3, b: 4}
	fmt.Println(s.Hypotenuse())
	// Ergebnis = 5
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}
