package selector

import (
	"fmt"
	"github.com/nasermirzaei89/tictactoe"
	"io"
)

type input struct {
	r io.Reader
}

func NewInput(r io.Reader) AI {
	return &input{
		r: r,
	}
}

func (ai input) ChooseBox(game tictactoe.Game) string {
	var in string
	_, _ = fmt.Fscan(ai.r, &in)
	return in
}

func (ai input) Title(n int) string {
	return fmt.Sprintf("Player %d", n)
}

func (ai input) Prompt() string {
	return "Enter a box number:"
}
