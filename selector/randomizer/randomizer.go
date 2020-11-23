package randomizer

import (
	"fmt"
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector"
	"math/rand"
	"time"
)

type random struct {
	rnd *rand.Rand
}

func New() selector.AI {
	return &random{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ai random) ChooseBox(game tictactoe.Game) string {
	moves := game.PossibleMoves()

	if len(moves) == 0 {
		return "e"
	}

	return moves[ai.rnd.Intn(len(moves))]
}

func (ai random) Title(n int) string {
	return fmt.Sprintf("Random %d", n)
}

func (ai random) Prompt() string {
	return "Randomising..."
}
