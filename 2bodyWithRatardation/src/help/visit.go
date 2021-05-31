package help

import (
	"log"
	"os"
	"path/filepath"
)

func Visit(dirs *[]string, condition func(os.FileInfo) bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if condition(info) {
			*dirs = append(*dirs, info.Name())
		}
		return nil
	}
}

func FileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}