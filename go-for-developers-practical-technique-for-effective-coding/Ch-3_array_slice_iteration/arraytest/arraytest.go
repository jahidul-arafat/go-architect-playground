package main

import "fmt"

func main() {
	fmt.Println("Array Testing")

	// Create an array of 4 strings
	// (if you know the memory requried at creation time)
	nameArr := [4]string{"Ally", "Billy", "Cilly", "Dilly"}
	fmt.Printf("Type: %[1]T, Elements: %[1]v\n", nameArr)
	for i, n := range nameArr {
		fmt.Printf("Index: %d, Name: %v\n", i, n)
	}

	// Create a slice of strings
	// if you dont know the memory requried at creation time
	// Most often used in daily situation
	nameSlice := []string{"Ally", "Billy", "Cilly", "Dilly"}
	fmt.Printf("TYpe: %[1]T, Elements: %[1]v\n", nameSlice)

	// Arrays and slices dont need to be initialized before that are used
	// is the type initialzied
	a := [5]int{10, 20, 30, 40, 50}
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Printf("%#v\n", a)

	//f := []string{"Ally", "Billy", "Cily", "Dilly"}
	//fmt.Printf("%#v\n", f)
	//fmt.Println(f[6])

	var c [3]bool
	fmt.Printf("%#v\n", c)

	// Is the array type definition matching during array assignment
	a1 := [2]string{"one", "two"}
	var a2 [2]string
	a2 = a1 // a2 will have its own copy from a1 and after than changing a1 values would not impact the a2 valuies
	fmt.Println(a2)
	a1[0] = "anotherOne"
	fmt.Println(a2)
	fmt.Println(a1)

	//a3 := [3]string{}
	//a3 = a2
	//fmt.Println(a3)

	// Is the Slice type definition marching
	s1 := []string{"one", "two"}
	var s2 []string
	s2 = s1
	fmt.Println(s2)

	//s3 := []int{}
	//s3 = s2
	//fmt.Println(s3)

	// Appending an slice; append doesnt work with array
	s1 = append(s1, "Four", "Five", "Six")
	fmt.Println(s1)

	// Appendign a slice with the contents of another slice
	sourceNameList := []string{} // type string
	sourceNameList = append(sourceNameList, "X")
	fmt.Printf("Type: %[1]T, Value: %[1]v\n", sourceNameList)
	sourceNameList = append(sourceNameList, "B", "C")
	fmt.Println(sourceNameList)

	moreNameList := []string{"MoreName1", "MoreName2"} // type slice of string
	fmt.Printf("Type: %[1]T, Values: %[1]v\n", moreNameList)
	//sourceNameList = append(sourceNameList, moreNameList) // you cant directly append a slice with another slice
	// you would need to iterate each elements of the moreNameList slice to append to the sourceNameList
	// Non-efficient way of appending
	//for _, name := range moreNameList {
	//	sourceNameList = append(sourceNameList, name)
	//}
	sourceNameList = append(sourceNameList, moreNameList...) // variadic argument
	fmt.Printf("Type: %[1]T, Size: %[2]v, Elements: %[1]v\n", sourceNameList, len(sourceNameList))

}
