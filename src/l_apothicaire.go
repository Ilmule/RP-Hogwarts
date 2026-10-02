package src

import (
	"fmt"
	"strings"
)

func LancerMagasinapothicaire() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
		// === MATÉRIEL DE POTIONS ===
		{Nom: "Chaudron en étain (taille 2 - modèle standard)", PrixGallions: 15, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chaudron en cuivre (chauffage plus rapide)", PrixGallions: 21, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chaudron en laiton massif", PrixGallions: 25, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chaudron en or massif (pur luxe)", PrixGallions: 100, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Boîte de 10 fioles en verre épais", PrixGallions: 0, PrixMornilles: 10, PrixNoises: 0},
		{Nom: "Boîte de 10 fioles en cristal fin", PrixGallions: 2, PrixMornilles: 5, PrixNoises: 0},
		{Nom: "Balance en cuivre avec poids de précision", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Kit de préparation : Mortier et pilon en granit", PrixGallions: 1, PrixMornilles: 5, PrixNoises: 0},
		{Nom: "Couteau à lame d'argent", PrixGallions: 2, PrixMornilles: 8, PrixNoises: 0},
		
		// === MATÉRIEL D'ASTRONOMIE (NOUVEAU) ===
		{Nom: "Télescope en laiton (modèle standard)", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Télescope pliable enchanté (zoom automatique)", PrixGallions: 12, PrixMornilles: 8, PrixNoises: 0},
		{Nom: "Télescope en or avec filtre stellaire", PrixGallions: 35, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Carte du ciel interactive (mise à jour magique)", PrixGallions: 0, PrixMornilles: 12, PrixNoises: 15},
		{Nom: "Globe lunaire miniature de bureau", PrixGallions: 3, PrixMornilles: 4, PrixNoises: 0},
		{Nom: "Astrolabe de précision en cuivre", PrixGallions: 4, PrixMornilles: 10, PrixNoises: 0},

		// === INGRÉDIENTS D'ORIGINE ANIMALE ===
		{Nom: "Cornes de Bicorne en poudre (le sachet)", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Peau de Serpent du Cap (le mètre)", PrixGallions: 3, PrixMornilles: 5, PrixNoises: 0},
		{Nom: "Yeux de scarabées noirs (la poignée)", PrixGallions: 0, PrixMornilles: 2, PrixNoises: 15},
		{Nom: "Foie de dragon d'Australie (le morceau)", PrixGallions: 0, PrixMornilles: 16, PrixNoises: 0},
		{Nom: "Crocs de serpent venimeux pulvérisés", PrixGallions: 0, PrixMornilles: 12, PrixNoises: 0},
		{Nom: "Chenilles velues séchées (le pot)", PrixGallions: 0, PrixMornilles: 3, PrixNoises: 10},
		{Nom: "Sangsues fraîches des marais", PrixGallions: 0, PrixMornilles: 5, PrixNoises: 0},
		{Nom: "Bave de Crapaud concentrée (la fiole)", PrixGallions: 0, PrixMornilles: 4, PrixNoises: 20},
		{Nom: "Plumes de Floperce", PrixGallions: 1, PrixMornilles: 2, PrixNoises: 0},
		{Nom: "Bézoard (Antidote puissant)", PrixGallions: 10, PrixMornilles: 0, PrixNoises: 0},

		// === INGRÉDIENTS D'ORIGINE VÉGÉTALE ===
		{Nom: "Sachet d'Aconit (Tue-loup)", PrixGallions: 0, PrixMornilles: 8, PrixNoises: 0},
		{Nom: "Armoise en poudre", PrixGallions: 0, PrixMornilles: 5, PrixNoises: 10},
		{Nom: "Racines d'Asphodèle en poudre", PrixGallions: 0, PrixMornilles: 6, PrixNoises: 0},
		{Nom: "Sanguinaire fraîche", PrixGallions: 0, PrixMornilles: 4, PrixNoises: 5},
		{Nom: "Chrysope séchée (la boîte)", PrixGallions: 0, PrixMornilles: 7, PrixNoises: 0},
		{Nom: "Sisymbre cueilli à la pleine lune", PrixGallions: 1, PrixMornilles: 3, PrixNoises: 0},
		{Nom: "Dictame purifié (la fiole)", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Poudre de Pierre de Lune", PrixGallions: 2, PrixMornilles: 10, PrixNoises: 0},

		// === POTIONS PRÉPARÉES & KITS ===
		{Nom: "Kit d'ingrédients de base pour première année", PrixGallions: 2, PrixMornilles: 15, PrixNoises: 0},
		{Nom: "Potion de Soin des Brûlures et Furoncles", PrixGallions: 1, PrixMornilles: 8, PrixNoises: 0},
		{Nom: "Potion d'Aiguisage de la Pensée", PrixGallions: 3, PrixMornilles: 5, PrixNoises: 0},
		{Nom: "Antidote contre les Poisons Communs", PrixGallions: 2, PrixMornilles: 10, PrixNoises: 0},
	}

	for {
		fmt.Println("\n" + Yellow + "=== APOTHICAIRE - INGRÉDIENTS, POTIONS & ASTRONOMIE ===" + Reset)
		fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)
		
		// Affichage aligné des articles
		for i, item := range catalogue {
			fmt.Printf("%2d - %-55s (%2dG %2dM %2dN)\n", i+1, item.Nom, item.PrixGallions, item.PrixMornilles, item.PrixNoises)
		}
		fmt.Println(" Q - Quitter")

		choix := LireSaisie(reader, "Que voulez-vous acheter ? : ")
		if strings.ToLower(strings.TrimSpace(choix)) == "q" {
			ClearTerminal()
			Chemindetraverse()
		}

		var idx int
		_, err := fmt.Sscan(choix, &idx)
		if err != nil || idx < 1 || idx > len(catalogue) {
			continue
		}

		article := catalogue[idx-1]
		if p.Payer(article.PrixGallions, article.PrixMornilles, article.PrixNoises) {
			
			// Tri automatique du type pour l'inventaire
			typeItem := TypeConsommable
			nomLower := strings.ToLower(article.Nom)
			
			if strings.Contains(nomLower, "télescope") || strings.Contains(nomLower, "carte") || strings.Contains(nomLower, "globe") || strings.Contains(nomLower, "astrolabe") {
				typeItem = TypeEquipement
			} else if strings.Contains(nomLower, "chaudron") || strings.Contains(nomLower, "balance") || strings.Contains(nomLower, "fiole") || strings.Contains(nomLower, "kit") {
				typeItem = TypeDivers
			}

			// Le paramètre "PrixValeur" de l'inventaire stocke la valeur en Gallions à titre indicatif
			p.AjouterItem(article.Nom, typeItem, article.PrixGallions, 1)
			
			ClearTerminal()
			fmt.Printf(Green+"Vous avez acheté : %s !\n"+Reset, article.Nom)
			
		} else {
			fmt.Println(Red + "Vous n'avez pas assez d'argent ! Gringotts vous attend." + Reset)
			LancerMagasinapothicaire()
		}
	}
}