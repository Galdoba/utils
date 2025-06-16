package text

import (
	"strings"

	"github.com/Galdoba/utils/slicetricks"
)

type wrapper struct {
	lineMaxWidth int
	linesAllowed int
	leftOffset   int
	rightOffset  int
	wrapLimit    int
	wrapBefore   []string
	wrapAfter    []string
}

func Wrap(text string, options ...WrappingOption) string {
	wr := wrapper{
		lineMaxWidth: -1,
		linesAllowed: -1,
		leftOffset:   0,
		rightOffset:  0,
		wrapLimit:    0,
		wrapBefore:   []string{"(", "[", "{", "`", `"`, "<"},
		wrapAfter:    []string{" ", "-", "\t"},
	}
	for _, set := range options {
		set(&wr)
	}
	schema := schema(wr, text)
	if wr.lineMaxWidth < 1 {
		return text
	}
	lines := composeLines(wr, schema)

	return strings.Join(lines, "\n")
}

type WrappingOption func(*wrapper)

func LineMaxWidth(mw int) WrappingOption {
	return func(w *wrapper) {
		w.lineMaxWidth = mw
	}
}

func LinesAllowed(mw int) WrappingOption {
	return func(w *wrapper) {
		w.linesAllowed = mw
	}
}

func LeftOffset(lo int) WrappingOption {
	return func(w *wrapper) {
		w.leftOffset = lo
	}
}

func WrapAllowed(wl int) WrappingOption {
	return func(w *wrapper) {
		w.wrapLimit = wl
	}
}

type wrappingSchema map[int]litera

type litera struct {
	text       string
	rn         rune
	maybeStart bool
	maybeEnd   bool
}

var lit_Space = litera{
	text:       " ",
	rn:         32,
	maybeStart: true,
	maybeEnd:   true,
}

var lit_NewLine = litera{
	text:       "\n",
	rn:         10,
	maybeStart: false,
	maybeEnd:   true,
}

func schema(wr wrapper, text string) wrappingSchema {
	ws := make(wrappingSchema)
	lits := strings.Split(text, "")
	for i, l := range lits {
		lit := litera{}
		lit.text = l
		lit.maybeStart = slicetricks.Contains(wr.wrapBefore, l)
		lit.maybeStart = slicetricks.Contains(wr.wrapBefore, l)
		lit.maybeEnd = slicetricks.Contains(wr.wrapAfter, l)
		for _, r := range l {
			lit.rn = r
		}
		ws[i] = lit
	}
	return ws
}

func composeLines(wr wrapper, sc wrappingSchema) []string {
	lineMaxWidth := wr.lineMaxWidth
	lines := []string{}
	candidate := newCandidate(0)
	trimmed := -1
	lenText := len(sc)
	for i := 0; i < lenText; i++ {

		lit := sc[i]
		if lit.text == "\n" {
			lines = append(lines, toLine(candidate))
			candidate = newCandidate(wr.leftOffset)
			continue
		}
		candidate = append(candidate, lit)

		if len(candidate) == lineMaxWidth {
			candidate, trimmed = candidate.trimSuffix(wr.wrapLimit)
			switch trimmed {
			case 0:
				lines = append(lines, toLine(candidate))
				candidate = newCandidate(wr.leftOffset)
			default:
				lines = append(lines, toLine(candidate))
				candidate = newCandidate(wr.leftOffset)
				i = i - trimmed
			}

		}

	}
	lines = append(lines, toLine(candidate))
	lines = joinAfter(lines, wr.linesAllowed)
	return lines
}

func joinAfter(slice []string, n int) []string {
	switch n {
	case 1:
		return []string{strings.Join(slice, "")}
	default:
		if n < 1 {
			return slice
		}
	}

	joined := []string{}
	for i, line := range slice {
		if i == 0 {
			joined = append(joined, line)
			continue
		}
		switch i >= n {
		case false:
			joined = append(joined, line)
		case true:
			joined[len(joined)-1] = joined[len(joined)-1] + line
		}
	}
	return joined
}

type candidate []litera

func newCandidate(offset int) candidate {
	cnd := candidate{}
	for i := 0; i < offset; i++ {
		cnd = append(cnd, lit_Space)
	}
	return cnd
}

func (cnd candidate) trimSuffix(trimax int) (candidate, int) {
	trimmed := 0
	origin := cnd
	for !cnd[len(cnd)-1].maybeEnd {
		if trimmed >= trimax {
			return origin, 0
		}
		cnd = cnd[:len(cnd)-1]
		trimmed++
		if len(cnd) == 0 {
			return origin, 0
		}
		if cnd[len(cnd)-1].maybeStart {
			cnd = cnd[:len(cnd)-1]
			trimmed++
			return cnd, trimmed
		}

	}
	return cnd, trimmed
}

func toLine(cnd candidate) string {
	line := ""
	for _, lit := range cnd {
		line += lit.text
	}
	return line
}
