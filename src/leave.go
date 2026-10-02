package src

import (
	"os"
	"src/src/asciiart"
)

func Leave() {
	ClearTerminal()
	asciiart.PrintLeave()
	os.Exit(0)
}
