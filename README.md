# Tic Tac Toe

## Run in Terminal

### Build
`make build-terminal`

### Run
`./bin/terminal`

> ~~In this mode computer is unbeatable AI.~~ In this update you can select user input / unbeatable / random

> No external package used.

## Run in Browser

`make run-server`

Then goto http://127.0.0.1:8080

## AIs

### Unbeatable AI

I use [MiniMax Algorithm](https://en.wikipedia.org/wiki/Minimax) to write this AI.

I found 3 other algorithms as below:

1. Tabular Q-Learning
1. MCTS
1. Neural Network

But I implemented minimax only :D

### Random AI

This AI randomly select an available box to play.

## Test AIs

I wrote 3 games in [test](./selector/ai_test.go):

1. Unbeatable AI versus Unbeatable AI (No one can beat each other)
1. Unbeatable (as X) versus  (as O) (Random can't beat Unbeatable)
1. Random (as X) versus Unbeatable (as O) (Random can't beat Unbeatable)

I ran these tests 1000 times:

`go test -count 1000`

and no one failed.
