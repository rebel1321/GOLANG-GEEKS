package main

import (
	"errors"
	"fmt"
)
var ErrInvalidFile = errors.New("Invalid file")
func check(num int) (string, error) {
	if num < 0 {
		return "", errors.New("Number is negative")
	}
	return "Number is positive", nil
}

func read(filename string) error {
	if filename == "" {
		return ErrInvalidFile
	}
	return nil
}

func file() (string, error) {
	e := read("")
	if e != nil {
		return "", fmt.Errorf("Error reading file: %w", e)

	}
	return "File exists", nil
}
func main() {
	// Basic error handling example
	result, err := check(-1)
	if err != nil {
		println(err.Error())
	} else {
		println(result)
	}
//Wrapping of errors
	_, e := file()
	if e != nil {
		fmt.Println(e)
		we := errors.Unwrap(e)
		if errors.Is(e, ErrInvalidFile) {
			fmt.Println("Unwrapped error:", we)
		}
	}
}