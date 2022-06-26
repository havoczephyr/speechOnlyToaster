package app

import "os"

func DirCheck(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
