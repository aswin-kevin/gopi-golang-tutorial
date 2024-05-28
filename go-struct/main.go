package main

import "fmt"

type StudentSchema struct {
	Name       string
	Std        int
	Address    string
	Phone      []string
	IsHosteler bool
	Subjects   map[string]float32
}

func main() {
	// var age int
	var Student StudentSchema

	Student.Name = "John"
	Student.Std = 10
	fmt.Println(Student)

	var Student2 StudentSchema

	Student2.Name = "Jane"

	var Student3 = StudentSchema{
		Name: "Doe",
		Std:  12,
	}

	fmt.Println(Student3)
}
