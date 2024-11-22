package main

import "fmt"

/*
	Heap- a larger, slower memory region used to store variables that have a dynamic or global lifetime.
	Stack - a small, fast memory region used for temporary storage of local variables within a function.

	1. If function returns a pointer to <variable>; this forcing Go to allocate <variable> on the heap
	to ensure it remains valid even after the function exists
	2. If function returns a <variable> by value; this is stored on the stack and destroy when the function exits.

	Note: Go only support call by value, go doesn't support call by references

	Which is more difficult to Clean up?
	========================================
	- Heap memory is more difficult to clean up because the Go runtime's GC has to track references to <variable>
	and determine when they are no longer in use.
	- Improper memory management (e.g. unintentional retention of references) can lead to memory leaks.

	- Stack memory (LIFO) is easier to clean up because it is automatically reclaimed when the function scope ends.
	- this doesn't require GC.
*/

// create a heap, that returns a pointer to a <variable> which eventually stores at heap memory, havign global Lifecycle
func createHeapVariable() *int {
	x := 42   // local variable
	return &x // escape to heap because the address is returned

}

// create a stack
// allocate memory on the stack
func createStackVariable() int {
	x := 42
	return x
}
func main() {
	fmt.Println("Heap Vs Stack Operations in Go....")

	heapVar := createHeapVariable()
	fmt.Printf("Heap Variable [address]: %+v\n", heapVar)                   // this will print the address since &x was returned
	fmt.Printf("Heap Varibale [value(after unpacking)]: %+v\n\n", *heapVar) // unpacking the Heap variable

	stackvar := createStackVariable()
	fmt.Printf("Stack Variable: %+v\n", *&stackvar) // &x-> referencing *&-> deferencing(refercenced x)-> cancels it out

}
