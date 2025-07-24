package git

import (
	"level_6/git/helper"
	"level_6/git/helper/commands"
	"level_6/git/helper/remote"
	"strconv"
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


