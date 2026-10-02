package src

import (
	"fmt"
)

func Sallesdecours() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Salle de Potions ")
	fmt.Println()
	fmt.Println(" 2 - Serre de botanique ")
	fmt.Println()
	fmt.Println(" 3 - Salle de sortilèges ")
	fmt.Println()
	fmt.Println(" 4 - Salle de métamorphose ")
	fmt.Println()
	fmt.Println(" 5 - La bibliothèque ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray + "Dans quelle salle de cours voulez-vous vous rendre ? : " + Reset)

	switch Option {
	case "1" :
		ClearTerminal()
		Courspotions()
	case "2" :
		ClearTerminal()
		Coursbotanique()
	case "3" :
		ClearTerminal()
		Courssortilège()
	case "4" :
		ClearTerminal()
		Coursmétamorphose()
	case "5" :
		ClearTerminal()
		Coursbiblio()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Sallesdecours)
	default :
		ClearTerminal()
		Sallesdecours()
	}
}
