package util

import (
	"os"
	"path/filepath"
)

func GetExecutablePath() string {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(executable)
}
