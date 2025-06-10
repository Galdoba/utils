package pathfinder

import (
	"fmt"
)

const (
	ROOT   = "root"
	HOME   = "home"
	CONFIG = "config"
	LOG    = "log"
	DATA   = "data"
	WORK   = "work"
)

type pathPlot struct {
	root        string
	upperLayers []string
	target      string
	lowerLayers []string
	file        string
}

func NewPlot(options ...PlotOption) (pathPlot, error) {
	pp := pathPlot{
		root:        HOME,
		upperLayers: []string{},
		target:      "",
		lowerLayers: []string{},
	}
	for _, modify := range options {
		modify(&pp)
	}
	return pp, assertPlot(pp)
}

func assertPlot(pp pathPlot) error {
	if (len(pp.upperLayers)+len(pp.lowerLayers) > 0) && pp.target == "" {
		return fmt.Errorf("no layers definded")
	}
	return nil
}

type PlotOption func(*pathPlot)

func WithRoot(r string) PlotOption {
	return func(pp *pathPlot) {
		switch r {
		case ROOT, HOME, CONFIG, LOG, DATA, WORK:
			pp.root = r
		}
	}
}

func WithTarget(tg string) PlotOption {
	return func(pp *pathPlot) {
		pp.target = tg
	}
}

func WithUpperLayers(layers ...string) PlotOption {
	return func(pp *pathPlot) {
		pp.upperLayers = layers
	}
}

func WithLowerLayers(layers ...string) PlotOption {
	return func(pp *pathPlot) {
		pp.lowerLayers = layers
	}
}

func WithFile(file string) PlotOption {
	return func(pp *pathPlot) {
		pp.file = file
	}
}

// Project - return absolute path of the plot or string of plot's error.
func Project(pp pathPlot) string {
	if err := assertPlot(pp); err != nil {
		return fmt.Sprintf("[bad plot: %v]", err)
	}
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
	root = asDir(root)
	if pp.file != "" {
		root += pp.file
	}

	return root
}
