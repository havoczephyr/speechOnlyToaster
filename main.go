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
	fmt.Printf("copying from %s to %s", src, dst)
	_, err2 := io.Copy(source, destination)
	return err2
}

func createExport(tsvArr []TsvInfo, rootdir string, firstName string, lastName string) {
	now := time.Now()
	exportName := fmt.Sprintf("%s_%s_tsv_%d%d%d", firstName, lastName, now.Month(), now.Day(), now.Year())
	exportPath := filepath.Join(rootdir, exportName)
	err := os.Mkdir(exportPath, 0777)
	if err != nil {
		fmt.Printf("could not make dir: %s", err)
		panic(err)
	}
	for _, item := range tsvArr {
		sessionFileName := item.SessionName + ".tsv"
		itemPath := filepath.Join(exportPath, sessionFileName)
		err := copy(item.TsvPath, itemPath)
		if err != nil {
			fmt.Printf("failed to copy: %s tsv path: %s", err, item.TsvPath)
			panic(err)
		}
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

func processFiles(dir string) []TsvInfo {
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
			tsvPath, _ := tsvCheck(filePath)
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

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	firstName := os.Args[1]
	lastName := os.Args[2]
	dirExists, _ := dirCheck(wd)
	if dirExists {
		tsvInfos := processFiles(wd)
		createExport(tsvInfos, wd, firstName, lastName)
	} else {
		fmt.Printf("could not find directory %s", wd)
	}
}
