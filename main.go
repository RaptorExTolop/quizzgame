package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileContents, err := os.ReadFile("helloWorld.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	splittedFile := strings.Split(string(fileContents), "\n")
	for i := 0; i < len(splittedFile); i++ {
		fmt.Println(splittedFile[i])
	}
}
