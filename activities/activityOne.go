package activities

import (
	"fmt"
	"strings"
)

// Activity 1. Fork a repocitory.



func makeThemFork() bool {
	fmt.Println("Did you fork this repo? (y/n)")
	var ansOne string
	fmt.Scan(&ansOne)
	return strings.ToLower(ansOne) == "y"
}


func ActivityOne() {
	fmt.Println("Activity 1. You have to go to github and fork this repocitory 👍 ")
	if makeThemFork() {
		
}
