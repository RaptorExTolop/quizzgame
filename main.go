package main

import (
	"log"
	"strings"
)

func main() {
	questions := Questions{}
	var blankStringList []string

	for {
		userInput, err := getInput("Command > ")
		if err != nil {
			log.Fatal("error getting command:", err)
			continue
		}

		userInput = strings.Replace(userInput, "\n", "", -1)

		switch userInput {
		case "add":
			questions.add("testing", blankStringList, blankStringList)
		case "delete":
			questions.delete(1)
		case "print":
			questions.printQuestions()
		}
	}
}
