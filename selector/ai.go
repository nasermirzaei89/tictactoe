package selector

import "github.com/nasermirzaei89/tictactoe"

type AI interface {
	ChooseBox(game tictactoe.Game) string
	Title(n int) string
	Prompt() string
}
