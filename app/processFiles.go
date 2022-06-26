package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ProcessFiles(dir string) []TsvInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("read failure: %s", err)
		panic(err)
	}
	tsvInfos := make([]TsvInfo, 0)

	for _, file := range files {
		isSession := file.IsDir() && strings.HasPrefix(file.Name(), "session-")
		if isSession {
			filePath := filepath.Join(dir, file.Name())
			tsvPath, _ := TsvCheck(filePath)
			if tsvPath != "" {
				fmt.Printf("adding tsv session name %s and tsvPath %s \n", file.Name(), tsvPath)
				info := TsvInfo{file.Name(), tsvPath}
				tsvInfos = append(tsvInfos, info)

			}
		}
	}
	fmt.Printf("length of tsvInfos: %d \n", len(tsvInfos))
	return tsvInfos
}
