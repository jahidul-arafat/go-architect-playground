package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// go statically typed compiled programming language
	var pi float64 = 3.14
	fmt.Println(pi)

	var week int = 7
	fmt.Println(week)

	//var results float64 = nil / 3.14
	//fmt.Println(results)

	// 2.3 Numbers
	var maxUint8 uint8 = 255           //Max Uint8 size = 255; Beyond that a value overflow will trigger
	fmt.Println("value: ", maxUint8+5) // will wrap tp 255+5=260, it wraps around 0, then 1,2,3,4

	// 2.4 String and UTF-8
	a := "Hello\t\tWorld"
	fmt.Println(a)

	b := `Say "Hello" to Golang`
	fmt.Println(b)

	c := `Say \t\t\n "hello" to golang` // inside backslash, \t\n would not work
	fmt.Println(c)

	d := `# json data for testing
{
	"id":1,
	"name": "Jahidul",
	"email": "test@test"
}
	`
	fmt.Println(d)

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", p) // Output: Person: {Name:Alice Age:30}

	name := "Alice"
	//age := 30
	//score := 95.5

	fmt.Printf("Value: %s, Type: %T", name, name) // Should show: Type of score: float64

}
