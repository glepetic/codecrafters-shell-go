package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadNext() string {
	return read("$", false)
}

func read(token string, isOpenQuote bool) string {
	fmt.Printf("%s ", token)

	scanner.Scan()
	userInput := scanner.Text()

	if len(userInput) == 0 {
		return read(token, isOpenQuote)
	}

	var openQuote = isOpenQuote
	var output strings.Builder
	for _, char := range userInput {
		if char == '"' {
			openQuote = !openQuote
		} else {
			output.WriteString(string(char))
		}
	}

	if openQuote {
		return output.String() + "\r\n" + read(">", openQuote)
	}

	return output.String()
}
