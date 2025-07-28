package options

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"level_6/git"
	"level_6/initializers"
	qanda "level_6/qAndA"
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

	initializers.App.Run()

	fmt.Println("You have succesfully marked your attendance.")

	defer outro()
}

func outro() {
	chCheckUpdate := make(chan bool)
	go git.CheckForUpdate(chCheckUpdate)
	fmt.Println("Now you want to push your attendance to the remote repository.")
	fmt.Println("Did you finish Pushing? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	answer := qanda.ReadYesNoAnswer(reader)

	if answer {
		return
	}
	if !<-chCheckUpdate {
		fmt.Println("Please update the repo first")
	}
}
