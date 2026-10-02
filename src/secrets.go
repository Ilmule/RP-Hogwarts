package src

import (
	"fmt"
)

func Lieuxsecrets() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Le bureau du directeur ")
	fmt.Println()
	fmt.Println(" 2 - La salle sur demande ")
	fmt.Println()
	fmt.Println(" 3 - Les toilettes de Mimi Geignarde ")
	fmt.Println()
	fmt.Println(" 4 - La chambre des secrets ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1" :
		ClearTerminal()
		Secretdirecteur()
	case "2" :
		ClearTerminal()
		Secretssallesurdemande()
	case "3" :
		ClearTerminal()
		Secretsmimigeignarde()
	case "4" :
		ClearTerminal()
		Secretchambredessecrets()
	case "inv" :
		ClearTerminal()
		AfficherInventaire(Lieuxsecrets)
	default :
		ClearTerminal()
		Halldepoudlard()
	}
}
