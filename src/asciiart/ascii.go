package asciiart

import (
	_ "embed"
	"fmt"
)

//go:embed welcometo.txt
var asciiWelcometo string
func PrintWelcometo() {
	fmt.Println(Welcometo + asciiWelcometo + Reset)
}

//go:embed Hogwarts.txt
var asciiHogwarts string
func PrintHogwarts() {
	fmt.Println(Hogwarts + asciiHogwarts + Reset)
}

//go:embed Leave.txt
var asciiLeave string
func PrintLeave() {
	fmt.Println(Purple + asciiLeave + Reset)
}