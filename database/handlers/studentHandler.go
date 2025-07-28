package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"level_6/database/models"
	"level_6/database/repository"
)

type StudentsHandler struct {
	repo repository.StudentsRepository
}

/*
Takes a repo of type repository.StudentsRepository
*/
func NewStudentsHandler(repo repository.StudentsRepository) *StudentsHandler {
	return &StudentsHandler{repo: repo}
}

func (h *StudentsHandler) CreateStudent() {
	fmt.Println("\nEnter new student details:")
	name, username := h.GetUserInput()

	student := &models.Students{Name: name, Username: username}
	if err := h.repo.Create(student); err != nil {
		fmt.Printf("Error creating student: %v\n", err)
	} else {
		fmt.Printf("Student created successfully with ID: %d\n", student.ID)
	}
}

func (h *StudentsHandler) GetAllStudents() {
	students, err := h.repo.GetAll()
	if err != nil {
		fmt.Printf("Error reading students: %v\n", err)
		return
	}

	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	fmt.Println("\nAll students:")
	for _, s := range students {
		fmt.Printf("ID: %d, Name: %s, Username: %s\n",
			s.ID, s.Name, s.Username)
	}
}

func (h *StudentsHandler) UpdateStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter student ID to update: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	student, err := h.repo.GetByID(uint(id))
	if err != nil {
		fmt.Printf("Student not found: %v\n", err)
		return
	}

	fmt.Printf("Current student: Name=%s, Username=%s\n", student.Name, student.Username)
	fmt.Println("Enter new details:")
	name, username := h.GetUserInput()

	student.Name = name
	student.Username = username

	if err := h.repo.Update(student); err != nil {
		fmt.Printf("Error updating student: %v\n", err)
	} else {
		fmt.Printf("Student ID %d updated successfully\n", student.ID)
	}
}

func (h *StudentsHandler) DeleteStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter student ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	// Check if student exists
	_, err = h.repo.GetByID(uint(id))
	if err != nil {
		fmt.Printf("Student not found: %v\n", err)
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		fmt.Printf("Error deleting student: %v\n", err)
	} else {
		fmt.Printf("Student ID %d deleted successfully\n", uint(id))
	}
}

func (h *StudentsHandler) GetUserInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter student username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	return name, username
}
