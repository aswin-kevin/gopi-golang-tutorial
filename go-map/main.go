package main

import "fmt"

func main() {
	// var students = make(map[string]int)

	var students = map[string]int{
		"john": 90,
		"jane": 95,
	}

	students["john"] = 92
	students["jane2"] = 96

	delete(students, "john")

	fmt.Println(students)
}
