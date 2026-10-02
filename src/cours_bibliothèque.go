package src

import (
	"fmt"
)

func Coursbiblio() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Retourner au Hall ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray + "Dans quelle salle de cours voulez-vous vous rendre ? : " + Reset)

	switch Option {
	case "1" :
		ClearTerminal()
		Halldepoudlard()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Coursbiblio)
	default :
		ClearTerminal()
		Halldepoudlard()
	}
}