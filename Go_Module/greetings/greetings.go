package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// RandomFormat returns one of a set of greeting messages. The returned message is selected at random

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! well met!",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats

	return formats[rand.Intn(len(formats))]
}
