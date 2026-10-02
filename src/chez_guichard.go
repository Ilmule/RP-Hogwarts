package src

import (
	"fmt"
	"strings"
)

func LancerMagasinGuichard() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
		// Uniformes & Essentiels
		{Nom: "Robe de travail (noire)", PrixGallions: 5, PrixMornilles: 20, PrixNoises: 0},
		{Nom: "Chapeau pointu (noir)", PrixGallions: 1, PrixMornilles: 10, PrixNoises: 0},
		{Nom: "Cape d'hiver avec attaches d'argent", PrixGallions: 7, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Gants protecteurs en cuir de dragon", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		// Accessoires & Cérémonie
		{Nom: "Robe de cérémonie en velours prune", PrixGallions: 8, PrixMornilles: 7, PrixNoises: 0},
		{Nom: "Robe de sorcier brodée de fils d'or", PrixGallions: 150, PrixMornilles: 20, PrixNoises: 0},
		{Nom: "Gants en peau de Moke", PrixGallions: 8, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Écharpe aux couleurs de sa maison", PrixGallions: 2, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Ceinture en cuir de Grip-sec", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Cravate enchantée serrée automatique", PrixGallions: 7, PrixMornilles: 12, PrixNoises: 0},
		// Chapeaux & Look
		{Nom: "Chapeau pointu à plumes de Flupbert", PrixGallions: 4, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chapeau haut-de-forme à pigeon mécanique", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Gilet en laine tricoté main (style Weasley)", PrixGallions: 2, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Robe à motifs changeants selon l'humeur", PrixGallions: 7, PrixMornilles: 0, PrixNoises: 0},
		// Inutiles & Gadgets magiques
		{Nom: "Chaussettes dépareillées hurlantes", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Bonnet à oreilles de Morsag", PrixGallions: 2, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Manteau invisible (défectueux)", PrixGallions: 8, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Cape de lévitation ratée", PrixGallions: 6, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Bottes à ressorts pour lutins", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
	}

	for {
		fmt.Println("\n" + Yellow + "=== CHEZ GUICHARD - VÊTEMENTS ===" + Reset)
		fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)
		for i, item := range catalogue {
			fmt.Printf("%2d - %-50s (%2dG %2dM %2dN)\n", i+1, item.Nom, item.PrixGallions, item.PrixMornilles, item.PrixNoises)
		}
		fmt.Println(" Q - Quitter")

		choix := LireSaisie(reader, "Que voulez-vous acheter ? : ")
		if strings.ToLower(strings.TrimSpace(choix)) == "q" {
			ClearTerminal()
			Chemindetraverse()
			break
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