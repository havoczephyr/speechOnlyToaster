package app

import (
	"io/ioutil"
	"path/filepath"
)

//filters if the selected "session-" contains "curated_processed_speech_only.tsv
func TsvCheck(dir string) (string, error) {
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
