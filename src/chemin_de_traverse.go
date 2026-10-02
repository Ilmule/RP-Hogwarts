package src

import (
	"fmt"
)

func Chemindetraverse() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le Chemin de Traverse---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Gringotts ")
	fmt.Println()
	fmt.Println(" 2 - Ollivanders ")
	fmt.Println()
	fmt.Println(" 3 - Fleury et bott ")
	fmt.Println()
	fmt.Println(" 4 - L'apothicaire ")
	fmt.Println()
	fmt.Println(" 5 - Chez Guichard ")
	fmt.Println()
	fmt.Println(" 6 - La ménagerie magique ")
	fmt.Println()
	fmt.Println(" 7 - Boutique de Quidditch de qualité ")
	fmt.Println()
	fmt.Println(" 8 - Farces pour sorciers facétieux (Weasley & Weasley) ")
	fmt.Println()
	fmt.Println(" 9 - Glaces Fortescue ")
	fmt.Println()
	fmt.Println(" 10 - Aller dans l'allée des embrumes ")
	fmt.Println()
	fmt.Println(" 11 - Retourner au chaudron baveur ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1":
		ClearTerminal()
		MenuGringotts()
	case "2":
		ClearTerminal()
		LancerMagasinOllivander()
	case "3":
		ClearTerminal()
		LancerMagasinfleury()
	case "4":
		ClearTerminal()
		LancerMagasinapothicaire()
	case "5":
		ClearTerminal()
		LancerMagasinGuichard()
	case "6":
		ClearTerminal()
		LancerMagasinmenagerie()
	case "7":
		ClearTerminal()
		LancerMagasinquidditch()
	case "8":
		ClearTerminal()
		LancerMagasinfarces()
	case "9":
		ClearTerminal()
		LancerMagasinfortescue()
	case "10":
		ClearTerminal()
		Alleedesembrumes()
	case "11":
		ClearTerminal()
		Chaudronbaveur()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Chemindetraverse)
	case "fournitures":
		ClearTerminal()
		Printlistefourniture()
		Chemindetraverse()
	case "leave":
		ClearTerminal()
		Leave()
	default:
		ClearTerminal()
		Chemindetraverse()
	}
}
