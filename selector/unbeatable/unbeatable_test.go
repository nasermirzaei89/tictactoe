package unbeatable_test

import (
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector/unbeatable"
	"testing"
)

func TestUnbeatableVsUnbeatable(t *testing.T) {
	ai1 := unbeatable.New()
	ai2 := unbeatable.New()

	game := tictactoe.New()

	for {
		switch game.Turn() {
		case 1:
			input := ai1.ChooseBox(*game)
			game.Play(input)
		case -1:
			input := ai2.ChooseBox(*game)
			game.Play(input)
		}

		if ok, who := game.GameOver(); ok {
			if who != 0 {
				t.Error("an unbeatable player lost")
			} else {
				break
			}
		}
	}
}
