package main

import (
	"fmt"
	"os"

	"dndwiz/character"
	"dndwiz/tui"
)

func main() {
	c := character.NewCharacter()
	cfg := tui.Config{
		Character: &c,
	}

	p := tui.NewProgram(&cfg)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
