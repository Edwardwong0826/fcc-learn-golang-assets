package main

import (
	"fmt"
	"strings"
)

// & operator generates a pointer
// * dereferences a pointer to gain access to the value
func removeProfanity(message *string) {
	// ?
	var value string
	value = *message
	value = strings.ReplaceAll(value, "dang", "****")
	value = strings.ReplaceAll(value, "shoot", "*****")
	value = strings.ReplaceAll(value, "heck", "****")
	*message = value

}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}
