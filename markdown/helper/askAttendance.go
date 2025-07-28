package helper

import (
	"bufio"
	"fmt"
	qanda "level_6/qAndA"
	"os"
)

func AskForAttendance() []bool {
	answers := make([]bool, 7)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nDid You complete Level 0:")
	answers[0] = qanda.ReadYesNoAnswer(reader)
	fmt.Println("\nDid You complete Level 1:")
	answers[1] = qanda.ReadYesNoAnswer(reader)
	fmt.Println("\nDid You complete Level 2:")
	answers[2] = qanda.ReadYesNoAnswer(reader)
	fmt.Println("\nDid You complete Level 3:")
	answers[3] = qanda.ReadYesNoAnswer(reader)
	fmt.Println("\nDid You complete Level 4:")
	answers[4] = qanda.ReadYesNoAnswer(reader)
	fmt.Println("\nDid You complete Level 5:")
	answers[5] = qanda.ReadYesNoAnswer(reader)
	fmt.Println("\nDid You complete Level 6:")
	answers[6] = qanda.ReadYesNoAnswer(reader)
	return answers
}
