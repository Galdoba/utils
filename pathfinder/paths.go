package pathfinder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var sep = string(filepath.Separator)

func asDir(path string) string {
	return strings.TrimSuffix(path, sep) + sep
}

// HomeDir - return user home dir.
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("no home found: %v", err))
	}
	return asDir(home)
}

// WorkDir - return current working dir.
func WorkDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("no work directory found: %v", err))
	}
	return asDir(wd)
}

func ConfigDir() string {
	return asDir(HomeDir() + ".config")
}
func LogDir() string {
	return asDir(HomeDir() + ".log")
}
func ProgramsDir() string {
	return asDir(HomeDir() + "Programs")
}
