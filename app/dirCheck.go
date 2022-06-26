package app

import "os"

//Confirm directory is real, will return a boolean error if untrue.
func DirCheck(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
