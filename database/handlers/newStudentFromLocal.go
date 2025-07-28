// TODO: Check where to place it?
package handlers

import (
	"fmt"

	"level_6/database/models"
	"level_6/git"

	"gorm.io/gorm"
)

func (h *StudentsHandler) CreateLocalStudent() {
	// TODO: Check if both name and username are needed
	chName, chUsername := make(chan string), make(chan string)

	go git.GetUsernameid(chUsername)
	go git.GetUserName(chName)
	username := <-chUsername

	// Check if the user already exists
	existingStudent, err := h.repo.GetByUsername(username)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("Error checking for existing user: %v\n", err)
		return
	}

	if existingStudent != nil {
		fmt.Printf("Log: User with username '%s' already exists.\n", username)
		return
	}

	name := <-chName
	// If user does not exist, create a new one
	student := &models.Students{Name: name, Username: username}
	if err := h.repo.Create(student); err != nil {
		fmt.Printf("Error creating student: %v\n", err)
	} else {
		fmt.Printf("Student created successfully with ID: %d\n", student.ID)
	}
}
