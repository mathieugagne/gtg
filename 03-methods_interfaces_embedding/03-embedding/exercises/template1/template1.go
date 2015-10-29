// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/c9Qrsq8QFe

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the Search method for hockey, call into the Search method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing searches.
type matcher interface {
	Search(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// Search checks the value for the specified term.
func (s sport) Search(searchTerm string) bool {
	if strings.Contains(s.team, searchTerm) ||
		strings.Contains(s.city, searchTerm) {
		return true
	}

	return false
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.
type hockey struct {
	sport
	arena string
}

// Implement the matcher interface for hockey.
func (h hockey) Search(searchTerm string) bool {
	// Make sure you call into Search method for the embedded
	// sport type.
	return strings.Contains(h.arena, searchTerm) || h.sport.Search(searchTerm)
}

// main is the entry point for the application.
func main() {
	// Define the term to search.
	searchTerm := "Montreal"

	// Create a slice of matcher values to search.
	sportArenas := []matcher{
		sport{
			team: "Impact",
			city: "Montreal",
		},
		hockey{
			sport: sport{
				team: "Canadiens",
				city: "Montreal",
			},
			arena: "Bell Center",
		},
	}

	// Display what we are searching for.
	fmt.Printf("Looking for the term %s\n", searchTerm)

	// Range of each matcher value and check the search term.
	for _, match := range sportArenas {
		if match.Search(searchTerm) {
			fmt.Printf("%s matches %s\n", match, searchTerm)
		}
	}
}
