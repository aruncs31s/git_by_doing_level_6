// WARNING: UserName = "Arun CS"
// Username = "aruncs31s" , there is a typo
// TODO: Fix this latter.

// Package git ...NOTE: To avoid linter warnings
package git

import (
	"log"
	"strconv"
	"strings"

	"level_6/git/helper"
	"level_6/git/helper/commands"
	"level_6/git/helper/remote"
	"level_6/students"
)

// func GitMain() {
// fmt.Println(GetNewFiles())
// fmt.Println(len(GetNewFiles()))
// }

func GetModifiedFiles(ch chan []string) {
	// TODO: Check if i can avoid this nesting
	chNew := make(chan []string)
	go helper.ExecuteLinux(chNew, commands.ModifiedFilesCMD)
	ch <- <-chNew // ??
}

func GetUntrackedFiles(ch chan []string) {
	chNew := make(chan []string)
	go helper.ExecuteLinux(chNew, commands.UntrackedFilesCMD)
	ch <- <-chNew
}

func GetNumCommitsLocal(ch chan int) {
	chNew := make(chan []string)
	go helper.ExecuteLinux(chNew, commands.GetNumCommitsLocalCMD)
	// number, _ := strconv.Atoi(_str[0])
	// return number
	stdout := <-chNew
	count, err := strconv.Atoi(stdout[0])
	if err != nil {
		log.Fatal("something bad happend")
	}
	ch <- count
}

func GetNumCommitsRemote(ch chan int) {
	ch <- remote.GetGithubCommits()
}

func GetUserName(chName chan string) {
	ch := make(chan []string)
	go helper.ExecuteLinux(ch, commands.GetUserNameCMD)
	chName <- strings.Join(<-ch, "")
}

func GetUserEmail(chEmail chan string) {
	ch := make(chan []string)
	go helper.ExecuteLinux(ch, commands.GetUserEmailCMD)
	chEmail <- strings.Join(<-ch, "")
}

func CheckForUpdate(chIsUpdate chan bool) {
	// TODO: Compare local and remote commits and return true/false
	chLocal, chRemote := make(chan int), make(chan int)

	go GetNumCommitsLocal(chLocal)
	go GetNumCommitsRemote(chRemote)
	if <-chLocal != <-chRemote {
		// fmt.Println("Please Update this repo" )
		chIsUpdate <- false
		return
	}
	chIsUpdate <- true
}

func checkForUpdate(ch_isUpdate chan bool) {
	chLocal := make(chan []string)
	go helper.ExecuteLinux(chLocal, commands.GetNumCommitsLocalCMD)
}

func CheckNames(chChekName chan bool) {
	// Debug
	chJSONParsor, chGitUserName := make(chan []students.Student), make(chan string)
	// TODO: Execute parsing json and geting user.name from git parallel
	go students.ReadStudentsFromJSON(chJSONParsor)
	go GetUserName(chGitUserName)
	gitUserName := <-chGitUserName
	// fmt.Println(gitUserName)
	// fmt.Println(ch)
	for _, student := range <-chJSONParsor {
		if strings.Contains(student.Name, gitUserName) {
			chChekName <- true
			break
		} else {
			continue
		}
	}
	chChekName <- false
}

// NOTE: GetUsernameid <- "aruncs31s"
func GetUsernameid(chUsername chan string) {
	chJSONParsor, chGitUserName := make(chan []students.Student), make(chan string)
	// TODO: Execute parsing json and geting user.name from git parallel
	go students.ReadStudentsFromJSON(chJSONParsor)
	go GetUserName(chGitUserName)
	gitUserName := <-chGitUserName
	// fmt.Println(gitUserName)
	// fmt.Println(ch)
	for _, student := range <-chJSONParsor {
		if strings.Contains(student.Name, gitUserName) {
			chUsername <- student.Username
			break
		} else {
			continue
		}
	}
	chUsername <- ""
}

// func PushToRemoteRepo() {
// // func GetModifiedFiles(ch chan []string) {
// 	// TODO: Check if i can avoid this nesting

// 	chNew := make(chan []string)
// 	go helper.ExecuteLinux(chNew, commands.ModifiedFilesCMD)
// 	go helper.ExecuteLinux(chNew, commands.ModifiedFilesCMD)
// 	go helper.ExecuteLinux(chNew, commands.ModifiedFilesCMD)
// 	go helper.ExecuteLinux(chNew, commands.ModifiedFilesCMD)

// 	// ch <- <-chNew // ??
// 	/*
// 		value := <-chNew
// 		ch <- value
// 	*/
// // }

// 	helper.ChainExecute(
// 		commands.GitPullCMD,
// 		commands.GitAddCMD,
// 		commands.GitCommitCMD,
// 		commands.GitPushCMD,
// 	)

// }
