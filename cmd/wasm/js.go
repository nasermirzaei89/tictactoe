package main

import (
	"github.com/nasermirzaei89/tictactoe/selector"
	"strings"
	"syscall/js"
)

func setPlayer1(this js.Value, args []js.Value) interface{} {
	id := this.Get("id").String()
	switch id {
	case "userInput":
		player1Select <- selector.NewInput(ui)
	case "unbeatableAI":
		player1Select <- selector.NewUnbeatable()
	case "randomAI":
		player1Select <- selector.NewRandom()
	}

	return nil
}

func setPlayer2(this js.Value, args []js.Value) interface{} {
	id := this.Get("id").String()
	switch id {
	case "userInput":
		player2Select <- selector.NewInput(ui)
	case "unbeatableAI":
		player2Select <- selector.NewUnbeatable()
	case "randomAI":
		player2Select <- selector.NewRandom()
	}

	return nil
}

func play(this js.Value, args []js.Value) interface{} {
	if ok, _ := game.GameOver(); ok {
		return nil
	}

	id := this.Get("id").String()
	val := strings.TrimPrefix(id, "btn")

	_, _ = ui.Write([]byte(val))
	return nil
}

func reload(_ js.Value, _ []js.Value) interface{} {
	js.Global().Get("location").Call("reload")
	return nil
}
