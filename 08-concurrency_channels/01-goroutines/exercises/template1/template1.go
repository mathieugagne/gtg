// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/KlKIYq9s_3

// Create a program that declares two anonymous functions. Once that counts up to
// 100 from 0 and one that counts down to 0 from 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// init is called prior to main.
func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

// main is the entry point for all Go programs.
func main() {
	// Declare a wait group and set the count to two.
	var wg sync.WaitGroup
	wg.Add(1)

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		for {
			fmt.Println("I'm alive!")
		}
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts up from 0 to 100 and
		// display each value.
		for count := 0; count < 100; count++ {
			fmt.Printf("%v", count)
		}
		time.Sleep(10 * time.Second)
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	wg.Wait()
}
