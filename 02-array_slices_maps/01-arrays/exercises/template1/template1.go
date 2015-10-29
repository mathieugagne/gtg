// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/qKUNW0FSgC

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

// Add imports.
import "fmt"

// main is the entry point for the application.
func main() {
	// Declare an array of 5 strings set to its zero value.
	var strings [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	strings2 := [5]string{"Foo", "Bar", "FooBar", "Blah", "Blew"}

	// Assign the populated array to the array of zero values.
	strings = strings2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for _, s := range strings {
		fmt.Println(s, &s)
	}
}
