package greetings

import (
	"fmt"
	 "errors"
	 "math/rand"
	 "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greetimg that embeds the name in message.
	if name =="" {
		return "", errors.New("empty name") 
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi %v", 
		"Merhaba %v",
		"hello %v",
	}

	return formats[rand.Intn(len(formats))]
}