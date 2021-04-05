package picturecross

import "strings"

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
	cells [][]Cell
}

func NewBoardFromSettings(rowSettings, colSettings []Settings) *Board {
	// TODO: Invalid board detection.
	cells := make([][]Cell, len(rowSettings))
	for i := range cells {
		cells[i] = make([]Cell, len(colSettings))
	}

	return &Board{
		cells: cells,
	}
}

func (b Board) String() string {
	s := strings.Builder{}

	if len(b.cells) == 0 {
		// Draw empty board.
		s.WriteString(strings.Repeat("─", 2))
		s.WriteString("\n")
		s.WriteString(strings.Repeat("│", 2))
		s.WriteString("\n")
		s.WriteString(strings.Repeat("─", 2))
		s.WriteString("\n")
	} else {
		s.WriteString(strings.Repeat("─", len(b.cells[0])+2))
		s.WriteString("\n")

		for _, r := range b.cells {
			s.WriteString("│")
			for _, c := range r {
				s.WriteString(c.String())
			}

			s.WriteString("│")
			s.WriteString("\n")
		}

		// Overline
		s.WriteString(strings.Repeat("─", len(b.cells[0])+2))
		s.WriteString("\n")
	}

	return s.String()
}
