package src

import (
	"fmt"
)

func Preaulard() {
	reader := Lecteur
	fmt.Println(Red + "---------------Pré au lard---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Retourner à King's Cross ")
	fmt.Println()
	fmt.Println(" 2 - Aller à Poudlard ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1":
		ClearTerminal()
		Kingscross()
	case "2":
		ClearTerminal()
		Halldepoudlard()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Preaulard)
	default:
		ClearTerminal()
		Preaulard()
	}
}
