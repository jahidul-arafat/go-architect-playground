package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Functions!!")
	fmt.Println(sayHello("Hello", "Ally"))
	uGreet, uName := returnTwo("hello", "cilly")
	//fmt.Sprintf("%s,%s", uGreet, uName) // this not gonna print anything, this is only for string formatting
	fmt.Printf("%s,%s!!\n", uGreet, uName)

	// print the result of returnThree function
	greet, name, err := returnThree("Hello", "Dilly")
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%s,%s!!", greet, name)
	}

	// Call the function with variadic arguments
	getVariadicResult := sayHelloVariadic("Ally", "Billy", "Cilly", "Kelly", "Ally")
	fmt.Printf("%s\n", getVariadicResult)
	greetListNames := strings.Fields(strings.ToLower(getVariadicResult)) // Slice of strings
	fmt.Printf("%v\n", greetListNames)

	// count how many times each name is greeted
	counter := map[string]int{}
	for _, name := range greetListNames {
		counter[name]++
	}
	fmt.Printf("%#v\n", counter)

	//greetVariadic := sayHelloVariadic("Hello", "Ally", "Billy", "Cilly")
	//fmt.Println(greetVariadic)

	// info function with multiple return
	names := []string{"Ally", "Billy", "Cilly", "Dilly"}
	//var names []string
	//names = make([]string, 0, len(names))
	gs, l, _ := info(names)
	fmt.Printf("gs=%v, l=%v\n", gs, l)

	// Example with named return
	// Example-1
	status, name := isValid()
	fmt.Printf("Named Return=%v, Name=%v\n", status, name)

	// Example-2
	lat, lng := coordinates()
	fmt.Printf("Latitude: %v, Longitude: %v", lat, lng)

}

// create a variadic function
func sayHelloVariadic(greetings ...string) string { // greeting... -> is a Slice of strings
	result := ""
	fmt.Printf("%T\n", greetings)
	for _, greet := range greetings {
		result += fmt.Sprintf("%s ", greet)
	}
	return strings.TrimSpace(result)
}

//func sayHelloVariadic(greetings ...string) string {
//	result := ""
//	for _, greeting := range greetings {
//		result += fmt.Sprintf("%s, ", greeting)
//	}
//	return strings.TrimSpace(result)
//}

func sayHello(greeting, name string) string {
	return fmt.Sprintf("%s, %s!!", greeting, name) // don't immediately print anything,
	// it's a string formatting which can be returned
}

func returnTwo(greeting, name string) (string, string) {
	return strings.ToUpper(greeting), strings.ToUpper(name)

}

// create a function that returns three items with an error return at the end
// if error to return-> this to be at the last
func returnThree(greeting, name string) (string, string, error) {
	if len(greeting) == 0 || len(name) == 0 {
		return "", "", fmt.Errorf("Name or Greeting must not be nil") // early return
	}

	return strings.ToUpper(greeting), strings.ToUpper(name), nil
}

func info(s []string) (string, int, int) {
	gs := fmt.Sprintf("%#v", s)
	l := len(s)
	c := cap(s)
	return gs, l, c
}

// create a function with named return, which automatically return the argument
func isValid() (valid bool, name string) { // this will return the "valid" variable; since the return is tagged with variable "valid", this is called named return
	// valid is initialized with <false>
	name = "Olly"
	fmt.Printf("Before=%v\n", valid)
	valid = true
	fmt.Printf("After=%v\n", valid)

	return // returning variable "valid" with its latest <value>
}

func coordinates() (lat float64, lng float64) {
	lat = 42.1344354
	lng = -71.2344535
	return
}
