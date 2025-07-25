package options

import (
	"fmt"
	"log"

	"level_6/git"
	helper "level_6/helper"
)

func Play() {
	chCheckNames, isSameNoOfCommits := make(chan bool), make(chan bool)

	// NOTE: Check if their namess match
	go helper.CheckNames(chCheckNames)
	go git.CheckForUpdate(isSameNoOfCommits)

	if !<-chCheckNames {
		log.Fatal("Configure You git user.name first")
	}
	if !<-isSameNoOfCommits {
		log.Fatal("Please update the repo")
	}
	fmt.Println("You have succesfully Completed this level")
}
