package options

import (
	"errors"
	"fmt"
	"log"

	"level_6/git"
	"level_6/initializers"
)

// initialize performs pre-flight checks and sets up the application.
func initialize() error {
	chCheckNames, isSameNoOfCommits := make(chan bool), make(chan bool)

	/*

	 */
	go git.CheckNames(chCheckNames)
	go git.CheckForUpdate(isSameNoOfCommits)

	// Wait for checks to complete
	if !<-chCheckNames {
		return errors.New("Configure You git user.name first, check level 5")
	}
	if !<-isSameNoOfCommits {
		return errors.New("Please update the repo")
	}

	// app := <-chDatabase
	return nil
}

func Play() {
	err := initialize()
	if err != nil {
		log.Fatal(err)
	}

	defer fmt.Println("You have succesfully Completed this level")
	initializers.App.Run()
}
