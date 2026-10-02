package src

import (
	"fmt"
	"strings"
)

func Kingscross() {
	reader := Lecteur
	p := JoueurActuel

	fmt.Println(Red + "--------------- King's Cross ---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Aller au Chaudron Baveur")
	fmt.Println()
	fmt.Println(" 2 - Traverser la barrière vers la Voie 9 ¾ (Poudlard)")
	fmt.Println()
	fmt.Println(Red + "--------------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
		
	case "1":
		ClearTerminal()
		Chaudronbaveur()

	case "2":
		ClearTerminal()

		// Vérification du Ticket pour la Voie 9 ¾ dans l'inventaire
		aLeTicket := false
		for _, item := range p.Inventaire {
			nomLower := strings.ToLower(item.Nom)
			if strings.Contains(nomLower, "ticket") && strings.Contains(nomLower, "9") {
				aLeTicket = true
				break
			}
		}

		if aLeTicket {
			fmt.Println(Green + "Vous présentez votre Ticket doré pour la Voie 9 ¾." + Reset)
			fmt.Println(Yellow + "Vous foncez à travers le mur de briques entre les voies 9 et 10..." + Reset)
			fmt.Println(Cyan + "Le Poudlard Express vous attend dans une fumée écarlate ! En route pour le château !" + Reset)
			AnimerPoudlardExpress()
			fmt.Println()
			Halldepoudlard()
		} else {
			fmt.Println(Red + "=== ACCÈS REFUSÉ À LA VOIE 9 ¾ ===" + Reset)
			fmt.Println("Hagrid vous retient par l'épaule :")
			fmt.Println(Yellow + "« Holà gamin ! Tu vas où comme ça ?! Tu peux pas monter dans l'Poudlard Express sans ton ticket pour la voie 9 ¾ ! »" + Reset)
			fmt.Println(Red + "« Va d'abord vérifier tes fournitures avec moi pour que je te donne ton billet ! »" + Reset)
			fmt.Println()
			Kingscross()
		}

	case "inv":
		ClearTerminal()
		AfficherInventaire(Kingscross)

	default:
		ClearTerminal()
		Kingscross()
	}
}