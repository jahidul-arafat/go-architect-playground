package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Defer and Panic Testing") // Defer is a LIFO

	//defer fmt.Println("one")
	////defer fmt.Println("two")
	//defer panic("two") // this will be re-executed last, after finishing all the non-panic defers
	//defer fmt.Println("three")

	// wht if the os.Exit(1); the program is explicitly terminated
	fmt.Println("One")
	defer fmt.Println("Two") // will never be executed since os.Exit(1) will terminate the program even before reachign to the defer call
	os.Exit(1)               // see, program is explicitly shutdown
	fmt.Println("Three")     // will never be executed

}
