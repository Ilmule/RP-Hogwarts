package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Sallecommunespoufsouffle() {
	p := JoueurActuel
	reader := Lecteur

	for {
		fmt.Println("\n" + Yellow + "=== SALLE COMMUNE DE POUFSOUFFLE ===" + Reset)
		fmt.Println("Une douce odeur de pâtisserie flotte dans l'air. La salle ronde ressemble à un immense terrier douillet.")
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Profiter du buffet avec les amis")
		fmt.Println(" 2 - Se prélasser dans un fauteuil jaune (Restaurer ses PV)")
		fmt.Println(" 3 - Travailler la Botanique (+ Expérience)")
		fmt.Println(" 4 - Faire une partie d'Échecs Version Sorcier")
		fmt.Println(" 5 - Soin du Mimbulus (Mini-Jeu)")
		fmt.Println(" 6 - Aller dans son dortoir (Changer de vêtements / Dormir)")
		fmt.Println(" R - Sortir dans le Hall")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			dialogues := []string{
				"Cedric : « Si tu as besoin d'aide pour le sortilège d'Allégresse, n'hésite pas ! »",
				"Hannah : « Les elfes ont préparé des choux farcis ce midi, j'ai hâte ! »",
				"Ernie : « J'ai compté les sabliers, nous avons presque rattrapé Serdaigle ! »",
			}
			fmt.Println(Cyan + dialogues[rand.Intn(len(dialogues))] + Reset)
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		case "2":
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Green + "Le confort de la salle vous redonne toutes vos forces." + Reset)
		case "3":
			ClearTerminal()
			gainExp := rand.Intn(10) + 5
			p.Exp += gainExp
			fmt.Printf(Yellow+"Vous étudiez les propriétés des mandragores (+%d Exp).\n"+Reset, gainExp)
		case "4":
			ClearTerminal()
			JeuEchecsSorcier()
		case "5":
			ClearTerminal()
			SoinPlante()
		case "6":
			ClearTerminal()
			MenuDortoirPoufsouffle()
		case "R":
			ClearTerminal()
			Lessallescommunes()
			return
		}
	}
}

func SoinPlante() {
	p := JoueurActuel
	reader := Lecteur
	fmt.Println(Yellow + "=== Soin du Mimbulus Mimbletonia ===" + Reset)
	fmt.Println("La plante tremble. Si vous la caressez au mauvais moment, elle vous aspergera de pus !")
	time.Sleep(1 * time.Second)
	fmt.Println("3... 2... 1...")
	
	chance := rand.Intn(100)
	if chance > 40 {
		fmt.Println(Green + "Vous calmez la plante doucement. Elle ronronne presque (+10 Exp)." + Reset)
		p.Exp += 10
	} else {
		fmt.Println(Red + "SPOUTCH ! Vous êtes recouvert de pus d'empestine nauséabond !" + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func MenuDortoirPoufsouffle() {
	p := JoueurActuel
	reader := Lecteur
	for {
		fmt.Println("\n" + Yellow + "=== DORTOIR DE POUFSOUFFLE ===" + Reset)
		fmt.Println(" 1 - Fouiller dans son armoire (Changer de vêtements)")
		fmt.Println(" 2 - Dormir au chaud (Soigne totalement)")
		fmt.Println(" R - Retourner au salon")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nAction : ")))
		if choix == "1" {
			ClearTerminal()
			ChangerVetementsDortoir()
		} else if choix == "2" {
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Cyan + "Vous vous endormez paisiblement dans votre lit douillet." + Reset)
		} else if choix == "R" {
			ClearTerminal()
			return
		}
	}
}