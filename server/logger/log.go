package logger

import "fmt"

// Info logs a message in green
func Info(message string) {
	// Print the message in green
	fmt.Printf("\033[32m INFO | " + message + "\033[0m\n")
}

// Error logs a message in red
func Error(message string, err error) {
	// Print the message in red
	fmt.Printf("\033[31m ERROR | " + message + "\033[0m\n")
	fmt.Println(err)
}
