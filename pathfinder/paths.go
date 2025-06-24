package pathfinder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var sep = string(filepath.Separator)

func asDir(path string) string {
	if path == "" {
		return sep
	}
	if strings.HasSuffix(path, sep) {
		return path
	}
	return path + sep
}

var (
	homeDirCache     string
	homeOnce         sync.Once
	configDirCache   string
	configOnce       sync.Once
	logDirCache      string
	logOnce          sync.Once
	programsDirCache string
	programsOnce     sync.Once
)

// HomeDir - return user home dir.
func HomeDir() string {
	homeOnce.Do(func() {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("no home found: %v", err))
		}
		homeDirCache = asDir(home)
	})
	return homeDirCache
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
	configOnce.Do(func() {
		configDirCache = asDir(HomeDir()) + ".config"
	})
	return configDirCache
}
func LogDir() string {
	logOnce.Do(func() {
		logDirCache = asDir(HomeDir()) + ".log"
	})
	return logDirCache
}
func ProgramsDir() string {
	programsOnce.Do(func() {
		programsDirCache = asDir(HomeDir()) + "Programs"
	})
	return configDirCache
}
