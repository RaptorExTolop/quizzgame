package main

import (
	"fmt"
	"os"
)

func main() {
	questions := Questions{}

	var blankStringList []string

	questions.add("testing 1", blankStringList, blankStringList)
	questions.add("testing 2", blankStringList, blankStringList)
	questions.add("testing 3", blankStringList, blankStringList)
	err := questions.editName(2, "testing 3.0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = questions.delete(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(questions)
}
