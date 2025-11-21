package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Leon"
	var lastName string = "Tran"
	var dayOfBirth int = 26
	var monthOfBirth int = 8
	var yearOfBirth int = 2008
	var numberOfSiblings int = 3
	var heightInMeters float32 = 1.72
	var zodiacSign byte = 'V'
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
