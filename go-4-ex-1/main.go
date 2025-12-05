package main

import "fmt"

func computeGrade(gotPoints float32, maxPoints float32) (float32, error) {
	if gotPoints <= 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0.0, fmt.Errorf("cannot compute grade with invalid points")
	}
	return (gotPoints/maxPoints)*5 + 1, nil
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0))
	fmt.Println(computeGrade(-1, 28.0))
	fmt.Println(computeGrade(17.5, -200))
	fmt.Println(computeGrade(30.0, 28.0))
}
