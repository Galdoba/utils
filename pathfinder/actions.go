package pathfinder

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	var sb strings.Builder
	switch pp.root {
	case ROOT:
		sb.WriteString(sep)
	case HOME:
		sb.WriteString(HomeDir())
	case LOG:
		sb.WriteString(LogDir())
	case CONFIG:
		sb.WriteString(ConfigDir())
	case DATA:
		sb.WriteString(ProgramsDir())
	case WORK:
		sb.WriteString(WorkDir())
	}
	for _, layer := range pp.upperLayers {
		if layer == "" {
			continue
		}
		sb.WriteString(layer)
		sb.WriteString(sep)
	}
	if pp.target != "" {
		sb.WriteString(pp.target)
		sb.WriteString(sep)
	}
	for _, layer := range pp.lowerLayers {
		if layer == "" {
			continue
		}
		sb.WriteString(layer)
		sb.WriteString(sep)
	}
	return sb.String()
}

// Status - return status message of plots projection.
func Status(pp pathPlot) string {
	if err := assertPlot(pp); err != nil {
		return fmt.Sprintf("path status: bad plot")
	}
	path := Project(pp)
	layers := strings.Split(path, sep)
	last := len(layers) - 1
	// reconstructed := ""
	var sb strings.Builder
	for lNumper, layer := range layers {
		switch lNumper {
		default:
			// reconstructed += layer + sep
			sb.WriteString(layer)
			sb.WriteString(sep)
			// st, err := os.Stat(reconstructed)
			st, err := os.Stat(sb.String())
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					return fmt.Sprintf("path status: upper layer directory does not exist (%v)", sb.String())
				}
				return fmt.Sprintf("path status: no stats for layer %v: %v", sb.String(), err)
			}
			if !st.IsDir() {
				return fmt.Sprintf("path status: upper layer is not a directory (%v)", sb.String())
			}
		case last:
			switch pp.file {
			case "":
				// reconstructed += layer + sep
				sb.WriteString(layer)
				sb.WriteString(sep)
				st, err := os.Stat(sb.String())
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
				// 				sb.String()+= pp.file
				sb.WriteString(pp.file)
				st, err := os.Stat(sb.String())
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
	dir := projectDirectory(pp)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	if pp.file == "" {
		return nil
	}
	filepath := filepath.Join(dir, pp.file)
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("failed to read/create plotted file")
	}
	return f.Close()
}
