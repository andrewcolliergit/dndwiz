package main

import (
	"fmt"
	"os"

	"dndwiz/character"
	root "dndwiz/tui"
)

func main() {
	c := character.NewCharacter()
	cfg := root.Config{
		Character: &c,
	}

	p := root.NewProgram(&cfg)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
