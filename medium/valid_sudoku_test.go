package medium

import "testing"

func TestValidSudoku(t *testing.T) {
	tests := []struct {
		board [][]byte
		want  bool
	}{
		{
			board: [][]byte{
				[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: true,
		},
		{
			board: [][]byte{
				[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: false,
		},
		{
			board: [][]byte{
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			want: true,
		},
	}

	for i, tt := range tests {
		got := isValidSudoku(tt.board)
		if got != tt.want {
			t.Fatalf("%d. got: %v; want: %v", i, got, tt.want)
		}
	}
}
