package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf(randomFormat(), name), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	format := []string{
		"Hello %s, welcome",
		"Hey %s, Good to see you",
		"Holla %s",
	}
	return format[rand.Intn(len(format))]
}
