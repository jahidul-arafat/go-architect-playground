package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("If Early Return - Avoid Complicated If-ElseIf-Else statements")

	// In Go, its good practice to avoid else, instead use "early return" strategy
	var greet bool // default is "False"
	if greet {
		fmt.Println("Hello")
		return // if return is used, make sure there are no subsequent major code portion left after than
	}

	fmt.Println("GoodBye")

	/*
		How to clean up the below complex if-elseif-else statements
	*/
	var name string
	fmt.Println("Enter a name: ")
	fmt.Scanln(&name) // store the entered name in the address of variable "name";
	// that's why &->address of name is used
	name = strings.TrimSpace(name)

	//name := "Ally"
	if strings.ToLower(name) == "ally" {
		fmt.Printf("Hello %s\n", "ally")
	} else if strings.ToLower(name) == "billy" {
		fmt.Printf("Hello %s\n", "billy")
	} else {
		fmt.Printf("[Default] Hello %s\n", name)
	}

	originalString := "Original string"
	dontModifyOriginalString(originalString)
	fmt.Printf("Original string after x Call: %s\n",
		originalString)

	modifyOriginalString(&originalString)
	fmt.Printf("Original String After mod Call: %s\n", originalString)

}

func dontModifyOriginalString(s string) {
	s = "dont modify string"
	fmt.Printf("Inside xModifiyOriginalString: %s\n", s)
}

// modifyOriginalString(&originalString) // passign the address
func modifyOriginalString(s *string) {
	// * -> is a deferencing operator. It's used to access the value stored at a memory address
	*s = "modified original string"
}
