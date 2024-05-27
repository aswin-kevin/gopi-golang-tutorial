package main

import "fmt"

// variable , constant

var age int // decraling a variable

func main() {
	age = 20 // assigning a value to a variable

	// var userId int = 12
	userId := 12 // short hand declaration and assingment

	userId = 13 // reassigning a value to a variable

	fmt.Println("Hello, My age is :", age)
	fmt.Println("Hello, My user id is :", userId)
}
