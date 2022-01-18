package greetings_1

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Greetings_1, %v. Welcome!", name)
	return message
}
