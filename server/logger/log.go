package logger

import "fmt"

// Info logs a message in green
func Info(message string) {
	// Print the message in green
	fmt.Printf("\033[32m INFO | " + message + "\033[0m\n")
}
