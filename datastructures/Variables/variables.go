package main

import "fmt"

func one() {
	// Use the var keyword to define a variable and set it's data type
	var variOne int
	var variTwo int
	// Set values to the variables
	variOne = 1
	variTwo = 2
	// print both variable values to the console
	fmt.Println("valueOne:", variOne)
	fmt.Println("valueTwo:", variTwo)
}

func two() {
	// Use the var key word to define two variables and their data type in one line of code
	var variOne, variTwo int
	// set a value only to the first declared variable
	variOne = 3
	// print both variable values to the console without assigning a value to the second variable
	fmt.Println("valueOne:", variOne)
	fmt.Println("valueTwo:", variTwo)
}

func three() {
	// define two variables and set a value to them without specifying the data type allowed to be
	// assigned to them. Let the go compliler infer their types
	variOne := 4
	variTwo := 5
	// print both variable values to the console
	fmt.Println("valueOne:", variOne)
	fmt.Println("valueTwo:", variTwo)
}

// const PI = 3.14

// Declare global variables
var globalVari int = 8

func four() {
	// Declare two variables with the shorthand syntax and let go infer their types
	variOne := 6
	variTwo := 7
	// print both function scoped variables and the globally scoped variable
	fmt.Println("valueOne:", variOne)
	fmt.Println("valueTwo:", variTwo)
	fmt.Println("valueGlobal:", globalVari)
}

func main() {
	one()
	two()
	three()
	four()
}
