package main

// import format from the go standard library
import "fmt"

// write the main package main function
func main() {
	// create a variable with a 10 integer array type specification
	var arr [10]int

	// print the array to the console before adding elements to the array
	// output: [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(arr)

	// create a for loop for 10 iterations
	for i := 0; i < 10; i++ {
		// every iteration set arr at index i to the current i value which will be the current iteration
		// number
		arr[i] = i
	}

	// print arr again to the console
	// output: [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(arr)

	// set a variable called count to the length of arr
	count := len(arr)

	// print to the console a descriptive string and the count variable value
	fmt.Println("Length of array", count)
}

/*
[0 0 0 0 0 0 0 0 0 0]
[0 1 2 3 4 5 6 7 8 9]
Length of array 10
*/
