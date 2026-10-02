package src

import "os"

func Leave() {
	ClearTerminal()
	os.Exit(0)
}
