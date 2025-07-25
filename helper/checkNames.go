package helper

import (
	"strings"

	"level_6/git"
	"level_6/students"
)

func CheckNames(chChekName chan bool) {
	// Debug
	chJSONParsor, chGitUserName := make(chan []students.Student), make(chan string)
	// TODO: Execute parsing json and geting user.name from git parallel
	go students.ReadStudentsFromJSON(chJSONParsor)
	go git.GetUserName(chGitUserName)
	gitUserName := <-chGitUserName
	// fmt.Println(gitUserName)
	// fmt.Println(ch)
	for _, student := range <-chJSONParsor {
		if strings.Contains(student.Name, gitUserName) {
			chChekName <- true
			break
		} else {
			continue
		}
	}
}
