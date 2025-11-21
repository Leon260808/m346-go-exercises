package main

import "fmt"

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}
	type Class struct {
		Students []Student
	}
	modules := map[int]Class{
		1: {
			Students: []Student{
				{FirstName: "Leon", LastName: "Tran"},
				{FirstName: "Alina", LastName: "Fischer"},
			},
		},
		2: {
			Students: []Student{
				{FirstName: "Nevin", LastName: "Wyss"},
				{FirstName: "Max", LastName: "Mustermann"},
			},
		},
	}
	fmt.Println(modules)
}
