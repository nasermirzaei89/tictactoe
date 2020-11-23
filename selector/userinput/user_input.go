package userinput

import (
	"fmt"
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector"
	"io"
)

type userInput struct {
	r io.Reader
}

func New(r io.Reader) selector.AI {
	return &userInput{
		r: r,
	}
}

func (ai userInput) ChooseBox(_ tictactoe.Game) string {
	var in string
	_, _ = fmt.Fscan(ai.r, &in)
	return in
}

func (ai userInput) Title(n int) string {
	return fmt.Sprintf("User %d", n)
}

func (ai userInput) Prompt() string {
	return "Enter a box number:"
}
