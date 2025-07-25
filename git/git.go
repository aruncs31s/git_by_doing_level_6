package git

import (
	"strconv"
	"strings"

	"level_6/git/helper"
	"level_6/git/helper/commands"
	"level_6/git/helper/remote"
)

// func GitMain() {
// fmt.Println(GetNewFiles())
// fmt.Println(len(GetNewFiles()))
// }

func GetModifiedFiles() []string {
	return helper.Execute_Linux(commands.ModifiedFilesCMD)
}

func GetUntrackedFiles() []string {
	return helper.Execute_Linux(commands.UntrackedFilesCMD)
}

func GetNumCommitsLocal() int {
	_str := helper.Execute_Linux(commands.GetNumCommitsLocalCMD)
	number, _ := strconv.Atoi(_str[0])
	return number
}

func GetNumCommitsRemote() int {
	return remote.GetGithubCOmmits()
}

func GetUserName(chName chan string) {
	nameSlice := helper.Execute_Linux(commands.GetUserNameCMD)
	chName <- strings.Join(nameSlice, "")
}

func GetUserEmail() string {
	emailSlice := helper.Execute_Linux(commands.GetUserEmailCMD)
	return strings.Join(emailSlice, "")
}
