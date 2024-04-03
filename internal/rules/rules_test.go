package rules

import (
	. "example/tictactoe/internal/types"
	"testing"
)

var win = [3]string{"X", "X", "X"}
var empty = [3]string{"", "", ""}
var first = [3]string{"X", "", ""}

func TestDetectWinFirstRow(t *testing.T) {
	var player Player

	var board = [3][3]string{win, empty, empty}

	DetectWin(board, &player)

	if player != X {
		t.Fatal("X is not the winner")
	}
}

func TestDetectWinFirstColumn(t *testing.T) {
	var player Player

	var board = [3][3]string{first, first, first}

	DetectWin(board, &player)

	if player != X {
		t.Fatal("X is not the winner")
	}
}

func TestDetectWinLeftDiagonal(t *testing.T) {
	var player Player

	var board = [3][3]string{first, {"", "X", ""}, {"", "", "X"}}

	DetectWin(board, &player)

	if player != X {
		t.Fatal("X is not the winner")
	}
}

func TestDetectWinRightDiagonal(t *testing.T) {
	var player Player

	var board = [3][3]string{{"", "", "X"}, {"", "X", ""}, first}

	DetectWin(board, &player)

	if player != X {
		t.Fatal("X is not the winner")
	}
}
