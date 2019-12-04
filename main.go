package main

import (
	"github.com/IacopoMelani/musicgang/command"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	command.InterpretingHumanWord()
}
