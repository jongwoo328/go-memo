package note

import (
	"fmt"
	"github.com/google/uuid"
	"go-memo/note/util"
	"os"
	"os/exec"
)

func Add() {
	var fileName string
	if len(os.Args) >= 3 {
		fileName = os.Args[2]
	} else {
		uuidRandom, err := uuid.NewRandom()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fileName = uuidRandom.String()
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	editorCommand, editorOptions := util.ParseCommand(editor)

	memoDir, err := util.GetMemoDirOrCreate()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	editorOptions = append(editorOptions, fmt.Sprintf("%s/%s.md", memoDir, fileName))
	cmd := exec.Command(editorCommand, editorOptions...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
