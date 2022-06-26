package app

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateExport(tsvArr []TsvInfo, rootdir string, firstName string, lastName string) {
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
		err := Copy(item.TsvPath, itemPath)
		if err != nil {
			fmt.Printf("failed to copy: %s tsv path: %s", err, item.TsvPath)
			panic(err)
		}
	}

}
