package implementation

import (
	"fmt"
	"os"
)

type FileLogger struct {
	FileName string // field
}

func (f FileLogger) Log(message string) {
	// create a file and write the message to it
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close() // will be called when the function returns
	file.WriteString("File: " + message + "\n")
}

func (f FileLogger) PrintMe() {
	fmt.Println("Print Me>>[F]>>")

}
