package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(message string) (userInput string, err error) {
	/* reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return "error"
	}

	return name*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v", message)

	userInput, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return userInput, nil
}
