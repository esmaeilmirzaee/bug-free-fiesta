package main

import (
	"myapp/game"
)

func main() {
	playAgain := true

	for playAgain {
		playAgain = game.GetYesOrNo("Would you like play again(y/n)?")
	}

}
