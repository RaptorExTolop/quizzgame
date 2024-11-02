package main

import "fmt"

func main() {
	questions := Questions{}

	var blankStringList []string

	questions.add("testing 1", "this is test 1", blankStringList, blankStringList)
	questions.add("testing 2", "this is test 2", blankStringList, blankStringList)
	questions.add("testing 3", "this is test 3", blankStringList, blankStringList)

	questions.delete(1)

	fmt.Println(questions)
}
