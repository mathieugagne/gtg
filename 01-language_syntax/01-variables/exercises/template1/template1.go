// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/1xUWjHMB3I

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var a int
	var b int
	var c int

	// Display the value of those variables.
	fmt.Println("Value of a is: %v", a)
	fmt.Println("Value of b is: %v", b)
	fmt.Println("Value of c is: %v", c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := 10
	cc := 10

	// Display the value of those variables.
	fmt.Println("Value of aa is: %v", aa)
	fmt.Println("Value of bb is: %v", bb)
	fmt.Println("Value of cc is: %v", cc)

	// Perform a type conversion.
	aaa := int32(10)

	// Display the value of that variable.
	fmt.Println("Value of aaa is: %v", aaa)
}
