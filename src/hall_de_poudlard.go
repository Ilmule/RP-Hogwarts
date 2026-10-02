package src

import (
	"fmt"
)

func Halldepoudlard() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - La grande salle")
	fmt.Println()
	fmt.Println(" 2 - Les salles communes ")
	fmt.Println()
	fmt.Println(" 3 - Les salles de cour ")
	fmt.Println()
	fmt.Println(" 4 - Les sous sols ")
	fmt.Println()
	fmt.Println(" 5 - Les tours ")
	fmt.Println()
	fmt.Println(" 6 - L'extérieur ")
	fmt.Println()
	fmt.Println(" 7 - Les lieux secrets ")
	fmt.Println()
	fmt.Println(" 8 - Retourner à Pré au lard ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1":
		ClearTerminal()
		GrandeSalle()
	case "2":
		ClearTerminal()
		Lessallescommunes()
	case "3":
		ClearTerminal()
		Sallesdecours()
	case "4":
		ClearTerminal()
		Soussols()
	case "5":
		ClearTerminal()
		Lestours()
	case "6":
		ClearTerminal()
		Lexterieur()
	case "7":
		ClearTerminal()
		Lieuxsecrets()
	case "8":
		ClearTerminal()
		Preaulard()
	case "inv" : 
		ClearTerminal()
		AfficherInventaire(Halldepoudlard)
	default:
		ClearTerminal()
		Halldepoudlard()
	}
}
