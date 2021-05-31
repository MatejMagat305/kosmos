package help

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FindPath(configs, config string) string {
	if FileExist(configs) {
		return "./" + configs + "/" + getConfigOs(config)
	}
	return fmt.Sprint(RekSearch("./", configs), getConfigOs(config))
}

func RekSearch(path string, configs string) string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return ""
	}
	for _, f := range files {
		if f.IsDir() {
			if f.Name() == configs {
				return fmt.Sprint(path, f.Name(), "/")
			}
			r := RekSearch(fmt.Sprint(path, f.Name(), "/"), configs)
			if r != "" {
				return r
			}
		}
	}
	return ""
}

func getConfigOs(config string) string {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		return config
	}
	return argsWithoutProg[0] + ".txt"
}