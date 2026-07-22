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
		fmt.Print("$ ")
		n, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Printf("Internal error when scanning user input: %s", err)
			os.Exit(-1)
		}
		var command string
		if n > 1 {
			command = strings.Split(userInput, " ")[0]
		} else {
			command = userInput
		}
		fmt.Printf("%s: command not found\r\n", command)
	}
}
