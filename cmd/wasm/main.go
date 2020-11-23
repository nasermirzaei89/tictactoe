// +build js

package main

import (
	"fmt"
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector"
	"syscall/js"
)

var (
	game          *tictactoe.Game
	ui            *userInput
	player1       selector.AI
	player2       selector.AI
	player1Select = make(chan selector.AI)
	player2Select = make(chan selector.AI)
)

func getPlayerSign(turn int) string {
	if turn == 1 {
		return "×"
	}

	if turn == -1 {
		return "○"
	}

	return ""
}

func setMessage(text string) {
	js.Global().Get("document").Call("getElementById", "message").Set("innerText", text)
}

func clear() {
	doc := js.Global().Get("document")
	container := doc.Call("getElementById", "container")
	container.Set("innerHTML", "")
	p := doc.Call("createElement", "p")
	p.Call("setAttribute", "id", "message")
	container.Call("appendChild", p)
}

func main() {
	game = tictactoe.New()

	ui = new(userInput)
	*ui = make(chan []byte)

	doc := js.Global().Get("document")

	container := doc.Call("getElementById", "container")

	// select player 1
	clear()
	doc.Call("getElementById", "message").Set("innerText", "Select Player 1")

	btn := doc.Call("createElement", "BUTTON")
	btn.Set("innerText", "User Input")
	btn.Call("setAttribute", "id", "userInput")
	btn.Call("addEventListener", "click", js.FuncOf(setPlayer1))
	btn.Get("classList").Call("add", "btn")
	container.Call("appendChild", btn)

	btn = doc.Call("createElement", "BUTTON")
	btn.Set("innerText", "Unbeatable AI")
	btn.Call("setAttribute", "id", "unbeatableAI")
	btn.Get("classList").Call("add", "btn")
	btn.Call("addEventListener", "click", js.FuncOf(setPlayer1))
	container.Call("appendChild", btn)

	btn = doc.Call("createElement", "BUTTON")
	btn.Set("innerText", "Random AI")
	btn.Call("setAttribute", "id", "randomAI")
	btn.Get("classList").Call("add", "btn")
	btn.Call("addEventListener", "click", js.FuncOf(setPlayer1))
	container.Call("appendChild", btn)

	player1 = <-player1Select

	// select player 2
	clear()
	doc.Call("getElementById", "message").Set("innerText", "Select Player 2")

	btn = doc.Call("createElement", "BUTTON")
	btn.Set("innerText", "User Input")
	btn.Call("setAttribute", "id", "userInput")
	btn.Get("classList").Call("add", "btn")
	btn.Call("addEventListener", "click", js.FuncOf(setPlayer2))
	container.Call("appendChild", btn)

	btn = doc.Call("createElement", "BUTTON")
	btn.Set("innerText", "Unbeatable AI")
	btn.Call("setAttribute", "id", "unbeatableAI")
	btn.Get("classList").Call("add", "btn")
	btn.Call("addEventListener", "click", js.FuncOf(setPlayer2))
	container.Call("appendChild", btn)

	btn = doc.Call("createElement", "BUTTON")
	btn.Set("innerText", "Random AI")
	btn.Call("setAttribute", "id", "randomAI")
	btn.Get("classList").Call("add", "btn")
	btn.Call("addEventListener", "click", js.FuncOf(setPlayer2))
	container.Call("appendChild", btn)

	player2 = <-player2Select

	// start game
	clear()
	js.Global().Set("play", js.FuncOf(play))
	for i := 0; i < 9; i++ {
		btn := doc.Call("createElement", "BUTTON")
		btn.Call("setAttribute", "id", fmt.Sprintf("btn%d", i+1))
		btn.Get("classList").Call("add", "box")
		btn.Set("innerHTML", "&nbsp;")
		btn.Call("addEventListener", "click", js.Global().Get("play"))

		container.Call("appendChild", btn)
	}

	// game loop
	for {
		var input string

		switch {
		case game.Turn() == 1:
			setMessage(player1.Title(1))
			input = player1.ChooseBox(*game)
		case game.Turn() == -1:
			setMessage(player2.Title(2))
			input = player2.ChooseBox(*game)
		}

		if game.IsValidMove(input) {
			game.Play(input)
			id := fmt.Sprintf("btn%s", input)
			doc.Call("getElementById", id).Set("innerText", getPlayerSign(-game.Turn()))
			doc.Call("getElementById", id).Set("disabled", true)
		} else {
			fmt.Println("Invalid answer received!")
			continue
		}

		if ok, who := game.GameOver(); ok {
			switch {
			case who == 0:
				setMessage("Draw!")
			case who == 1:
				setMessage("Player 1 Won!")
			case who == -1:
				setMessage("Player 2 Won!")
			}

			btn := doc.Call("createElement", "BUTTON")
			btn.Set("innerText", "Reload")
			btn.Call("setAttribute", "id", "reload")
			btn.Get("classList").Call("add", "btn")
			btn.Call("addEventListener", "click", js.FuncOf(reload))
			container.Call("appendChild", btn)
			break
		}
	}

	// block
	c := make(chan bool)
	<-c
}
