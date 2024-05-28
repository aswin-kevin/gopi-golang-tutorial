package main

import "fmt"

func main() {
	marks := []int{12}

	marks = append(marks, 20, 30, 40)

	newMarks := []int{50, 60, 70}

	marks = append(marks, newMarks...)

	fmt.Println(marks)

	// using make command

	// realMarks := make([]int, 4, 6)

}
