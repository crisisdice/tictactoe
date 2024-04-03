package opponent

import (
	. "example/tictactoe/internal/rendering"
	. "example/tictactoe/internal/types"
	"testing"
)

var win = [3]string{"X", "X", "X"}
var empty = [3]string{"", "", ""}
var first = [3]string{"X", "", ""}

func TestComputerMoveFirstMove(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	ComputerMove(&board, 0, player)

	if board[0][0] != "X" {
		t.Fatal("Corner Not X")
	}
}

func TestComputerMoveSecondMove01(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[0][1] = "O"

	ComputerMove(&board, 1, player)

	if board[2][0] != "X" {
		t.Fatal("Corner Not X")
	}
}

func TestComputerMoveThirdMove01Blocked(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[0][1] = "O"
	board[2][0] = "X"
	board[1][0] = "O"

	ComputerMove(&board, 2, player)

	if board[2][2] != "X" {
		t.Fatal("Corner Not X")
	}
}

func TestComputerMoveThirdMove01UnBlocked(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[0][1] = "O"
	board[2][0] = "X"
	board[2][2] = "O"

	ComputerMove(&board, 1, player)

	if board[1][0] != "X" {
		t.Fatal("Corner Not X")
	}
}

func TestComputerMoveSecondMove11(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[1][1] = "O"

	ComputerMove(&board, 1, player)

	if board[2][2] != "X" {
		t.Fatal("Corner Not X")
	}
}

func TestComputerMoveThirdMove11_02(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	board[0][2] = "O"

	ComputerMove(&board, 2, player)

	if board[2][0] != "X" {
		t.Fatal("Corner Not X")
	}
}

func TestComputerMoveForthMove11_02Unblocked(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	board[0][2] = "O"
	board[2][0] = "X"
	board[0][1] = "O"

	ComputerMove(&board, 3, player)

	if board[2][1] != "X" {
		t.Fatal("Win not taken")
	}
}

func TestComputerMoveForthMove11_02Blocked10(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	board[0][2] = "O"
	board[2][0] = "X"
	board[1][0] = "O"

	ComputerMove(&board, 3, player)

	if board[2][1] != "X" {
		t.Fatal("Win not taken")
	}
}

func TestComputerMoveForthMove11_02Blocked21(t *testing.T) {
	var player Player = X
	var board = InitBoard()

	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	board[0][2] = "O"
	board[2][0] = "X"
	board[2][1] = "O"

	ComputerMove(&board, 3, player)

	if board[1][0] != "X" {
		t.Fatal("Win not taken")
	}
}
