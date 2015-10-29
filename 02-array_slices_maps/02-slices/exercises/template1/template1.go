// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/mPKmyGNR2L

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var slice []int

	// Appends numbers to the slice.
	for i, _ := range [5]int{} {
		slice = append(slice, i)
	}

	// Display each value in the slice.
	for _, num := range slice {
		fmt.Println(num)
	}

	// Declare a slice of strings and populate the slice with names.
	slice2 := []string{"Mathieu", "John", "Khaled"}

	// Display each index position and slice value.
	for i, name := range slice2 {
		fmt.Printf("Position %v = %v\n", i, name)
	}

	// // Take a slice of index 1 and 2 of the slice of strings.
	// Display each index position and slice values for the new slice.
	for i, name := range slice2[1:3] {
		fmt.Printf("Position %v = %v\n", i, name)
	}
}
