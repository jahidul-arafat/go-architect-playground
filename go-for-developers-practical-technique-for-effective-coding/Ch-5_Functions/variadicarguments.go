package main

import "fmt"

func main() {
	fmt.Println("Function Variadic Argument Testing...")
	fmt.Printf("Concated Names: %s\n\n", sayHelloToNames("A", "B", "C"))
	fmt.Printf("Concated Names: %s\n\n", sayHelloToNames("J"))
	fmt.Printf("Concated Names: %s\n\n", sayHelloToNames())

	nameList := []string{"AA", "BB", "CC"}         // slice of string
	sayHelloVariadicAtLast("AU-CSSE", nameList...) // must need to unpack the slice of strings to pass as a variadic arguments

	// Bad code to be replaced by variadic arguments
	id1 := 1
	id2 := 2
	id3 := 3
	ids := []int{id1, id2, id3}

	//lookUpUserInDB([]int{id1})
	//lookUpUserInDB([]int{id2})
	//lookUpUserInDB([]int{id2,id3})
	//lookUpUserInDB(ids)

	lookUpUserInDB(id1)
	lookUpUserInDB(id2)
	lookUpUserInDB(id1, id2)
	lookUpUserInDB(ids...) // unpacking the slice

}

/*
- Variadic argument can be used with other arguments, but must be the last in the parameter list
*/

// Example-1: Using variadic argument as a silo parameter
func sayHelloToNames(names ...string) (nameConcates string) { // names -> is a slice of string [aka Array]
	fmt.Printf("Type: %T, Length: %d\n", names, len(names))
	//for i, v := range names {
	//	fmt.Printf("%d->%s\n", i, v)
	//}
	nameConcates = ""
	for i, name := range names {
		if i == len(names)-1 {
			nameConcates += fmt.Sprintf("%s ", name)
			return
		}
		nameConcates += fmt.Sprintf("%s->", name)

	}
	return // returning <named-argument> nameConcates
}

// Example-2: Using variadic argument with other parameters in the function
// Note: Variadic argument must be the last in the parameter list

func sayHelloVariadicAtLast(group string, names ...string) {
	for _, name := range names { // names is a Slice (aka array) of string
		fmt.Printf("Hello %s to %s\n", name, group)
	}
}

// Bad Code to be replaced by Variadic arguments
//func lookUpUserInDB(ids []int){// agrument as slice of IDs
//	fmt.Println("Looking up IDs in the DB: ",ids)
//}

// Reconstucting using variadic argument
func lookUpUserInDB(ids ...int) { // will automatically convert the agrument as slice of IDs
	fmt.Println("Looking up IDs in the DB: ", ids)
}
