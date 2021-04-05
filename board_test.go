package picturecross

import "testing"

func TestBoardString(t *testing.T) {

	tests := []struct {
		board *Board
		st    string
	}{
		{
			board: &Board{},
			st:    "<empty>",
		},
		{
			board: NewBoardFromSettings([]Settings{{1}, {1}}, []Settings{{1}, {1}, {1}}),
			st: " │111\n" +
				"─────\n" +
				"1│   \n" +
				"1│   \n",
		},
		{
			board: NewBoardFromSettings([]Settings{{1, 1}, {1}}, []Settings{{1, 1}, {1}, {1}}),
			st: "  │1  \n" +
				"  │111\n" +
				"─────\n" +
				"11│   \n" +
				" 1│   \n",
		},
	}

	for i, test := range tests {
		str := test.board.String()
		if str != test.st {
			t.Errorf("test case %d did not give the expected result. Expected:\n%sfound:\n%s", i, test.st, str)
		}
	}
}
