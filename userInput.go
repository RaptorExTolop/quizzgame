package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func addInput(qs *Questions) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the question: ")
	questionName, err := reader.ReadString('\n')
	fmt.Println(strings.Replace(questionName, "\n", "", -1))
	if err != nil {
		return err
	}
	fmt.Print("What are the possible answers: ")
	possibleAnswers, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	splittedPAns := strings.Split(possibleAnswers, " | ")
	for i := range len(splittedPAns) {
		fmt.Printf("P ANS -> %v\n", strings.TrimSpace(splittedPAns[i]))
	}
	fmt.Print("What are the correct answer/s: ")
	correctAnswers, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	splittedCAns := strings.Split(correctAnswers, "|")
	for i := range correctAnswers {
		fmt.Printf("C And -> %v\n", strings.TrimSpace(splittedCAns[i]))
	}

	return nil
}
