package src

import (
	"fmt"
)

func Lexterieur() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - La cabane d'Hagrid ")
	fmt.Println()
	fmt.Println(" 2 - Le terrain de quidditch ")
	fmt.Println()
	fmt.Println(" 3 - La volière ")
	fmt.Println()
	fmt.Println(" 4 - La forêt interdite ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray + "Que voulez-vous faire ? : " + Reset)

	switch Option {
	case "1" :
		ClearTerminal()
		Exterieurhagrid()
	case "2" :
		ClearTerminal()
		Exterieurquidditch()
	case "3" :
		ClearTerminal()
		Exterieurvoliere()
	case "4" :
		ClearTerminal()
		Exterieurforetinterdite()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Lexterieur)
	default :
		ClearTerminal()
		Lexterieur()
	}
}
