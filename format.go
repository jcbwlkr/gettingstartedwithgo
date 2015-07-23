package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func Notify(u User) (string, error) {
	if u.Age < (72 / 4) {
		return "", errors.New("too young")
	}

	return fmt.Sprintf("Hello, %s!", u.Name), nil
}
