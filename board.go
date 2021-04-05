package picturecross

import (
	"strconv"
	"strings"
)

type Value int

const (
	ValueUnset Value = iota
	ValueCross
	ValueBlock
)

func (v Value) String() string {
	switch v {
	case ValueCross:
		return "X"
	case ValueBlock:
		return "█"
	}

	return " "
}

type Cell struct {
	value   Value
	certain bool // If true then the value above is definitely correct
}

func (c Cell) String() string {
	if !c.certain {
		return " "
	}

	return c.value.String()
}

type Settings []int

type Board struct {
	cells       [][]Cell
	rowSettings []Settings
	colSettings []Settings
}

func NewBoardFromSettings(rowSettings, colSettings []Settings) *Board {
	// TODO: Invalid board detection.
	cells := make([][]Cell, len(rowSettings))
	for i := range cells {
		cells[i] = make([]Cell, len(colSettings))
	}

	return &Board{
		cells:       cells,
		rowSettings: rowSettings,
		colSettings: colSettings,
	}
}

func (b Board) String() string {
	if len(b.cells) == 0 {
		return "<empty>"
	}

	maxRowSettings := maxSettingsLength(b.rowSettings)
	maxColSettings := maxSettingsLength(b.colSettings)

	s := strings.Builder{}

	// ColSettings at the top
	for i := 0; i < maxColSettings; i++ {
		// Padding for row settings
		s.WriteString(strings.Repeat(" ", maxRowSettings))
		s.WriteString("│")
		for _, ss := range b.colSettings {
			idx := i + len(ss) - maxColSettings
			if idx < 0 {
				s.WriteString(" ")
			} else {
				s.WriteString(strconv.Itoa(ss[idx]))
			}
		}
		s.WriteString("\n")
	}

	s.WriteString(strings.Repeat("─", len(b.cells[0])+1+maxRowSettings))
	s.WriteString("\n")

	for rowIdx, r := range b.cells {
		if len(b.rowSettings[rowIdx]) < maxRowSettings {
			s.WriteString(strings.Repeat(" ", maxRowSettings-len(b.rowSettings[rowIdx])))
		}

		for _, sv := range b.rowSettings[rowIdx] {
			s.WriteString(strconv.Itoa(sv))
		}

		s.WriteString("│")
		for _, c := range r {
			s.WriteString(c.String())
		}

		s.WriteString("\n")
	}

	return s.String()
}

func maxSettingsLength(s []Settings) int {
	maxLength := 0
	for _, ss := range s {
		if len(ss) > maxLength {
			maxLength = len(ss)
		}
	}

	return maxLength
}
