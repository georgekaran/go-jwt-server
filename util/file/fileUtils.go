package file

import (
	"github.com/georgekaran/go-jwt-server/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Read a .properties file and create a map.
func ToMap(filename string) map[string]string {
	fileMap := make(map[string]string)

	pathBase := filepath.Join(ConfigDir(), filename)
	bs, err := ioutil.ReadFile(pathBase)
	util.Must(err)

	lines := strings.Split(string(bs), "\n")
	for _, line := range lines {
		if strings.Contains(line, "#") || line == "" {
			continue
		}
		lineSplit := strings.Split(line, "=")
		fileMap[lineSplit[0]] = lineSplit[1]
	}
	return fileMap
}

// Return the root directory of the project.
func RootDir() string {
	dir, err := os.Getwd()
	util.CheckPrint(err)
	return dir
}

// Return the config directory of the project.
func ConfigDir() string {
	return filepath.Join(RootDir(), "config")
}