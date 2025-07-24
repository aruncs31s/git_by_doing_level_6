package helper

import (
	"strings"
)

// NOTE: Initially copying the values.
func SplitCommands(theCMD string) []string {
	splitCMD := strings.Split(strings.TrimSpace(theCMD), " ")
	return splitCMD
}
