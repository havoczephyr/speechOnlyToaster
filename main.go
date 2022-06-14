package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// be careful with apostrophes, since this will be used in the filename.
const FIRSTNAME string = "Giovanni"
const LASTNAME string = "DAmico"

type TsvInfo struct {
	SessionName string
	TsvPath     string
}

func copy(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err2 := io.Copy(source, destination)
	return err2
}

func createExport(arr []TsvInfo, dir string, rootdir string) {
	now := time.Now()
	exportName := fmt.Sprintf("%s_%s_tsv_%d%d%d", FIRSTNAME, LASTNAME, now.Month(), now.Day(), now.Year())
	exportPath := filepath.Join(rootdir, exportName)
	err := os.Mkdir(exportPath, 0777)
	if err != nil {
		panic(err)
	}
	for _, item := range arr {
		itemPath := filepath.Join(exportPath, item.TsvPath)
		os.Mkdir(itemPath, 0777)
		copy(item.TsvPath, itemPath)
	}

}

func dirCheck(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
func tsvCheck(dir string) (string, error) {
	const tsvName string = "curated_processed_speech_only.tsv"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, file := range files {

		if file.Name() == tsvName {
			return filepath.Join(dir, tsvName), nil
		}
	}

	return "", nil
}

func processFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	tsvInfos := make([]TsvInfo, len(files))

	for _, file := range files {
		isSession := file.IsDir() && strings.HasPrefix(file.Name(), "session-")
		if isSession {
			filePath := filepath.Join(dir, file.Name())
			tsvPath, _ := tsvCheck(filePath)
			if tsvPath != "" {
				info := TsvInfo{file.Name(), tsvPath}
				tsvInfos = append(tsvInfos, info)
				createExport(tsvInfos, filePath, dir)
			}
		}
	}
}

func main() {
	dirArgs := os.Args[1]
	dirExists, _ := dirCheck(dirArgs)
	if dirExists {
		fmt.Println("it exists!")
		processFiles(dirArgs)
	} else {
		fmt.Printf("could not find directory %s", dirArgs)
	}
}
