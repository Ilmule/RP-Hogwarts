package src

import (
	"fmt"
	"src/src/asciiart"
)

func Creationperso() {
	ClearTerminal()
	reader := Lecteur
	joueur := CreatePlayerFromInput()
	InitJeu(joueur)
	ClearTerminal()
	fmt.Println(Red + "\n=== PERSONNAGE PRÊT ===" + Reset)
	fmt.Printf(Orange+"Bienvenue, %s!"+Reset, joueur.Name,)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	asciiart.PrintWelcometo()
	asciiart.PrintHogwarts()
	fmt.Println()
	fmt.Println()
	fmt.Println(Purple + "\nAppuyez sur Entrée pour valider votre personnage" + Reset)
	fmt.Println()
	reader.ReadString('\n')
	ClearTerminal()
	Debutjeu()
}
