package main

import "fmt"

// Here we will demonstrate how defer function execution works
// The deferred functions are executed in Last-In-First-Out (LIFO) order when the surrounding function (deferExample) returns.
func main() {
	fmt.Println("Starting Defer Function Example...")

	// defer a function call
	//defer func() {
	//	fmt.Println("Line 3")
	//}()
	//
	//defer fmt.Println("Line 2")
	//
	//fmt.Println("Line 1")

	fmt.Printf("main: %v", MeaningOfLife())
}

// MeaningOfLife - defer function with named return and ambiguity
/*
The change from 42 to 1 happens at runtime, not compile time. Here's a detailed explanation of the process:

1.
At compile time, the compiler sees the return 42 statement and the deferred function, but it doesn't execute any of this code.
2.
At runtime, when the function is called:
a. The return 42 statement is encountered first.
b. This sets the named return value meaning to 42.
c. However, the function doesn't immediately return to its caller.
3.
Before the function actually returns, Go executes any deferred functions in last-in-first-out order.
*/

func MeaningOfLife() (meaning int) { // default int value is 0 if not else defined
	fmt.Println("Start: ", meaning)

	defer func() { // this will execute after we return 42 from the function
		fmt.Println("Deferred: ", meaning)
		meaning = 1 // this will be the final return
	}()

	return 42
}
