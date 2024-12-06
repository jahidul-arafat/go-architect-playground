package implementation

import "fmt"

// logs message to the console
type ConsoleLogger struct {
}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console: ", message)
}

func (c ConsoleLogger) PrintMe() {
	fmt.Println("Print Me>>[C]>>")
}
