package app

import (
	"io/ioutil"
	"os"
)

// reads, creates and writes curated_speech_only.tsv's
// from the current batch and drops them into the generated folder.
// along with its new name.
func Copy(src, dst string) error {
	// Read the File
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	bs, err := ioutil.ReadAll(source)
	if err != nil {
		return err
	}
	//Creates the new File destination
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	//writes from memory to the destination
	_, err2 := destination.Write(bs)
	return err2
}
