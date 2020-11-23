package selector_test

import (
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector/randomizer"
	"github.com/nasermirzaei89/tictactoe/selector/unbeatable"
	"testing"
)

func TestUnbeatableVsRandom(t *testing.T) {
	ai1 := unbeatable.New()
	ai2 := randomizer.New()

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
			if who == -1 {
				t.Error("unbeatable player lost")
			} else {
				break
			}
		}
	}
}

func TestRandomVsUnbeatable(t *testing.T) {
	ai1 := randomizer.New()
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
			if who == 1 {
				t.Error("unbeatable player lost")
			} else {
				break
			}
		}
	}
}
