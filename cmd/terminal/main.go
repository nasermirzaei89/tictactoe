package main

import (
	"fmt"
	"github.com/nasermirzaei89/tictactoe"
	"github.com/nasermirzaei89/tictactoe/selector"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Loading...")

	game := tictactoe.New()

	var (
		p1 selector.AI
		p2 selector.AI
	)

	// select player 1
	for {
		time.Sleep(time.Second / 2)
		render(*game)
		fmt.Print(`Choose Player 1:
(1) User input
(2) Unbeatable AI
(3) Random AI
`)

		var input string
		_, _ = fmt.Scan(&input)
		if input == "1" {
			p1 = selector.NewInput(os.Stdin)
			break
		}

		if input == "2" {
			p1 = selector.NewUnbeatable()
			break
		}

		if input == "3" {
			p1 = selector.NewRandom()
			break
		}
		fmt.Println("Invalid answer received!")
		continue
	}

	// select player 2
	for {
		time.Sleep(time.Second / 2)
		render(*game)
		fmt.Print(`Choose Player 2:
(1) User input
(2) Unbeatable AI
(3) Random AI
`)

		var input string
		_, _ = fmt.Scan(&input)
		if input == "1" {
			p2 = selector.NewInput(os.Stdin)
			break
		}

		if input == "2" {
			p2 = selector.NewUnbeatable()
			break
		}

		if input == "3" {
			p2 = selector.NewRandom()
			break
		}
		fmt.Println("Invalid answer received!")
		continue
	}

	// game loop
	for {
		time.Sleep(time.Second / 2)
		render(*game)

		var input string

		switch {
		case game.Turn() == 1:
			fmt.Println(p1.Title(1))
			fmt.Println(p1.Prompt())
			input = p1.ChooseBox(*game)
		case game.Turn() == -1:
			fmt.Println(p2.Title(2))
			fmt.Println(p2.Prompt())
			input = p2.ChooseBox(*game)
		}

		if game.IsValidMove(input) {
			game.Play(input)
		} else {
			fmt.Println("Invalid answer received!")
			continue
		}

		if ok, who := game.GameOver(); ok {
			render(*game)
			switch {
			case who == 0:
				fmt.Println("Draw!")
			case who == 1:
				fmt.Println("Player 1 Won!")
			case who == -1:
				fmt.Println("Player 2 Won!")
			}
			break
		}
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	default:
		log.Fatalln(fmt.Sprintf("platform '%s' is not supported", runtime.GOOS))
	}
}

func render(game tictactoe.Game) {
	clearScreen()
	fmt.Print(`
┏━━━━━━━━━━━━━┓
┃ Tic Tac Toe ┃
┣┳━━━┯━━━┯━━━┳┫
`)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := game.Board()[i][j]
			switch j {
			case 0:
				fmt.Printf("┃┃ %s ", renderValue(i, j, v))
			case 1:
				fmt.Printf("│ %s │", renderValue(i, j, v))
			case 2:
				fmt.Printf(" %s ┃┃", renderValue(i, j, v))
			}
		}

		if i < 2 {
			fmt.Println("\n┣╉───┼───┼───╊┫")
		} else {
			fmt.Println("\n┗┻━━━┷━━━┷━━━┻┛")
		}
	}
}

func renderValue(i, j, v int) string {
	if v == 1 {
		return "×"
	}

	if v == -1 {
		return "○"
	}

	// v == 0
	return strconv.Itoa(i*3 + j + 1)
}
