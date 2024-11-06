package main

import (
	"bufio"
	"errors"
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
	if err != nil {
		return err
	}
	questionNameRL := strings.ReplaceAll(questionName, "\n", "")
	fmt.Printf("%v\n", questionNameRL)

	fmt.Print("What are the possible answers a user can chose from? ")
	possibleAnswers, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	possibleAnswersSplit := strings.Split(possibleAnswers, "|")
	for i := range possibleAnswersSplit {
		fmt.Println(strings.TrimSpace(possibleAnswersSplit[i]))
	}

	fmt.Print("What are the corect answer/s: ")
	correctAnswers, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	correctAnswersSplit := strings.Split(correctAnswers, "|")
	for i := range correctAnswersSplit {
		fmt.Println(strings.TrimSpace(correctAnswersSplit[i]))
	}

	var possibleAnswersSplitRS, correctAnswersSplitRS []string

	var corAnsInPosAns bool = false
	for i := range possibleAnswersSplit {
		possibleAnswersSplitRS = append(possibleAnswersSplitRS, strings.TrimSpace(possibleAnswersSplit[i]))
		if len(correctAnswersSplit) == 1 {
			if correctAnswersSplit[0] == possibleAnswersSplit[i] {
				corAnsInPosAns = true
				correctAnswersSplitRS[0] = strings.TrimSpace(correctAnswersSplit[i])
				break
			}
		} else {
			for n := range correctAnswersSplit {
				if correctAnswersSplit[n] == possibleAnswersSplit[i] {
					correctAnswersSplitRS = append(correctAnswersSplitRS, strings.TrimSpace(correctAnswersSplit[n]))
					corAnsInPosAns = true
					break
				}
			}
		}
	}

	if !corAnsInPosAns {
		fmt.Println("No correct answer is in the possible answers")
		err := errors.New("no correct answer is in the possible answers")
		return err
	}

	qs.add(questionNameRL, possibleAnswersSplitRS, correctAnswersSplitRS)
	return nil
}
