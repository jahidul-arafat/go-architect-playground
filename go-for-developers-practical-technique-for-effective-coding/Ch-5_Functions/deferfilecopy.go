package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// let's log how long the program takes to run
	currentTime := time.Now()

	fmt.Println("Defer and FileCopy Testing")
	fileCopyStatus, err := fileCopy("go.mod", "go.mod.copy", currentTime)
	fmt.Println(fileCopyStatus, err)
	fmt.Printf("Duration of Execution: %s\n", time.Since(currentTime))
}

func fileCopy(sourceFileName, destFileName string, t time.Time) (string, error) {
	// open the source
	srcFileHandler, err := os.Open(sourceFileName)
	if err != nil {
		return "", fmt.Errorf("file Opening failed: %v", sourceFileName)
	}

	// close the source file handler
	defer func() {
		fmt.Printf("Closing Source File: %v\n", sourceFileName)
		srcFileHandler.Close()
	}()
	//defer srcFileHandler.Close() // close-2// will be called at last during the normal or error return

	// create the destination file

	dstFileHandler, err := os.Create(destFileName)
	if err != nil {
		return "", fmt.Errorf("cant create the destination file: %v", destFileName)
	}

	// close the destination file handler
	defer func() {
		fmt.Printf("Closing Destination File: %v\n", destFileName)
		dstFileHandler.Close()
	}()
	//defer dstFileHandler.Close() // close-1

	// Copy the file contents from sourceFile to destinationFile
	fileCopyStatus, err := func() (string, error) {
		fmt.Printf("Copying %v to %v\n", sourceFileName, destFileName)
		_, err = io.Copy(dstFileHandler, srcFileHandler)
		if err != nil {
			return "", fmt.Errorf("file Copying Erroor: [%v]->[%v]", sourceFileName, destFileName)
		}
		return "File Copied Successful\n", nil
	}()
	//fmt.Println("Sleeping for 50ms .....")
	//time.Sleep(50 * time.Millisecond)
	return fileCopyStatus, nil

}
