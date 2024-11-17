package main

import "fmt"

func main() {
	var month int
	fmt.Println("Enter month [integer 1-12]: ")
	fmt.Scanln(&month)

	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	default:
		fmt.Printf("Month %d not found\n", month)
	}
}
