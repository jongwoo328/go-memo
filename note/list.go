package note

import (
	"fmt"
	"go-memo/note/util"
	"os"
	"sort"
)

func List() (int, error) {
	memoDir, err := util.GetMemoDirOrCreate()
	if err != nil {
		fmt.Println("Error: ", err)
		return 1, err
	}

	files, err := os.ReadDir(memoDir)
	if err != nil {
		fmt.Println("Error: ", err)
		return 1, err
	}

	sort.Slice(files, func(i, j int) bool {
		fileInfo, err := files[i].Info()
		if err != nil {
			return false
		}
		fileInfo2, err := files[j].Info()
		if err != nil {
			return false
		}
		return fileInfo.ModTime().After(fileInfo2.ModTime())
	})

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Println("Error: ", err)
			return 1, err
		}
		fmt.Println(fmt.Sprintf("\033[90m%s\033[0m\t", info.ModTime().Format("2006-01-02 15:04:05")), file.Name())
	}

	return 0, nil
}
