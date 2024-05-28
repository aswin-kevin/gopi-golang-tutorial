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
	var allStudentsData = make([]StudentSchema, 0)

	var studentNames = []string{"John", "Jane", "Doe"}

	for _, studentName := range studentNames {
		var student = StudentSchema{
			Name: studentName,
		}
		allStudentsData = append(allStudentsData, student)
	}

	fmt.Println(allStudentsData)
}
