// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/-WjYGqauiY

// Declare a named type called counter with a base type of int. Declare a variable
// named c of type counter set to its zero value. Display the value of c.
//
// Declare a variable named c2 of type counter set to the value of 10. Display the value
// of c2.
//
// Declare a variable named i of the base type int. Attempt to assign the value
// of i to c. Does the compiler allow the assignment?
package main

// Add imports.
import "fmt"

// Declare the named type called counter with a base
// type of int.
type counter int

// main is the entry point for the application.
func main() {
	// Declare a variable named c of type counter
	// set to its zero value.
	var c counter
	c = 0

	// Display the value of c.
	fmt.Println(c)

	// Declare a variable named c2 of type counter
	// set to the value of 10.
	var c2 counter
	c = 10

	// Display the value of c2.
	fmt.Println(c2)

	// Declare a variable named i of the base type int.
	var i int

	// Assign the value of i to the variable named c.
	c = i

	// Does it compile?
}
