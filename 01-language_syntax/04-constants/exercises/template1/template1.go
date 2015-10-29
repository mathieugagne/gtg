// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/qG3-9yn5_6

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import "fmt"

// Declare a constant named server of kind string and assign a value.
const server string = "mtlpc-mgagne"

// Declare a constant named port of type integer and assign a value.
const port int = 8080

// main is the entry point for the application.
func main() {
	// Display the value of both server and port.
	fmt.Println(server, ":", port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	const answer = 10 / 2.0

	// Display the value of the variable.
	fmt.Println(answer)
}
