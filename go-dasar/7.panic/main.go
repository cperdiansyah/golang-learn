package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		defer fmt.Println("terbaca")
		panic(err.Error())
		fmt.Println("end")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
