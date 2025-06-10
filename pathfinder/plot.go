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
