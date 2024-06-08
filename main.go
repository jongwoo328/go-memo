package main

import (
	"fmt"
	"go-memo/note"
	"os"
)

func main() {
	arguments := os.Args
	subCommands := arguments[1:]
	if len(subCommands) < 1 {
		fmt.Println("No subcommand")
		return
	}

	subCommand := subCommands[0]
	switch subCommand {
	case "add":
		note.Add()
	case "list":
		note.List()
	default:
		fmt.Println("Invalid subcommand")
	}

}
