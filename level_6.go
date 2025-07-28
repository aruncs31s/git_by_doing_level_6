package main

import (
	"log"

	"level_6/cmd"
	"level_6/options"
)

/* NOTE: Flow
- check which option using ./cmd/commandLine.go
- play {}
	- Check if they played the previous level
		- it can be done by checking if their git config user.name is set,
		- if the *getUserName* returns something , check if its their actual name from the ./students/students_details.json file
	- Then check if there is any update to the repo , if yes promt them to update,and restart the level(wait for it to fetch and continue)
*/

func main() {
	switch selectedOption := cmd.WhichCMD(); selectedOption {
	case options.SelectError:
		// TODO: Check if i can be more specific about the err
		log.Fatal("Error")
	case options.SelectDoPlay:
		options.Play()
	case options.SelectShowError:
		options.ShowErrorDetails("test error")
	default:
		log.Fatal("Error")
	}
}
