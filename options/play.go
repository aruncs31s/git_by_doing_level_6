package options

import (
	"fmt"
	"strings"

	"level_6/git"
	"level_6/students"
)

func checkNames(chParser chan string) {
	chJsonParsor, chGitUserName := make(chan []students.Student), make(chan string)
	// TODO: Execute parsing json and geting user.name from git parallel
	go students.ReadStudentsFromJSON(chJsonParsor)
	go git.GetUserName(chGitUserName)
}

func Play() {
	ch_CheckNames, ch_counts := make(chan bool), make(chan )

	// NOTE: Check if their namess match
	go checkNames(ch_CheckNames)

	studentsList, _ := students.ReadStudentsFromJSON()

	// for _, student := range studentsList {
	// 	fmt.Println(student.Name)
	// }
	fmt.Printf("Total Number of Students %v \n", len(studentsList))

	// Check if they are in this list
	for _, student := range studentsList {
		if theirUserName == strings.TrimSpace(student.Name) {
			fmt.Println("They are in the list")
		}
	}
	Intro(theirUserName)
}
