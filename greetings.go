package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo a una persona especifica.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Welcome, %v!",
		"Hello, %v. Welcome!",
		"Hi, %v! Welcome!",
		"What's up, %v! Welcome!",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}
