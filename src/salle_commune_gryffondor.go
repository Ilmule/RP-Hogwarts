package src

import (
	"fmt"
	"math/rand"
	"strings"
)

func Sallecommunesgryffondor() {
	p := JoueurActuel
	reader := Lecteur

	for {
		fmt.Println("\n" + Red + "=== SALLE COMMUNE DE GRYFFONDOR ===" + Reset)
		fmt.Println("Le feu crépite dans l'immense cheminée. La pièce est chaleureuse et remplie de fauteuils rouges moelleux.")
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Discuter avec des camarades")
		fmt.Println(" 2 - Se poser au coin du feu (Restaurer ses PV)")
		fmt.Println(" 3 - Travailler sur ses leçons (+ Expérience)")
		fmt.Println(" 4 - Faire une partie d'Échecs Version Sorcier")
		fmt.Println(" 5 - Jouer aux Bavboules de Feu (Mini-Jeu)")
		fmt.Println(" 6 - Monter dans le dortoir (Changer de vêtements / Dormir)")
		fmt.Println(" R - Sortir dans le Hall")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			dialogues := []string{
				"Neville : « Tu n'aurais pas vu mon crapaud Trevor ? Je l'ai encore perdu... »",
				"Seamus : « Regarde ça ! Je crois que j'ai enfin réussi à transformer cette eau en rhum... *BOUM* »",
				"Hermione : « Tu as déjà fait ton devoir pour le professeur McGonagall ? Il faut au moins deux rouleaux de parchemin ! »",
			}
			fmt.Println(Yellow + dialogues[rand.Intn(len(dialogues))] + Reset)
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		case "2":
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Green + "La chaleur du feu vous revigore. (PV restaurés au maximum !)" + Reset)
		case "3":
			ClearTerminal()
			gainExp := rand.Intn(10) + 5
			p.Exp += gainExp
			fmt.Printf(Cyan+"Vous révisez la Métamorphose et gagnez %d points d'expérience !\n"+Reset, gainExp)
		case "4":
			ClearTerminal()
			JeuEchecsSorcier()
		case "5":
			ClearTerminal()
			JeuBavboules()
		case "6":
			ClearTerminal()
			MenuDortoirGryffondor()
		case "R":
			ClearTerminal()
			Lessallescommunes()
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

func JeuBavboules() {
	p := JoueurActuel
	reader := Lecteur
	fmt.Println(Yellow + "=== Lancer de Bavboules de Feu ===" + Reset)
	fmt.Println("Le but : approcher le score de 100 sans le dépasser, sinon la bavboule explose !")
	score := 0
	for {
		jet := rand.Intn(30) + 10
		score += jet
		fmt.Printf("Vous lancez une bavboule... Score actuel : %d\n", score)
		if score > 100 {
			fmt.Println(Red + "BOUM ! La bavboule explose et vous recouvre de liquide puant. Perdu !" + Reset)
			break
		} else if score >= 90 {
			fmt.Println(Green + "Magnifique ! Vous remportez 2 Gallions !" + Reset)
			p.Gallions += 2
			break
		}
		choix := LireSaisie(reader, "Relancer ? (o/n) : ")
		if strings.ToLower(strings.TrimSpace(choix)) != "o" {
			fmt.Printf("Vous vous arrêtez à %d. Pas mal !\n", score)
			break
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func MenuDortoirGryffondor() {
	p := JoueurActuel
	reader := Lecteur
	for {
		fmt.Println("\n" + Red + "=== DORTOIR DE GRYFFONDOR ===" + Reset)
		fmt.Println(" 1 - Fouiller sa malle (Changer de vêtements)")
		fmt.Println(" 2 - Dormir jusqu'au lendemain matin (Soigne totalement)")
		fmt.Println(" R - Redescendre")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nAction : ")))
		if choix == "1" {
			ClearTerminal()
			ChangerVetementsDortoir()
		} else if choix == "2" {
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Cyan + "Vous passez une bonne nuit... Vous êtes en pleine forme !" + Reset)
		} else if choix == "R" {
			ClearTerminal()
			return
		}
	}
}