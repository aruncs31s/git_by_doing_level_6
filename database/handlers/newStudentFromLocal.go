// TODO: Check where to place it?
package handlers

import (
	"fmt"

	"level_6/database/models"
	"level_6/git"
)

func (h *StudentsHandler) CreateLocalStudent() {
	fmt.Println("\nEnter new student details:")

	// TODO: Check if both name and username are needed
	chName, chUsername := make(chan string), make(chan string)

	go git.GetUsernameid(chUsername)
	go git.GetUserName(chName)
	name, username := <-chName, <-chUsername

	// Check if the user is already exist in the database
	// if not create one
	// if yes return false / print alredy exisit
	// NOTE: better automatically check this in the beggining
	student := &models.Students{Name: name, Username: username}
	if err := h.repo.Create(student); err != nil {
		fmt.Printf("Error creating student: %v\n", err)
	} else {
		fmt.Printf("Student created successfully with ID: %d\n", student.ID)
	}
}
