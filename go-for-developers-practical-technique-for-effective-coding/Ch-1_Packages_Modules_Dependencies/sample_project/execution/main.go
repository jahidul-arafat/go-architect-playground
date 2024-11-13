package main

import (
	"fmt"
	"sampleproject/mathutils"
)

func main() {
	a, b := 10, 5
	fmt.Printf("Addition of %d and %d is %d\n", a, b, mathutils.Add(a, b))
	fmt.Printf("Subtraction of %d and %d is %d\n", a, b, mathutils.Subtract(a, b))
	demo()
}
