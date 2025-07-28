package git

import (
	"level_6/git/helper"
	"level_6/git/helper/commands"
)

func checkForUpdate(ch_isUpdate chan bool) {
	chLocal := make(chan []string)
	go helper.ExecuteLinux(chLocal, commands.GetNumCommitsLocalCMD)
}
