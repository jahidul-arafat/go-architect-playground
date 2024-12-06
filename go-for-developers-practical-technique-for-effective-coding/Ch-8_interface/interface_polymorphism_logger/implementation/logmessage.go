package implementation

import "interface_polymorphism_logger/interfaces"

// LogMessage
// will behave like polymorphic since could log both the ConsoleMessage and the FileMessage
func LogMessage(loggerInterface interfaces.Logger, message string) {
	loggerInterface.Log(message)
	loggerInterface.PrintMe()
}
