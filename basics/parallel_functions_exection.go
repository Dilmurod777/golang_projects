package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func parallelFunctionsExecution(functions []func() error, N int, maxErrors int) {
	channelErrors := make(chan error, maxErrors)

	for _, function := range functions {
		go func(function func() error) {
			channelErrors <- function()

			if len(channelErrors) == cap(channelErrors) {
				os.Exit(0)
			}
		}(function)

		time.Sleep(1 * time.Microsecond) //
	}
}

func main() {
	functions := []func() error{
		func() error {
			fmt.Println("Function 1")
			return errors.New("function 1")
		},
		func() error {
			fmt.Println("Function 2")
			return errors.New("function 2")
		},
		func() error {
			fmt.Println("Function 3")
			return errors.New("function 3")
		},
		func() error {
			fmt.Println("Function 4")
			return errors.New("function 4")
		},
		func() error {
			fmt.Println("Function 5")
			return errors.New("function 5")
		},
	}

	parallelFunctionsExecution(functions, len(functions), 2)

	time.Sleep(2 * time.Second)
}
