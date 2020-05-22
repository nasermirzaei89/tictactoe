package tictactoe

import (
	"strconv"
)

type Game struct {
	board [3][3]int
	turn  int
}

func New() *Game {
	game := Game{
		board: [3][3]int{},
		turn:  1,
	}

	return &game
}

func (ttt Game) Turn() int {
	return ttt.turn
}

func (ttt Game) Board() [3][3]int {
	return ttt.board
}

func (ttt Game) IsValidMove(input string) bool {
	for _, v := range ttt.PossibleMoves() {
		if input == v {
			return true
		}
	}

	return false
}

func (ttt *Game) Play(input string) {
	i, j := ttt.getIndexes(input)
	ttt.board[i][j] = ttt.Turn()
	ttt.turn = -ttt.turn
}

func (ttt Game) getIndexes(s string) (int, int) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return -1, -1
	}

	i := (v - 1) / 3
	j := (v - 1) % 3
	return i, j
}

func (ttt Game) PossibleMoves() []string {
	var moves []string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if ttt.board[i][j] == 0 {
				moves = append(moves, strconv.Itoa(i*3+j+1))
			}
		}
	}

	return moves
}

func (ttt Game) GameOver() (bool, int) {
	turn := -ttt.Turn()
	win :=
		(ttt.board[0][0] == turn && ttt.board[0][1] == turn && ttt.board[0][2] == turn) ||
			(ttt.board[1][0] == turn && ttt.board[1][1] == turn && ttt.board[1][2] == turn) ||
			(ttt.board[2][0] == turn && ttt.board[2][1] == turn && ttt.board[2][2] == turn) ||
			(ttt.board[0][0] == turn && ttt.board[1][0] == turn && ttt.board[2][0] == turn) ||
			(ttt.board[0][1] == turn && ttt.board[1][1] == turn && ttt.board[2][1] == turn) ||
			(ttt.board[0][2] == turn && ttt.board[1][2] == turn && ttt.board[2][2] == turn) ||
			(ttt.board[0][0] == turn && ttt.board[1][1] == turn && ttt.board[2][2] == turn) ||
			(ttt.board[0][2] == turn && ttt.board[1][1] == turn && ttt.board[2][0] == turn)

	if win {
		return true, turn
	}

	return len(ttt.PossibleMoves()) == 0, 0
}
