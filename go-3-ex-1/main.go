package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'
	if p.BirthDate.Day >= 21 && p.BirthDate.Month == 3 || p.BirthDate.Day <= 20 && p.BirthDate.Month == 4 {
		zodiacSign = Aries
	} else if p.BirthDate.Day >= 21 && p.BirthDate.Month == 4 || p.BirthDate.Day <= 21 && p.BirthDate.Month == 5 {
		zodiacSign = Taurus
	} else if p.BirthDate.Day >= 22 && p.BirthDate.Month == 5 || p.BirthDate.Day <= 21 && p.BirthDate.Month == 6 {
		zodiacSign = Gemini
	} else if p.BirthDate.Day >= 22 && p.BirthDate.Month == 6 || p.BirthDate.Day <= 22 && p.BirthDate.Month == 7 {
		zodiacSign = Cancer
	} else if p.BirthDate.Day >= 23 && p.BirthDate.Month == 7 || p.BirthDate.Day <= 22 && p.BirthDate.Month == 8 {
		zodiacSign = Leo
	} else if p.BirthDate.Day >= 23 && p.BirthDate.Month == 8 || p.BirthDate.Day <= 22 && p.BirthDate.Month == 9 {
		zodiacSign = Virgo
	} else if p.BirthDate.Day >= 23 && p.BirthDate.Month == 9 || p.BirthDate.Day <= 22 && p.BirthDate.Month == 10 {
		zodiacSign = Libra
	} else if p.BirthDate.Day >= 23 && p.BirthDate.Month == 10 || p.BirthDate.Day <= 22 && p.BirthDate.Month == 11 {
		zodiacSign = Scorpius
	} else if p.BirthDate.Day >= 23 && p.BirthDate.Month == 11 || p.BirthDate.Day <= 20 && p.BirthDate.Month == 12 {
		zodiacSign = Sagittarius
	} else if p.BirthDate.Day >= 21 && p.BirthDate.Month == 12 || p.BirthDate.Day <= 19 && p.BirthDate.Month == 1 {
		zodiacSign = Capricornus
	} else if p.BirthDate.Day >= 20 && p.BirthDate.Month == 1 || p.BirthDate.Day <= 18 && p.BirthDate.Month == 2 {
		zodiacSign = Aquarius
	} else if p.BirthDate.Day >= 19 && p.BirthDate.Month == 2 || p.BirthDate.Day <= 20 && p.BirthDate.Month == 3 {
		zodiacSign = Pisces
	} else {
		zodiacSign = '?'
	}
	// NOTE: The runes are defined above as constants.

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)

	leon := Person{FullName{"Leon", "Tran"}, BirthDate{26, 8, 2008}}
	alina := Person{FullName{"Alina", "Fischer"}, BirthDate{20, 11, 2007}}
	sophie := Person{FullName{"Sophie", "Tran"}, BirthDate{31, 5, 2010}}

	outputWithZodiacSign(leon)
	outputWithZodiacSign(alina)
	outputWithZodiacSign(sophie)
}
