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

func Hellos(names []string) (map[string]string, error) {
	msgMap := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgMap[name] = msg
	}
	return msgMap, nil
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
