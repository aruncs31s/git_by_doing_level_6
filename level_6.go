package main

import (
	"fmt"
	"level_6/git"
)

func main() {
	// Check if any file has changed

	modifiedFiles := git.GetModifiedFiles()
	for _, file := range modifiedFiles {
		if len(file) != 0 {
			fmt.Printf("Modified: %v\n", file)

		}
	}

	// Untracked files
	UntrackedFiles := git.GetUntrackedFiles()
	for _, file := range UntrackedFiles {
		if len(file) != 0 {
			fmt.Printf("Untracked: %v\n", file)

		}
	}
	// remote.GetGithubCOmmits()
	// Total Local Commits.
	totalLocalCommit := git.GetNumCommitsLocal()
	fmt.Print("Total Local Commits ")
	fmt.Println(totalLocalCommit)
	totalRemoteCommit := git.GetNumCommitsRemote()

	fmt.Print("Total Remote Commits ")
	fmt.Println(totalRemoteCommit)
}
