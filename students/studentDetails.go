// Package students .. NOTE:to avoid linter error
package students

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
type Students struct {
	Students []Student `json:"students"`
}

func ReadStudentsFromJSON(chParser chan []Student) {
	data, err := os.ReadFile("students/students_details.json")
	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}
	// Parse as map[string]string (username -> name)
	var studentMap map[string]string
	err = json.Unmarshal(data, &studentMap)
	if err != nil {
		fmt.Printf("error parsing JSON: %v", err)
	}

	// Convert to []Student slice
	var students []Student
	for username, name := range studentMap {
		students = append(students, Student{
			Name:     name,
			Username: username,
		})
	}

	chParser <- students
	// return students, nil
}

func DisplayStudents(students []Student) {
	fmt.Println("📚 Student Details:")

	for _, student := range students {
		fmt.Printf("   Name: %s\n", student.Name)
		fmt.Printf("   username: %s\n", student.Username)
	}
}
