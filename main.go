package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile() string {
	c, err := os.ReadFile("helloWorld.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(c)
}

func splitFileWithParams(fc string, rc []string) []string {
	var finalString []string
	for r := 0; r < len(rc); r++ {
		finalString = strings.Split(fc, rc[r])
	}
	return finalString
}

func main() {
	fileContents := readFile()
	fileContentsString := string(fileContents)
	splittedFile := splitFileWithParams(fileContentsString, nil)
	for i := 1; i < len(splittedFile); i++ {
		fmt.Println(splittedFile[i])
	}
}
