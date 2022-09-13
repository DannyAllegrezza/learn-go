package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "dang", "****")
	*message = strings.ReplaceAll(*message, "heck", "****")
	*message = strings.ReplaceAll(*message, "shoot", "*****")
}

// don't touch below this line

func main() {
	messages := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}
