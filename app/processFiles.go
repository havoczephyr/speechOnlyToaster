package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//sweeps through working directory for all folders that has the prefix of "session-" and utlizes TsvCheck() to verify
//if it contains "curated_processed_speech_only.tsv". if folder satisfies all requirements, its path and its name is placed in tsvInfos
//and is provided as a return.
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
