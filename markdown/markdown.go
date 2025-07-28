package markdown

import (
	"fmt"
	"io/ioutil"
	"level_6/markdown/helper"
	"log"
	"os"
)

// Usage:
// func main() {
// 	MarkAttendance("aruncs", "Arun CS", true, true, false, true, false, true, true)
// }

func MarkAttendance(username, name string) {
	attendance := helper.AskForAttendance()
	// Read the markdown ,
	// Enter attentance
	// Ask should i push it ?
	// If yes, push it to the remote repository

	// Read the README.md file
	readmeContent, err := readFile("./README.md")
	if err != nil {
		log.Fatalf("Failed to read README file: %v", err)
	}

	// Print the content of the README file
	fmt.Println("README Content:")
	contents := string(readmeContent)
	fmt.Println(contents)
	// Create the attendance entry
	attendanceEntry := fmt.Sprintf("|%s| [%s](https://github.com/%s) | %s | %s | %s | %s | %s | %s |%s |\n",
		name, name, username, checkmark(attendance[0]), checkmark(attendance[1]), checkmark(attendance[2]), checkmark(attendance[3]), checkmark(attendance[4]), checkmark(attendance[5]), checkmark(attendance[6]))
	fmt.Println(attendanceEntry)
	updatedContent := contents + attendanceEntry
	fmt.Println("Updated README Content:")
	fmt.Println(updatedContent)
	// Write that to the file
	err = os.WriteFile("README.md", []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write to README file: %v", err)
	}
	fmt.Println("Attendance entry added successfully to README.md")
}

func checkmark(condition bool) string {
	if condition {
		return "✅"
	}
	return "❌"
}

// readFile reads the content of a file and returns it as a byte slice
func readFile(filename string) ([]byte, error) {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filename)
	}

	// Read file content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}
