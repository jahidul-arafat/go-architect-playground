package main

import "fmt"

// Creating some custom simple Go Types from the existing
// All new types must be based on the existing types
type MyInt int
type MyString string
type MyMap map[string]string // map[key][]value

// create a custom type for a func() which returns a string
type Greeter func() string
type Adder func(int, int, int) (int, int)
type Subtractor func(int, int, int) (int, int)

// 1. Methods for simple custom types
// method for MyInt
func (i MyInt) IntGreeter() {
	fmt.Printf("Hello, I am %d\n", i)
}

// crete a sayHello function that takes the Greeter as an argument
func sayHelloToCustomTyPeGreeter(greetFunc Greeter) {
	fmt.Println(greetFunc()) // printing the returned string of the greetFunc
}

// create a function that takes the Adder as an argument
func addNumbers(addFunc Adder, num1, num2, counter int) {
	result, counter := addFunc(num1, num2, counter)
	fmt.Printf("[Add] Result: %d, Counter: %d\n", result, counter)
}

// create a function that takes the Subtractor as an argument
func subtractNumbers(subtractFunc Subtractor, num1, num2, counter int) {
	result, counter := subtractFunc(num1, num2, counter)
	fmt.Printf("[Sub] Result: %d, Counter: %d\n", result, counter)
}

// Create some complex types using struct
// where struct = zero/more fields + methods

// create a new struct named Student with fields name, age, and grades and methods like PrintStudentDetails
// this struct has a method called PrintStudentDetials
// Method name should not be in the struct definition
// struct definition should only have the fields like below
type Student struct {
	Name       string   // field 01 // non-reference field
	Age        int      // field 02	// non-reference field
	Grades     []string // field 03; slice of strings (index, value) // reference filed
	IsModified bool     // default bool value is "False" // non-reference field
}

// (s Student) indicates that this is a method, not a standalone function
// this (s Student) part is called "receiver"; this indicates that this methods is associated with the Student type
// s -> is a variable of type "Student" that you can use within the method body to access the struct's fields
func (s Student) PrintStudentDetails() {
	fmt.Printf("Name: %s, Age: %d, Grades: %v, IsModified: %v\n", s.Name, s.Age, s.Grades, s.IsModified)
}

// Methods-2: ImproveGrade check if any grade is below B, if so then changes this to
// Why using (s *Student) pointer? Because we want to modify the actual Student struct, not a copy of it
// if you are only changing the student's grade than you dont need to pass a *Student pointer, because s.Grades is a slice
// since s.Grades is a slice which is a reference type in Go. When you modify a slice inside a method, you are modifying the underlying array that the slice points to irrespective of whether you are passing a value or a pointer
// but if (s Student), a value is passed and if you want to change the student name, then the original student object will not be impacted
// only a local copy of that student object will be modified

// What about the potential bugs that can be introduced if pointers and values are not carefully handled in method receiver?

// func (s Student) ImproveGrades
// modification of the non-reference field requires a pointer to be passed as method receiver
// passing an entire value receiver will gonna have performance impact since entire struct has to be copied each time this method is called

func (s *Student) ImproveGrades() {
	s.Name = "Jahidul Arafat"
	s.IsModified = true
	for i, grade := range s.Grades { // s.Grades is a slice, so this will be upgraded even if (s Student) value/a copy is passed instead of a pointer
		if grade > "D" { // "A"<"B"<"C"<"D"<"E"<"F"
			s.Grades[i] = "D"
		}
	}
}

func main() {
	fmt.Println("Struct Basic Operations ...")
	//testingUserDefinedTypes()

	// create a new student object
	//var student Student
	//student.Name = "John Doe"
	//student.Age = 25
	//student.Grades = []string{"A", "F", "C"}

	myInt := MyInt(20)
	myInt.IntGreeter()

	// custom type func Greeter() testing

	sayHelloToCustomTyPeGreeter(func() string {
		return "Greeter Type Function: 10"
	})

	// create a variable of type Greeter
	//newGreeter := Greeter(func() string {
	//	return "Greeter Function: 20"
	//})
	//sayHelloToCustomTyPeGreeter(newGreeter)

	var newGreeter Greeter = func() string {
		return "This is a new Greeter"
	}

	sayHelloToCustomTyPeGreeter(newGreeter)

	// create a new adder function
	//var addFunc Adder = func(a, b, counter int) (int, int) {
	//	return a + b, counter + 1
	//}
	//addNumbers(addFunc, 10, 20, 0)

	// create a new adder function using anonymous function
	//addFunc := func(a, b, counter int) (int, int) {
	//    return a + b, counter + 1
	//}
	//addNumbers(addFunc, 10, 20, 0)

	// create a new adder function
	addFunc01 := Adder(func(a, b, counter int) (int, int) {
		return a + b, counter + 1
	})

	subtarctFunc01 := Subtractor(func(a, b, counter int) (int, int) {
		return a - b, counter - 1
	})

	addNumbers(addFunc01, 10, 20, 0)
	subtractNumbers(subtarctFunc01, 10, 20, 0)

	// create a new student object
	//var student Student
	//student.Name = "John Doe"
	//student.Age = 25
	//student.Grades = []string{"A", "F", "C"}

	// Now let's create a new student object using the User Defined type

	student := Student{
		Name:       "John Doe",
		Age:        25,
		Grades:     []string{"A", "F", "C"},
		IsModified: false, // default value is "False"
	}

	fmt.Printf("Student Details [before upgrading]: %+v\n", student)
	student.PrintStudentDetails()

	// lets modify the students grade and print the student details again
	student.ImproveGrades()
	fmt.Printf("Student Details [after upgrading]: %+v\n", student)
	student.PrintStudentDetails()
}

func testingUserDefinedTypes() {
	// Now use these new types like the other types we have encountered
	var i1 MyInt = 1
	fmt.Printf("%[1]T->%[1]v", i1)

	// define a new map using the regular type "map"
	var trees = make(map[string]string)
	//trees = make(map[string]string)
	trees["A"] = "One"
	fmt.Printf("%[1]T->%[1]v\n", trees)

	//Now create a new map named "newTrees" using the user defined type "MyMap"
	var newTrees MyMap = make(MyMap)
	newTrees["A"] = "One"
	fmt.Printf("%[1]T->%[1]v\n", newTrees)

	// implicit way
	anotherTree := MyMap{
		"A": "one",
		"B": "two",
	}
	fmt.Printf("%[1]T->%[1]v\n", anotherTree)

	// lets cat a map to MyMap type
	var castingTree = map[string]string{"A": "One", "C": "Three"}
	castingTree["B"] = "Two"
	fmt.Printf("%[1]T->%[1]v\n", castingTree)

	// Create a new string using the user defined type "MyString"
	var newStr MyString = "New String defined with New String Type"
	fmt.Printf("%[1]T->%[1]v\n", newStr)

	// Create a new string usign the default string type
	var myStr string = "original string"
	fmt.Printf("%[1]T->%[1]v\n", myStr)

	// implicit way
	anotherStr := MyString("Another String")
	fmt.Printf("%[1]T->%[1]v\n", anotherStr)

}
