package options

import (
	"fmt"
	"log"

	"level_6/database"
	"level_6/git"
	helper "level_6/helper"
)

func Play() {
	chCheckNames, isSameNoOfCommits := make(chan bool), make(chan bool)
	chDatabase := make(chan *database.App)
	defer fmt.Println("You have succesfully Completed this level")
	// NOTE: Check if their namess match

	go helper.CheckNames(chCheckNames)
	go git.CheckForUpdate(isSameNoOfCommits)
	go database.Database(chDatabase)

	if !<-chCheckNames {
		log.Fatal("Configure You git user.name first")
	}
	if !<-isSameNoOfCommits {
		log.Fatal("Please update the repo")
	}

	app := <-chDatabase
	// fmt.Println("Inside play after calling database")
	app.Run()
}
