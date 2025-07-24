package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

func Execute_Linux(theCMD string) []string {
	CMD := SplitCommands(theCMD) // []string
	// Check if we have at least one element (the command name)
	if len(CMD) == 0 {
		fmt.Println("Error: empty command")
		return []string{}
	}
	var output []byte
	if len(CMD) == 1 {
		stdout := exec.Command(CMD[0])
		output, _ = stdout.Output()
	} else {
		stdout := exec.Command(CMD[0], CMD[1:]...)
		output, _ = stdout.Output()
	}
	return []string(strings.Split(string(output), "\n"))
}
