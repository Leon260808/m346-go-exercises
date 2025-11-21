package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  int16
}

type Profile struct {
	Name             FullName
	Born             BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Leon",
			LastName:  "Tran",
		},
		Born: BirthDate{
			Day:   26,
			Month: 8,
			Year:  2008,
		},
		NumberOfSiblings: 3,
		ZodiacSign:       'V',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
