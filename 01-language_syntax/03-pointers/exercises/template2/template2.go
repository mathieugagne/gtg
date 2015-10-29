// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/b6-FNFOToO

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name string
}

// Create a function that changes the value of one of the user fields.
func setName(u *user, newName string) {
	u.name = newName
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	myUser := user{
		name: "Mathieu",
	}

	// Display the value of the variable.
	fmt.Println("My Name:\t", myUser.name)

	// Share the variable with the function you declared above.
	setName(&myUser, "Khaled")

	// Display the value of the variable.
	fmt.Println("My new name:\t", myUser.name)
}
