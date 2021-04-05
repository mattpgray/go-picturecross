package picturecross

import "testing"

func TestBoardString(t *testing.T) {

	tests := []struct {
		board *Board
		st    string
	}{
		{
			board: &Board{},
			st: "──\n" +
				"││\n" +
				"──\n",
		},
		{
			board: NewBoardFromSettings([]Settings{{1}, {1}}, []Settings{{1}, {1}, {1}}),
			st: "─────\n" +
				"│   │\n" +
				"│   │\n" +
				"─────\n",
		},
	}

	for i, test := range tests {
		str := test.board.String()
		if str != test.st {
			t.Errorf("test case %d did not give the expected result. Expected:\n%qfound:\n%q", i, test.st, str)
		}
	}
}
