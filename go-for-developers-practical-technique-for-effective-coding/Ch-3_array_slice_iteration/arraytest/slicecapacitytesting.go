package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slice Capacity Testing")

	names := []string{}
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))

	names = append(names, "Ally")
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))

	names = append(names, "Billy")
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))

	names = append(names, "Cilly")
	fmt.Println("len: ", len(names))
	fmt.Println("cap: ", cap(names))

	// Different ways to create a slice
	a := []string{}
	var b []string
	c := make([]string, 1, 3)           // length=1, capacity=3
	c = append(c, "foo", "bar", "joll") //3+1(empty string:default)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%#v\n", c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Printf("%#v\n", letters)
	fmt.Printf("Capacity: %v\n", cap(letters))
	fmt.Printf("Length: %v\n", len(letters))

	// slice of slices
	fmt.Printf("Slice [2:5]: %v\n", letters[2:5]) //c,d,e
	fmt.Println(letters[:len(letters)])

	subLetters := letters[0:4]
	fmt.Printf("Sub Letters: %v\n", subLetters)
	subLetters[0] = "i"
	fmt.Printf("Original: %v\n, "+
		"After Slice Modification: %v\n", letters, subLetters)

	for i, val := range subLetters {
		subLetters[i] = strings.ToUpper(val)
	}

	fmt.Printf("Original: %v\n ToUpper: %v\n", letters, subLetters)

	subLetterWithCopy := make([]string, 3) // Length=3
	copy(subLetterWithCopy, letters[:3])
	fmt.Printf("%v\n", subLetterWithCopy)
	subLetterWithCopy[0] = "K"
	fmt.Printf("Original: %v\n SubLettersCopy: %v\n", letters, subLetterWithCopy)

	nameList := [4]string{"Ally", "Billy", "Cilly", "Dilly"} // its an array, not a slice
	//slicesOnly([]string(nameList))                           //failed the compiler-> invalid syntax                    // casting an array into a slice
	slicesOnly(nameList[:]) // converting the array into slice

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}

func slicesOnly(names []string) {
	for i := range names {
		fmt.Printf("%d-->%s\n", i, names[i])
	}

}
