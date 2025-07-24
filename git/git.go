package git

import (
	"fmt"
	"os/exec"
)

func GetNewFiles() []string {
	stdout := exec.Command("ls")
	fmt.Println(stdout)
	return []string{"", ""}
}
