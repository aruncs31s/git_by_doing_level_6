// Package cmd is used to ... // NOTE: to ignore linter error
package cmd

import (
	"log"
	"os"

	"level_6/options"
)

func WhichCMD() uint8 {
	// TODO: Convert it to Switch
	// Check if they provided args
	if len(os.Args) < 2 {
		log.Fatal("    No args")
	}

	if os.Args[1] == "" {
		log.Fatal("    Empty Arg 1")
	}
	// For Linux
	if os.Args[1] == "play" {
		return options.DoPlay
	}

	if os.Args[1] == "show" {
		if os.Args[2] == "help" {
			return options.Help
		} else if os.Args[2] == "error" {
			return options.ShowError
		} else {
			log.Panic("   given ", os.Args[1])
			log.Panic("   expects ", os.Args[1])
			log.Fatal("   try show help")
		}
	}
	return options.Error
}
