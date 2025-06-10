package pathfinder

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Project - return absolute path of the plot or string of plot's error.
func Project(pp pathPlot) string {
	if err := assertPlot(pp); err != nil {
		return fmt.Sprintf("[bad plot: %v]", err)
	}
	path := projectDirectory(pp)
	if pp.file != "" {
		path += pp.file
	}
	return path
}

func projectDirectory(pp pathPlot) string {
	root := ""
	switch pp.root {
	case ROOT:
		root = sep
	case HOME:
		root = HomeDir()
	case LOG:
		root = LogDir()
	case CONFIG:
		root = ConfigDir()
	case DATA:
		root = ProgramsDir()
	case WORK:
		root = WorkDir()
	}
	for _, layer := range pp.upperLayers {
		if layer == "" {
			continue
		}
		root += layer + sep
	}
	if pp.target != "" {
		root += pp.target + sep
	}
	for _, layer := range pp.lowerLayers {
		if layer == "" {
			continue
		}
		root += layer + sep
	}
	return asDir(root)
}

// Status - return status message of plots projection.
func Status(pp pathPlot) string {
	if err := assertPlot(pp); err != nil {
		return fmt.Sprintf("path status: bad plot")
	}
	path := Project(pp)
	layers := strings.Split(path, sep)
	last := len(layers) - 1
	reconstructed := ""
	for lNumper, layer := range layers {
		switch lNumper {
		default:
			reconstructed += layer + sep
			st, err := os.Stat(reconstructed)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return fmt.Sprintf("path status: upper layer directory does not exist (%v)", reconstructed)
				}
				return fmt.Sprintf("path status: no stats for layer %v: %v", reconstructed, err)
			}
			if !st.IsDir() {
				return fmt.Sprintf("path status: upper layer is not a directory (%v)", reconstructed)
			}
		case last:
			switch pp.file {
			case "":
				reconstructed += layer + sep
				st, err := os.Stat(reconstructed)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) {
						return fmt.Sprintf("path status: directory does not exist")
					}
					return fmt.Sprintf("path status: no stats directory: %v", err)
				}
				if !st.IsDir() {
					return fmt.Sprintf("path status: plot is not a directory")
				}
			default:
				reconstructed += pp.file
				st, err := os.Stat(reconstructed)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) {
						return fmt.Sprintf("path status: plot file does not exist")
					}
					return fmt.Sprintf("path status: no stats plot file: %v", err)
				}
				if st.IsDir() {
					return fmt.Sprintf("path status: plot is not a file")
				}
			}
		}
	}
	st, _ := os.Stat(path)
	perm := st.Mode().Perm().String()
	return fmt.Sprintf("path status: ok (permission: %v)", perm)
}

// Pave - create plot path. Return error if failed.
func Pave(pp pathPlot) error {
	if err := assertPlot(pp); err != nil {
		return fmt.Errorf("failed to assert path: %v", err)
	}
	path := Project(pp)
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if err := os.MkdirAll(projectDirectory(pp), 0777); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	if pp.file == "" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("failed to read/create plotted file")
	}
	return f.Close()
}
