package main

import "fmt"

func main() {
	/*
		Means
		1. Function can be argument or return statement of another function, just like other variables
		2. Create a function and assign it to a variable for later execution
	*/
	fmt.Println("Function First Class Citizen ...")
	fnWithIntReturn := func() (meaning int) {
		fmt.Println("Hello var-assigned function")
		return // returning integer value of <meaning>
	}
	//fmt.Printf("Meaning: %v", fn())
	sayHelloToFunction(fnWithIntReturn)

	fnWithoutReturn := func() {
		fmt.Println("This function passed as argument inside another function [No return]")
	}

	sayHelloToFunctionWithoutReturn(fnWithoutReturn)

	// Example- Creating Annonymous Function during the to avoid unneeded variable declaration
	name := "Annynomous Name"
	sayHelloToAnnynomousFunction(func() { // this function as a argument is annynomous because we didnt define it earlier
		fmt.Printf("Hello to Annynomous Function with [%v]", name)
	})
}

func sayHelloToFunction(fn func() int) {
	fmt.Printf("Func-as-Argument-Meaning: %v\n", fn())
}

func sayHelloToFunctionWithoutReturn(fn func()) {
	fmt.Println("Say hello to function without return..")
	fn()
}

func sayHelloToAnnynomousFunction(fn func()) {
	fn()
}
