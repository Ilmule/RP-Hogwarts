package src

import (
	"fmt"
	"strings"
)

// MenuGringotts gère l'interface et le coffre de la banque des sorciers
func MenuGringotts() {
	p := JoueurActuel // Utilise le joueur global
	reader := Lecteur

	for {
		fmt.Println("\n" + Yellow + "================ GOBELINS DE GRINGOTTS ================" + Reset)
		fmt.Printf("🪙 Sur vous         -> Gallions: %d | Mornilles: %d | Noises: %d\n", p.Gallions, p.Mornilles, p.Noises)
		fmt.Printf("🏦 Dans votre coffre -> Gallions: %d | Mornilles: %d | Noises: %d\n", p.CoffreGallions, p.CoffreMornilles, p.CoffreNoises)
		fmt.Println("-----------------------------------------------------")
		fmt.Println("1 - Déposer de l'argent dans le coffre")
		fmt.Println("2 - Retirer de l'argent du coffre")
		fmt.Println("3 - Ressortir de Gringotts")

		choix := LireSaisie(reader, "Que voulez-vous faire ? : ")

		switch strings.TrimSpace(choix) {
		case "1":
			DeposerArgent(p)
		case "2":
			RetirerArgent(p)
		case "3":
			fmt.Println("Les gobelins vous raccompagnent vers la sortie du Chemin de Traverse.")
			return
		case "inv" :
			ClearTerminal()
			AfficherInventaire(MenuGringotts)
		case "fournitures":
			ClearTerminal()
			Printlistefourniture()
			MenuGringotts()
		default:
			fmt.Println(Red + "Choix invalide, veuillez choisir une option valide." + Reset)
		}
	}
}

func DeposerArgent(p *Player) {
	fmt.Println("\n--- DÉPOSER DE L'ARGENT DANS LE COFFRE ---")
	var g, m, n int

	fmt.Print("Combien de Gallions déposer ? : ")
	_, err1 := fmt.Scanln(&g)
	fmt.Print("Combien de Mornilles déposer ? : ")
	_, err2 := fmt.Scanln(&m)
	fmt.Print("Combien de Noises déposer ? : ")
	_, err3 := fmt.Scanln(&n)

	if err1 != nil || err2 != nil || err3 != nil || g < 0 || m < 0 || n < 0 {
		fmt.Println(Red + "Montants invalides." + Reset)
		return
	}

	if g > p.Gallions || m > p.Mornilles || n > p.Noises {
		fmt.Println(Red + "Vous n'avez pas autant d'argent sur vous !" + Reset)
		return
	}

	p.Gallions -= g
	p.Mornilles -= m
	p.Noises -= n

	p.CoffreGallions += g
	p.CoffreMornilles += m
	p.CoffreNoises += n

	fmt.Println(Green + "Dépôt effectué avec succès dans votre coffre fort !" + Reset)
}

func RetirerArgent(p *Player) {
	fmt.Println("\n--- RETIRER DE L'ARGENT DU COFFRE ---")
	var g, m, n int

	fmt.Print("Combien de Gallions retirer ? : ")
	_, err1 := fmt.Scanln(&g)
	fmt.Print("Combien de Mornilles retirer ? : ")
	_, err2 := fmt.Scanln(&m)
	fmt.Print("Combien de Noises retirer ? : ")
	_, err3 := fmt.Scanln(&n)

	if err1 != nil || err2 != nil || err3 != nil || g < 0 || m < 0 || n < 0 {
		fmt.Println(Red + "Montants invalides." + Reset)
		return
	}

	if g > p.CoffreGallions || m > p.CoffreMornilles || n > p.CoffreNoises {
		fmt.Println(Red + "Votre coffre ne contient pas autant d'argent !" + Reset)
		return
	}

	p.CoffreGallions -= g
	p.CoffreMornilles -= m
	p.CoffreNoises -= n

	p.Gallions += g
	p.Mornilles += m
	p.Noises += n

	fmt.Println(Green + "Retrait effectué avec succès. L'argent est dans vos poches !" + Reset)
}