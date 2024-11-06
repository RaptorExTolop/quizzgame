package main

import (
	"fmt"
	"strings"
)

func main() {
	questions := Questions{}

	testingArrayString := [2]string{"hello", "world"}
	for {
		userInput, err := getInput("Command -> ")
		if err != nil {
			fmt.Println("error getting command:", err)
			continue
		}

		userInput = strings.Replace(userInput, "\n", "", -1)

		switch userInput {
		case "add":
			//questions.add("testing", blankStringList, blankStringList)
			err := addInput(&questions)
			if err != nil {
				fmt.Printf("error adding command: %v\n", err)
			}
		case "delete":
			questions.delete(1)
		case "print":
			questions.printQuestions()
		case "add debug":
			questions.add("Testing", testingArrayString[:], testingArrayString[:])
			questions.add("Testing", testingArrayString[:], testingArrayString[:])
			questions.add("Testing", testingArrayString[:], testingArrayString[:])
			questions.add("Testing", testingArrayString[:], testingArrayString[:])
			questions.add("Testing", testingArrayString[:], testingArrayString[:])
		default:
			fmt.Println("Unknown command")
		}
	}
}
