package main

import "interface_polymorphism_logger/implementation"

func main() {
	// create a ConsoleLogger and FileLogger
	consoleLogger := implementation.ConsoleLogger{}              // implemented the interface
	fileLogger := implementation.FileLogger{FileName: "log.txt"} // implemented the interface

	// Use the LogMessage()
	implementation.LogMessage(consoleLogger, "This is a console logger")
	implementation.LogMessage(fileLogger, "this is a file logger")
}
