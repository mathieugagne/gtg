// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/ItPe2EEy9X

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int16
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct
	myUser := user{
		name:  "Mathieu",
		email: "mathieu@gagne.com",
		age:   31,
	}

	// Display the field values.
	fmt.Printf("%+v\n", myUser)
	fmt.Println("Name: ", myUser.name)
	fmt.Println("Email: ", myUser.email)
	fmt.Println("Age: ", myUser.age)

	// Declare a variable using an anonymous struct.
	otherUser := struct {
		name  string
		email string
		age   int16
	}{
		name:  "Khaled",
		email: "kelbadawi@vg.com",
		age:   25,
	}

	// Display the field values.
	fmt.Printf("Name:\t%v\n", otherUser.name)
	fmt.Printf("Email:\t%v\n", otherUser.email)
	fmt.Printf("Age:\t%v\n", otherUser.age)
}
