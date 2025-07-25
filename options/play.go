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
	isSameName := <-chCheckNames
	if !isSameName {
		log.Fatal("Configure You git user.name first")
	}
	if !<-isSameNoOfCommits {
		fmt.Println("Please update the repo")
		return
	}
	fmt.Println("You have succesfully Completed this level")
}
