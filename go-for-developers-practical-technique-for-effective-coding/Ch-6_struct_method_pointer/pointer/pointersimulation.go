package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

/*
In Go, pass by reference is not supported in the traditional sense like in some other languages (e.g., C++). Instead, Go uses pass by value consistently, even when dealing with pointers. This means:
	•	When you pass a variable to a function, Go passes a copy of the value.
	•	When you pass a pointer, Go passes a copy of the pointer.
This is why, in your example, the pointer itself has a unique memory address when passed to a function, even though it still points to the same data.
*/

// create two function to check if the parameter <Student> passed to it holds the same memory address or not
/*
	1.	A copy of the pointer (s) is created and passed to the function.
	•	This copy has its own memory address (e.g., 0xc000006028).
	2.	However, the copy still stores the same address (0xc000016120) as the original pointer (&std01), so it points to the same data.
*/

// A value receiver cant mutate the data, when the pointer receiver can
func CheckAndChangeMe(s *Student) {
	fmt.Printf("3. [Check01] Address of the Pointer itself: %+v\n", &s)
	fmt.Printf("4. [Check02] Address stored in the pointer: %p\n", s)
	s.Name = "NewChangedName" // check if this has a global impact
}

func main() {
	fmt.Println("This is a Pointer Simulation ....")

	std01 := Student{
		Name: "Billy",
		Age:  20,
	}

	fmt.Printf("Student Details: %+v\n", std01)
	// print the memory address of std01
	fmt.Printf("1. Memory Address of std01: %p\n", &std01)

	// create a new pointer to std01 and if that pointer has the same memory address as that of <&std01>
	std01Ptr := &std01                                              // create a pointer with the address of the stu01
	fmt.Printf("2. Memory Address of std01Ptr: %p\n", std01Ptr)     // same memory address as of <&std01>
	fmt.Printf("Value of std01 through std01Ptr: %+v\n", *std01Ptr) //*-> means deferencing to get the actual values store in that specific address

	// change of the value of std01 name though the std01Ptr and check if it has a global impact, changing the Name in all of its references
	std01Ptr.Name = "changedPtrBilly"
	fmt.Printf("After Changes[std01Ptr]: %+v\n", *std01Ptr) // derefence first to get the value // dereference should always be done through (*)
	fmt.Printf("After Changes[std01]: %+v\n", std01)

	// check the address
	// even though the two points has different obejct id (aka different memory address), they both still points to the same underlying std01 data
	CheckAndChangeMe(&std01) // 0x140001b8078 // create a copy of the pointer, but still points to the same std01 values
	//CheckAndChangeMe(&std01) // 0x140001b8090 // create a copy of the pointer, but still points to the same std01 values

	fmt.Printf("Student Details: %+v\n", std01)
	fmt.Printf("Value of std01 through std01Ptr: %+v\n", *std01Ptr)
}
