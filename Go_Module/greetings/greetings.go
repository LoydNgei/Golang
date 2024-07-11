package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// The := operator is a shortcut for declaring and initializing a variable in one line
