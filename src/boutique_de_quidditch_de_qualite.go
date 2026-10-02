package src

import (
	"fmt"
	"strings"
)

func LancerMagasinquidditch() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
		// === ÉQUIPEMENT & MATÉRIEL ===
		{Nom: "Livre : Le Quidditch à travers les âges", PrixGallions: 2, PrixMornilles: 5, PrixNoises: 0},
		{Nom: "Lot de balles de Quidditch (1 Souaffle, 2 Cognards, 1 Vif d'or)", PrixGallions: 25, PrixMornilles: 10, PrixNoises: 0},
		{Nom: "Kit d'entretien pour balais", PrixGallions: 5, PrixMornilles: 15, PrixNoises: 0},
		{Nom: "Lunettes de vol (anti-sortilège & anti-pluie)", PrixGallions: 4, PrixMornilles: 7, PrixNoises: 0},
		{Nom: "Tenue complète de vol / Quidditch", PrixGallions: 10, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Boussole pour balais magiques", PrixGallions: 3, PrixMornilles: 23, PrixNoises: 0},
		// === GOODIES ===
		{Nom: "Fanion aux couleurs des Canons de Chudley", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Écharpe de supporter des Harpies de Holyhead", PrixGallions: 2, PrixMornilles: 4, PrixNoises: 0},
		{Nom: "Insigne officiel des Frelons de Wimbourne", PrixGallions: 1, PrixMornilles: 9, PrixNoises: 0},
		{Nom: "Figurine animée de Viktor Krum", PrixGallions: 5, PrixMornilles: 20, PrixNoises: 0},
		{Nom: "Chapeau rugissant en forme de tête de lion", PrixGallions: 4, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Rostres & Sifflets magiques de supporter", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		// === BALAIS ===
		{Nom: "L'Éclair de Feu Suprême (Firebolt Supreme)", PrixGallions: 350, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "L'Éclair de Feu (Firebolt)", PrixGallions: 250, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Nimbus 2001", PrixGallions: 100, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Nimbus 2000", PrixGallions: 50, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Flèche d'Argent (Silver Arrow)", PrixGallions: 30, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Comète 260", PrixGallions: 15, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Comète 220", PrixGallions: 12, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Comète 180", PrixGallions: 10, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Comète 140", PrixGallions: 8, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 11 (Cleansweep 11)", PrixGallions: 14, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 10", PrixGallions: 13, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 9", PrixGallions: 12, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 8", PrixGallions: 11, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 7", PrixGallions: 10, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 6", PrixGallions: 9, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 5", PrixGallions: 7, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 4", PrixGallions: 6, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 3", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 2", PrixGallions: 4, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brossdur 1", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Le Taureau Volant (Oakshaft 79)", PrixGallions: 18, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Le Poussière de Lune (Moondreamer)", PrixGallions: 15, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Le Swiftstick", PrixGallions: 11, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "L'Arbalète 360 (Tinderblast)", PrixGallions: 9, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "L'Étoile Filante / Mancheàbalai (Shooting Star)", PrixGallions: 4, PrixMornilles: 0, PrixNoises: 0},
	}

	for {
		fmt.Println("\n" + Yellow + "=== BOUTIQUE DE QUIDDITCH DE QUALITÉ ===" + Reset)
		fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)
		for i, item := range catalogue {
			fmt.Printf("%2d - %-60s (%3dG %2dM %2dN)\n", i+1, item.Nom, item.PrixGallions, item.PrixMornilles, item.PrixNoises)
		}
		fmt.Println(" Q - Quitter")

		choix := LireSaisie(reader, "Que voulez-vous acheter ? : ")
		if strings.ToLower(strings.TrimSpace(choix)) == "q" {
			ClearTerminal()
			Chemindetraverse()
			break
		}

		if strings.ToLower(strings.TrimSpace(choix)) == "inv" {
			ClearTerminal()
			AfficherInventaire(LancerMagasinquidditch)
		}

		if strings.ToLower(strings.TrimSpace(choix)) == "fournitures" {
			ClearTerminal()
			Printlistefourniture()
		}

		if strings.ToLower(strings.TrimSpace(choix)) == "leave" {
			ClearTerminal()
			Leave()
		}

		var idx int
		_, err := fmt.Sscan(choix, &idx)
		if err != nil || idx < 1 || idx > len(catalogue) {
			continue
		}

		article := catalogue[idx-1]
		if p.Payer(article.PrixGallions, article.PrixMornilles, article.PrixNoises) {
			p.AjouterItem(article.Nom, TypeEquipement, article.PrixGallions, 1)
			ClearTerminal()
			fmt.Printf(Green+"Vous avez acheté : %s !\n"+Reset, article.Nom)
		} else {
			fmt.Println(Red + "Vous n'avez pas assez d'argent !" + Reset)
		}
	}
}