package src

import (
	"fmt"
)

func Lestours() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - L'infirmerie ")
	fmt.Println()
	fmt.Println(" 2 - La salle des trophés ")
	fmt.Println()
	fmt.Println(" 3 - La tour d'astronomie ")
	fmt.Println()
	fmt.Println(" 4 - La salle de divination ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1" :
		ClearTerminal()
		Toursinfirmerie()
	case "2" :
		ClearTerminal()
		Tourstrophes()
	case "3" :
		ClearTerminal()
		Toursastronomie()
	case "4" :
		ClearTerminal()
		Toursdivination()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Lestours)
	default :
		ClearTerminal()
		Halldepoudlard()
	}
}
