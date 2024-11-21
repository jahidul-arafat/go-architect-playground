package main

import "fmt"

// if your package needs an init function, have only one

// init function- A function that takes no argument and returns no argument
// function of this function is to initialize the package with any needed setup code
// this init function get run when the package is imported or executed

/*
Expected output:
init
main
*/
func main() {
	fmt.Println("main")
}

// since this is an init function, this doesnt require to be called from the main
func init() {
	fmt.Println("init 01")
}

// Declaring multiple init
// unlike other functions, init function can be decalred multiple time
// however, multiple init might make it difficult to understand which init will be called first
// which one would come first is decided during the runtime
// however, multiple init could easily lead to bug
// care must be taken if any init declarations have an order of precedence when running and exists in different files
func init() {
	fmt.Println("init 02")
}
