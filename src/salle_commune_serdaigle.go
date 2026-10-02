package src

import (
	"fmt"
	"math/rand"
	"strings"
)

func Sallecommunesserdaigle() {
	p := JoueurActuel
	reader := Lecteur

	for {
		fmt.Println("\n" + Cyan + "=== SALLE COMMUNE DE SERDAIGLE ===" + Reset)
		fmt.Println("La pièce est aérée, tapissée de soie bleue, avec des bibliothèques recouvrant chaque mur.")
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Discuter avec les esprits brillants")
		fmt.Println(" 2 - Se détendre en lisant sous les étoiles (Restaurer ses PV)")
		fmt.Println(" 3 - Travailler sur l'Histoire de la Magie (+ Expérience)")
		fmt.Println(" 4 - Faire une partie d'Échecs Version Sorcier")
		fmt.Println(" 5 - Défi de l'Aigle (Mini-Jeu d'énigme)")
		fmt.Println(" 6 - Monter dans le dortoir (Changer de vêtements / Dormir)")
		fmt.Println(" R - Sortir dans le Hall")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			dialogues := []string{
				"Luna : « Tes joncheruines sont particulièrement agitées aujourd'hui. »",
				"Cho : « J'espère que l'entraînement de Quidditch ne sera pas annulé avec ce vent. »",
				"Terry : « Je viens de lire un ouvrage fascinant sur la métamorphose élémentaire... »",
			}
			fmt.Println(Yellow + dialogues[rand.Intn(len(dialogues))] + Reset)
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		case "2":
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Green + "Le calme de la salle vous apaise complètement. (PV restaurés)" + Reset)
		case "3":
			ClearTerminal()
			gainExp := rand.Intn(10) + 5
			p.Exp += gainExp
			fmt.Printf(Cyan+"Votre soif d'apprendre vous rapporte %d Exp !\n"+Reset, gainExp)
		case "4":
			ClearTerminal()
			JeuEchecsSorcier()
		case "5":
			ClearTerminal()
			DefiAigle()
		case "6":
			ClearTerminal()
			MenuDortoirSerdaigle()
		case "R":
			ClearTerminal()
			Lessallescommunes()
			return
		}
	}
}

func DefiAigle() {
	p := JoueurActuel
	reader := Lecteur
	fmt.Println(Cyan + "=== L'Énigme du Heurtoir ===" + Reset)
	
	enigmes := []struct { q, r string }{
		{"Je ne respire jamais mais j'ai beaucoup d'esprit. Qui suis-je ?", "fantome"},
		{"Plus elle est grande, moins on la voit. Qui est-elle ?", "obscurite"},
	}
	e := enigmes[rand.Intn(len(enigmes))]
	
	fmt.Println("« " + e.q + " »")
	reponse := strings.ToLower(strings.TrimSpace(LireSaisie(reader, "Votre réponse : ")))
	
	if strings.Contains(reponse, e.r) || strings.Contains(reponse, "fantôme") || strings.Contains(reponse, "obscurité") {
		fmt.Println(Green + "« Bien raisonné. » Vous gagnez en sagesse (+15 Exp)." + Reset)
		p.Exp += 15
	} else {
		fmt.Println(Red + "« La logique vous échappe... »" + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func MenuDortoirSerdaigle() {
	p := JoueurActuel
	reader := Lecteur
	for {
		fmt.Println("\n" + Cyan + "=== DORTOIR DE SERDAIGLE ===" + Reset)
		fmt.Println(" 1 - Parcourir sa penderie (Changer de vêtements)")
		fmt.Println(" 2 - Dormir (Soigne totalement)")
		fmt.Println(" R - Redescendre")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nAction : ")))
		if choix == "1" {
			ClearTerminal()
			ChangerVetementsDortoir()
		} else if choix == "2" {
			ClearTerminal()
			p.Health = p.MaxHealth
			fmt.Println(Cyan + "Sous les draps de soie bleue, vous récupérez toute votre énergie." + Reset)
		} else if choix == "R" {
			ClearTerminal()
			return
		}
	}
}