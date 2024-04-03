package rules

import (
	. "example/tictactoe/internal/types"
)

func IsLegalMove(board *Board, cursor *Cursor) bool {
	var (
		row         = cursor[1]
		collumn     = cursor[0]
		isLegalMove = board[row][collumn] == ""
	)

	return isLegalMove
}

func AllSame(arr [3]string) bool {
	return arr[0] == arr[1] && arr[0] == arr[2]
}

func detectWinner(char string) Player {
	if char == "X" {
		return X
	}

	return O
}

func SplitBoard(board *Board) (rows [3][3]string, collumns [3][3]string, leftDiagonal [3]string, rightDiagonal [3]string) {
	rows = *board

	collumns = [3][3]string{
		{board[0][0], board[1][0], board[2][0]},
		{board[0][1], board[1][1], board[2][1]},
		{board[0][2], board[1][2], board[2][2]},
	}

	leftDiagonal = [3]string{board[0][0], board[1][1], board[2][2]}
	rightDiagonal = [3]string{board[0][2], board[1][1], board[2][0]}

	return
}

func DetectWin(board Board, winner *Player) {
	var rows, collumns, leftDiagonal, rightDiagonal = SplitBoard(&board)

	for _, row := range rows {
		if AllSame(row) && row[0] != "" {
			*winner = detectWinner(row[0])
			return
		}
	}

	for _, collumn := range collumns {
		if AllSame(collumn) && collumn[0] != "" {
			*winner = detectWinner(collumn[0])
			return
		}
	}

	if AllSame(leftDiagonal) && leftDiagonal[0] != "" {
		*winner = detectWinner(leftDiagonal[0])
		return

	}

	if AllSame(rightDiagonal) && rightDiagonal[0] != "" {
		*winner = detectWinner(rightDiagonal[0])
		return
	}
}

func DetectDraw(board Board) bool {
	for _, row := range board {
		for _, collumn := range row {
			if collumn == "" {
				return false
			}
		}
	}

	return true
}

func GetOtherPlayer(activePlayer Player) Player {
	if activePlayer == X {
		return O
	} else {
		return X
	}
}

func SwitchActivePlayer(activePlayer *Player) {
	*activePlayer = GetOtherPlayer(*activePlayer)
}
