package main

import (
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		var userInput string
		n := readNext(&userInput)
		var command string
		if n > 1 {
			command = strings.Split(userInput, " ")[0]
		} else {
			command = userInput
		}
		fmt.Printf("%s: command not found\r\n", command)
	}
}

func readNext(input *string) int {
	fmt.Print("$ ")
	n, err := fmt.Scan(input)
	if err != nil {
		fmt.Printf("Internal error when scanning user input: %s", err)
		os.Exit(-1)
	}
	return n
}
