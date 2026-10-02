package src

import (
	"fmt"
)

func Soussols() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Les cuisines ")
	fmt.Println()
	fmt.Println(" 2 - Salle de bain des préfets ")
	fmt.Println()
	fmt.Println(" 3 - Les cachots ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1" :
		ClearTerminal()
		Soussolscuisine()
	case "2" :
		ClearTerminal()
		Soussolsprefets()
	case "3" :
		ClearTerminal()
		Soussolscachots()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Soussols)
	default :
		ClearTerminal()
		Halldepoudlard()
	}
}
