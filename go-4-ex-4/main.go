package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*(9.0/5.0) + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5.0 / 9.0)
}

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*(9.0/5.0) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5.0 / 9.0
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(10))
	fmt.Println(convertFahrenheitToCelsius(350))
	// Ergebnisse: 50 & 176.666...

	fmt.Println(convertCelsiusToFahrenheit(176.66666666666666))
	fmt.Println(convertFahrenheitToCelsius(50))

	var cozy Celsius = 23.0
	fmt.Println(cozy.ConvertToFahrenheit())

	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius())
}
