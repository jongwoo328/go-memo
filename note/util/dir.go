package util

import (
	"fmt"
	"os"
	"os/user"
)

func GetMemoDirOrCreate() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("error getting current user: %w", err)
	}

	homeDir := currentUser.HomeDir
	appDir := homeDir + "/.go-memo"
	memoDir := homeDir + "/.go-memo/memos"
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		err := os.Mkdir(appDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("error creating app directory: %w", err)
		}

		err = os.Mkdir(memoDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("error creating memo directory: %w", err)
		}
	}

	return memoDir, nil
}
