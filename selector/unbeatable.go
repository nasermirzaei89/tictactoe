package selector

import (
	"fmt"
	"github.com/nasermirzaei89/tictactoe"
	"math/rand"
	"time"
)

type unbeatable struct {
	rnd *rand.Rand
}

func NewUnbeatable() AI {
	return &unbeatable{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ai unbeatable) ChooseBox(game tictactoe.Game) (input string) {
	if game.Turn() == 1 {
		input, _ = ai.max(game)
	} else {
		input, _ = ai.min(game)
	}

	return
}

func (ai unbeatable) max(game tictactoe.Game) (string, int) {
	if ok, v := game.GameOver(); ok {
		return "e", v
	}

	moves := game.PossibleMoves()
	ai.rnd.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })
	bestS := "e"
	bestV := -2
	for _, move := range moves {
		game2 := game
		(&game2).Play(move)
		_, v := ai.min(game2)
		if v > bestV {
			bestV = v
			bestS = move
		}
	}

	return bestS, bestV
}

func (ai unbeatable) min(game tictactoe.Game) (string, int) {
	if ok, v := game.GameOver(); ok {
		return "e", v
	}

	moves := game.PossibleMoves()
	ai.rnd.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })
	bestS := "e"
	bestV := 2
	for _, move := range moves {
		game2 := game
		(&game2).Play(move)
		_, v := ai.max(game2)
		if v < bestV {
			bestV = v
			bestS = move
		}
	}

	return bestS, bestV
}

func (ai unbeatable) Title(n int) string {
	return fmt.Sprintf("Unbeatable %d", n)
}

func (ai unbeatable) Prompt() string {
	return "Thinking..."
}
