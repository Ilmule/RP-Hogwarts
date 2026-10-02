package src

import (
	"fmt"
	"strings"
)

func LancerMagasinfortescue() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
		{Nom: "Glace simple au beurre de cacahuète", PrixGallions: 1, PrixMornilles: 10, PrixNoises: 0},
		{Nom: "Glace à la vanille et jus de potiron", PrixGallions: 1, PrixMornilles: 10, PrixNoises: 8},
		{Nom: "Coupe géante Fortescue (+30 PV)", PrixGallions: 4, PrixMornilles: 20, PrixNoises: 9},
		{Nom: "Sorbet Poudlard aux fruits magiques (+15 PV)", PrixGallions: 0, PrixMornilles: 8, PrixNoises: 15},
		{Nom: "Glace à la Bièrobeurre frappée", PrixGallions: 0, PrixMornilles: 10, PrixNoises: 0},
		{Nom: "Jus de Potiron glacé", PrixGallions: 0, PrixMornilles: 4, PrixNoises: 20},
	}

	for {
		fmt.Println("\n" + Yellow + "=== GLACES FORTESCUE ===" + Reset)
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
			p.AjouterItem(article.Nom, TypeConsommable, article.PrixGallions, 1)
			ClearTerminal()
			fmt.Printf(Green+"Miam ! Vous avez dégusté : %s !\n"+Reset, article.Nom)
		} else {
			fmt.Println(Red + "Vous n'avez pas assez d'argent !" + Reset)
		}
	}
}