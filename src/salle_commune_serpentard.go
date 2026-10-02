package src

import (
	"fmt"
	"math/rand"
	"strings"
)

func Sallecommuneserpentard() {
	p := JoueurActuel
	reader := Lecteur

	for {
		fmt.Println("\n" + Green + "=== SALLE COMMUNE DE SERPENTARD ===" + Reset)
		fmt.Println("L'atmosphère est fraîche. La lueur verte des lampes se reflète sur les vitres donnant sous le Lac Noir.")
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Écouter les rumeurs (Discuter)")
		fmt.Println(" 2 - Se réchauffer près de la cheminée sculptée (Restaurer ses PV)")
		fmt.Println(" 3 - Étudier les Potions (+ Expérience)")
		fmt.Println(" 4 - Faire une partie d'Échecs Version Sorcier")
		fmt.Println(" 5 - Jouer au Poker Menteur Magique (Mini-Jeu)")
		fmt.Println(" 6 - Descendre dans le dortoir (Changer de vêtements / Dormir)")
		fmt.Println(" R - Sortir dans le Hall")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			dialogues := []string{
				"Drago : « Mon père en entendra parler si Dumbledore continue de favoriser Potter ! »",
				"Pansy : « As-tu vu la nouvelle coupe de cheveux de Miss Je-Sais-Tout ? Hilarant. »",
				"Blaise : « Il paraît que les cachots s'étendent bien plus loin qu'on ne le pense sous le lac... »",
			}
			fmt.Println(Yellow + dialogues[rand.Intn(len(dialogues))] + Reset)
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		case "2":
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Green + "Le feu vert pâle soigne vos blessures. (PV restaurés)" + Reset)
		case "3":
			ClearTerminal()
			gainExp := rand.Intn(10) + 5
			p.Exp += gainExp
			fmt.Printf(Cyan+"Vous révisez vos potions avec ferveur (+%d Exp) !\n"+Reset, gainExp)
		case "4":
			ClearTerminal()
			JeuEchecsSorcier()
		case "5":
			ClearTerminal()
			PokerMenteur()
		case "6":
			ClearTerminal()
			MenuDortoirSerpentard()
		case "R":
			ClearTerminal()
			Lessallescommunes()
			return
		}
	}
}

func PokerMenteur() {
	p := JoueurActuel
	reader := Lecteur
	fmt.Println(Yellow + "=== Poker Menteur Sorcier ===" + Reset)
	fmt.Println("Un 7ème année vous affirme qu'il détient une carte 'Dragon Noir'. Ment-il ?")
	choix := LireSaisie(reader, "1 - Il bluffe (Dénoncer) | 2 - Il dit vrai (Se coucher) : ")
	
	bluff := rand.Intn(2) == 0 // 50% de chance qu'il mente
	if (choix == "1" && bluff) || (choix == "2" && !bluff) {
		fmt.Println(Green + "Bien vu ! Vous l'avez percé à jour. Vous remportez 5 Gallions !" + Reset)
		p.Gallions += 5
	} else {
		fmt.Println(Red + "Faux ! Il ricane et empoche la mise. Vous avez été manipulé." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func MenuDortoirSerpentard() {
	p := JoueurActuel
	reader := Lecteur
	for {
		fmt.Println("\n" + Green + "=== DORTOIR DE SERPENTARD ===" + Reset)
		fmt.Println(" 1 - Ouvrir sa malle (Changer de vêtements)")
		fmt.Println(" 2 - Dormir (Soigne totalement)")
		fmt.Println(" R - Retourner au salon")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nAction : ")))
		if choix == "1" {
			ClearTerminal()
			ChangerVetementsDortoir()
		} else if choix == "2" {
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Cyan + "Le bruit de l'eau du lac vous berce. (PV max)" + Reset)
		} else if choix == "R" {
			ClearTerminal()
			return
		}
	}
}