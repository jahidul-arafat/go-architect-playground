package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Student struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Hello, Maps!")

	// Define a map
	// Create a map of users
	users := map[string]string{} // map[<keytype>]<valuetype>
	fmt.Printf("%#v\n", users)
	users["A"] = "Ally"
	users["B"] = "Billy"
	users["C"] = "Cilly"
	fmt.Printf("%[1]T--> %#[1]v\n", users)

	for key, val := range users {
		fmt.Println(key, val)
	}

	// Create another map of Cars
	cars := map[string]string{
		"T": "Toyota",
		"D": "Dodge",
	}
	cars["T"] = "Tata"

	fmt.Println(len(cars), cars["K"])

	// Define the map: create a nil map
	var trees map[string]int
	fmt.Printf("%[1]T --> %[1]v\n", trees)

	// initialize the map to add value to it
	trees = make(map[string]int)
	fmt.Printf("%[1]T --> %[1]v\n", trees)
	trees["Oak"] = 100
	trees["Pine"] = 300
	fmt.Printf("%[1]T --> %[1]v\n", trees)

	// check if a key exists in the map
	keyName := "Oak"
	//value, isExists := trees[keyName]
	//fmt.Println(value, isExists)
	if value, isExists := trees["Oak"]; isExists {
		fmt.Printf("Key: %v with value: %v exists: %v\n", keyName, value, isExists)
	} else {
		fmt.Printf("Key: %v doesnt found\n", keyName)
	}

	// Check if a value exists in the map
	valueToFind := 100
	for key, val := range trees { // iteration order is random
		if val == valueToFind {
			fmt.Printf("Value %v found in Key: %v\n", val, key)
		}
	}

	// delete a key from the map
	// Only one key at a time can be deleted from map
	keyToBeDeleted := "Pine"
	delete(trees, keyToBeDeleted)
	fmt.Println(trees, len(trees))

	// Using struct as keytype in map
	// Create a user map of keytype struct <Student>
	students := map[Student]string{}
	fmt.Printf("%#v\n", students)

	student1 := Student{ID: 1, Name: "Ally"}
	fmt.Println(student1)
	students[student1] = "ally@gmail.com"
	fmt.Printf("%v\n\n", students)

	// The bug in the map, key doesnt exists still returning ""
	datas := map[int]string{}
	datas[1] = "Hello World!!!"

	keyToCheck := 1
	_, isExists := datas[keyToCheck]
	if !isExists {
		fmt.Printf("Key %v doesnt exists\n", keyToCheck)
		os.Exit(1) // exist the program if key not found
	}

	// Map with sentence example
	// Define and initialize a map to keep the count of each words
	counts := map[string]int{}
	fmt.Printf("%#v\n", counts)

	sentence := "Auburn University, Samuel Ginn College of Engineering - Auburn"
	// Split the string into words, convert each word to lowercase and find the occurance of each work
	words := strings.Fields(strings.ToLower(sentence)) // a Slice of strings (aka Array of Strings)
	fmt.Printf("%[1]T -> %#[1]v\n", words)
	for _, word := range words { // iterate over a Slice or Array
		counts[word]++
	}
	fmt.Printf("%#v\n\n", counts)

	// Create a map where we will store the struct <Strudent> as value
	stdDataMap := map[int]Student{}
	stdDataMap[1] = student1
	fmt.Println(stdDataMap)

	// lets try to change the Name of the student for 1
	//stdDataMap[1].Name = "Billy" // Not the right approach to update the value of a complex struct in the map
	// Solution
	// 1. Change the Name of student in the student1 struct
	student1.Name = "Billy"
	stdDataMap[1] = student1 // reinsert it into the map
	fmt.Println(stdDataMap)
	fmt.Println(student1)

	// GO doesnt provide the default ways to get the list of keys and values from a map
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	// extract all the keys from the maps
	keys := make([]int, 0, len(months)) // create a slice of ints with initial length 0 and max capacity 12
	for k := range months {             // map <month> is unordered
		keys = append(keys, k)
	}

	sort.Ints(keys)   // sort the retrieved Key Slice
	fmt.Println(keys) // print the keys

	// extract all the values from the maps, sort them and print them
	//var monthNames []string
	monthNames := make([]string, 0, len(months))
	for _, month := range months {
		monthNames = append(monthNames, month)
	}
	sort.Strings(monthNames)
	fmt.Println(monthNames)

}
