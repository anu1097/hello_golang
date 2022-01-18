package main

import (
	"fmt"

	"github.com/anu1097/golang_module/greetings_1"
	"github.com/anu1097/golang_module/greetings_2"
)

func greeting_1() {
	message := greetings_1.Hello("Gladys")
	fmt.Println(message)
}

func greeting_2() {
	message := greetings_2.Hello("Gladys")
	fmt.Println(message)
}

func main() {
	// Get a greeting message and print it.
	greeting_1()
	greeting_2()
}
