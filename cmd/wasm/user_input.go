package main

import "io"

type userInput chan []byte

func (ui userInput) Read(p []byte) (int, error) {
	in := <-ui
	if len(in) > 0 {
		copy(p, in)
		return len(in), nil
	}

	return 0, io.EOF
}

func (ui *userInput) Write(p []byte) (n int, err error) {
	*ui <- p
	*ui <- []byte{}
	return len(p), nil
}
