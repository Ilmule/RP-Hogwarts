package src 

import (
	"fmt"
)

func Alleedesembrumes() {
	reader := Lecteur
	fmt.Println(Red + "---------------King's Cross---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Borgin et Burkes ")
	fmt.Println()
	fmt.Println(" 2 - La boutique de peaux d'anguilles ")
	fmt.Println()
	fmt.Println(" 3 - Le magasin de la vieille sorcière ")
	fmt.Println()
	fmt.Println(" 4 - Retourner au Chemin de Traverse ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1":
		ClearTerminal()
		Chemindetraverse()
	case "leave":
		ClearTerminal()
		Leave()
	case "inv" : 
		AfficherInventaire(Alleedesembrumes)
	default:
		ClearTerminal()
		Alleedesembrumes()
	}
}
