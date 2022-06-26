package main

import (
	"fmt"
	"os"
	"speechOnlyToaster/app"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	firstName := os.Args[1]
	lastName := os.Args[2]
	dirExists, _ := app.DirCheck(wd)
	if dirExists {
		tsvInfos := app.ProcessFiles(wd)
		app.CreateExport(tsvInfos, wd, firstName, lastName)
	} else {
		fmt.Printf("could not find directory %s", wd)
	}
}
