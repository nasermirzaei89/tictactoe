package selector_test

import (
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector"
	"testing"
)

func TestUnbeatableVsUnbeatable(t *testing.T) {
	ai1 := selector.NewUnbeatable()
	ai2 := selector.NewUnbeatable()

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

func TestUnbeatableVsRandom(t *testing.T) {
	ai1 := selector.NewUnbeatable()
	ai2 := selector.NewRandom()

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
	ai1 := selector.NewRandom()
	ai2 := selector.NewUnbeatable()

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
