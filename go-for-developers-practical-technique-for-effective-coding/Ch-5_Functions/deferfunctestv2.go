package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Second Line") // Defer is LIFO
	fmt.Println("First Line")
	//defer fmt.Println("Third Line") // Last-In-First-Out in the defer stack
	//fmt.Println(letReturnError("Test", "Ally"))

	fileOpsStatus, err := CreateFile("defer.txt")
	fmt.Println(fileOpsStatus, err)

}

//func letReturnError(greet, name string) (string, string, error) {
//	if len(greet) == 0 || len(name) == 0 {
//		return "", "", fmt.Errorf("length Greet %v, length Name: %v", greet, name)
//	}
//	return strings.ToUpper(greet), strings.ToUpper(name), nil
//
//}

// Example: Create a file and tries to write data to it
// How to use defer statement here to ensure the file closes properly?
/*
1.We add defer fileHandler.Close() immediately after successfully creating the file.
2.The defer statement ensures that fileHandler.Close() will be called when the function returns, whether it's a normal return or due to an error.
3.This guarantees that the file will be properly closed, releasing system resources and ensuring all data is flushed to disk.
*/
func CreateFile(fName string) (string, error) {

	fileHandler, err := os.OpenFile(fName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	//fileHandler, err := os.Create(fName)
	if err != nil {
		return "", fmt.Errorf("failed to create the file %s", fName)
	}

	defer fileHandler.Close() // this gonna call last // this ensures that the file is closed when returns either for Normal return or error return

	_, err = fileHandler.WriteString("Hello to File \n")
	if err != nil {
		return "", fmt.Errorf("failed to write to file %s", fName)
	}
	return fmt.Sprintf("File Creation Successful: %v", fName), nil

}
