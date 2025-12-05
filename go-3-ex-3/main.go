package main

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i <= Upper; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			println("FizzBuzz")
		case i%3 == 0:
			println("Fizz")
		case i%5 == 0:
			println("Buzz")
		default:
			println(i)
		}
	}
}
