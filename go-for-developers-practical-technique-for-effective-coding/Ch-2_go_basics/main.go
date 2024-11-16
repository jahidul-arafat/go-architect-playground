package main

import (
	"fmt"            // for printing and formatting values-> fmt package is used
	"gobasics/fruit" // <modulename>/<packagesname>
	gpath "path"     // aliasing the standard PATH package as gpath
)

type User struct {
	Name string
	Age  int
}

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

	fmt.Printf("Value: %s, Type: %T\n", name, name) // Should show: Type of score: float64

	anotherWorld := "Hello , another world িোে" // simple loop would fail to print the UFT characters --> Sol: use range
	for i := 0; i < len(anotherWorld); i++ {
		fmt.Printf("Char[%d]=%s\n", i, string(anotherWorld[i]))
	}
	fmt.Println("Using the Range for Printing")

	for i, c := range anotherWorld {
		fmt.Printf("Char[%d]=%s\n", i, string(c))
	}

	// 2.5 Variable Declaration
	fmt.Println("Variable Declaration\n--------------------")

	// declaring the variable
	// syntax: var <variable_name> <type> = <initialization value>
	//var i int = 42
	//var f float64 = 3.14
	//var bo bool = true
	//var s string = "test string"

	// lets go infer the type of the variable from the value being assigned
	i1 := 42
	j1 := uint8(45) // unsigned interger 8 bit = max value is 255
	f1 := 45.67
	f2 := float32(45) // 45.00
	bo1 := false
	s1 := "another test string"

	// printing the variables
	fmt.Printf("Value=%d, Type=%T\n", i1, i1)
	fmt.Printf("Value=%.2f, Type=%T\n", f1, f1)
	fmt.Printf("Value=%d, Type=%T\n", j1, j1)
	fmt.Printf("Value=%.2f, Type=%T\n", f2, f2)
	fmt.Println(bo1)
	fmt.Println(s1)

	// Create 4 new variables in the same line of different types
	//i2, f2, b2, s2 := 42, 53.56, false, "all-in-one-line"
	// make sure there is no unused variables--> else go compiler will return errors
	i2, f2, b2, _ := Values() // _ is the blank identifier

	fmt.Printf("var i %T=%v\n", i2, i2)
	fmt.Printf("var f %T=%f\n", f2, f2)
	fmt.Printf("var b %T=%v\n", b2, b2)
	//fmt.Printf("var s %T=%v\n", s2, s2)

	// Variable ZERO value testing
	fmt.Println("Variable Zero Value Testing")
	var (
		s     string
		r     rune
		bt    byte
		i     int
		ui    uint
		f     float32
		cx    complex64
		bx    bool
		myarr [2]int // an array of 02 integers
		myobj struct {
			myb  bool
			myar [3]int
		}
		si  []int          // array of integers; size is not defined
		ch  chan string    // go channel -> for goroutine
		mp  map[int]string // key->int, value-> string
		fn  func()
		ptr *string // string pointer // Initial value=NIL
	)

	// print all the variables
	fmt.Printf("var s %T=%v\n", s, s)
	fmt.Printf("var r %T=%v\n", r, r)
	fmt.Printf("var bt %T=%v\n", bt, bt)
	fmt.Printf("var i %T=%v\n", i, i)
	fmt.Printf("var ui %T=%v\n", ui, ui)
	fmt.Printf("var f %T=%v\n", f, f)
	fmt.Printf("var c %T=%v\n", cx, cx)
	fmt.Printf("var b %T=%v\n", bx, bx)
	fmt.Printf("var myarr %T=%v\n", myarr, myarr)
	fmt.Printf("var myobj %T=%v\n", myobj, myobj)
	fmt.Printf("var si %T=%v\n", si, si)
	fmt.Printf("var ch %T=%v\n", ch, ch)
	fmt.Printf("var mp %T=%v\n", mp, mp)
	fmt.Printf("var fn %T=%v\n", fn, fn)
	fmt.Printf("var ptr %T=%v\n", ptr, ptr)

	// 2.6 Constants
	const stdname = "JAHID"
	//stdname = "Another Name"
	fmt.Println(stdname)

	// Failing constants
	//const gopher = Names() // func Names()
	//const nameslist = []string{"A", "B", "C", "D"}
	//fmt.Println(gopher)

	// 2.7 Naming Identifiers
	names := []string{"Ally", "Billy", "Cilly", "Dilly"}
	for i, n := range names {
		fmt.Printf("%v-->%q\n", i, n)
	}

	// dont use the Go standard package name as identifier of your own
	path := gpath.Join("dir", "file.xtx")
	fmt.Println("Path: ", path)

	//anotherPATH := path.Join("dir", "test.txt") // you named your identifier as the go standard package name "path"
	// that's why original standard Path package becomes inaccessible
	// Calling the Print() function from the package <fruit>
	fruit.Print()
	fmt.Println(fruit.Apple)
	fmt.Println(fruit.Banana)
	//fmt.Println(fruit.cherry)     // not exported from the hosting package <fruit>
	//fmt.Println(fruit.strawberry) // not exported from the hosting package <fruit>

	// Printing struct data
	u := User{
		Name: "Jahid",
		Age:  30,
	}
	fmt.Printf("Struct Info: %[1]T and Struct Data: %#[1]v\n", u)

	// Reversing the order of the arguments
	fmt.Printf("Explicit: %[4]s, %[3]s, %[2]s,%[1]s\n", "one", "two", "three", "four")

}

func Values() (int, float32, bool, string) {
	return 42, 3.56, false, "all-in-one" // bad practice; returning so many values from the function
}

func Names() string {
	return "Gopher"

}
