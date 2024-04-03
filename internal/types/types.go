package types

type (
	Board  = [3][3]string
	Player int
	Cursor = [2]int
)

const (
	X = iota + 1
	O
)
