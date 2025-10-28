package main

import "fmt"

func exer5_19() {
	type expectedError struct{}
	defer func() {
		switch err := recover(); err {
		case nil:
			fmt.Println("no error")
		case expectedError{}:
			fmt.Println("expected error")
		}
	}()

	panic(expectedError{})
}
