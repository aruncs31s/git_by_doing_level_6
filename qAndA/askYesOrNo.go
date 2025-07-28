package qanda

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadYesNoAnswer(reader *bufio.Reader) bool {
	for {
		fmt.Print("Please enter 'yes' or 'no': ")
		answer, _ := reader.ReadString('\n')
		answer = answer[:len(answer)-1] // Remove the newline character
		answer = answer[:1]             // Get the first character

		switch strings.ToLower(answer) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Input error")
		}
	}
}
