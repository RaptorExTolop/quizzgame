package main

import (
	"fmt"
	"strings"
)

func main() {
	questions := Questions{}
	//var blankStringList []string

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
			addInput(&questions)
		case "delete":
			questions.delete(1)
		case "print":
			questions.printQuestions()
		default:
			fmt.Println("Unknown command")
		}
	}
}
