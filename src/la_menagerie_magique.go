package src

import (
	"fmt"
	"strings"
)

func LancerMagasinmenagerie() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
		// === ANIMAUX AUTORISÉS ===
		{Nom: "Hibou grand-duc (utile pour le courrier)", PrixGallions: 10, PrixMornilles: 10, PrixNoises: 0},
		{Nom: "Chouette effraie", PrixGallions: 8, PrixMornilles: 20, PrixNoises: 0},
		{Nom: "Petit duc roux", PrixGallions: 6, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chat de gouttière roux", PrixGallions: 9, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chat à poil long (pense être un Flairousteur)", PrixGallions: 9, PrixMornilles: 7, PrixNoises: 0},
		{Nom: "Crapaud vert tacheté (très lent)", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Rat brun dodu (semble vivre anormalement longtemps)", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		// === ANIMAUX INUTILES ===
		{Nom: "Escargot orange géant", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Crapaud pustuleux géant", PrixGallions: 2, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Tortue à carapace étincelante", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Bousier magique de collection", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Sangsue de compagnie", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Poisson-lune miniature en bocal", PrixGallions: 2, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chenille multicolore", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Chauve-souris naine", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		// === CRÉATURES D'AGRÉMENT ===
		{Nom: "Boursouflet Rose (Pygmy Puff)", PrixGallions: 5, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Flairousteur (Kneazle)", PrixGallions: 12, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Crabe de Feu", PrixGallions: 15, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Rat noir à queue articulée", PrixGallions: 4, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Oiseau Gobeur de mouches magique", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
		// === ACCESSOIRES ===
		{Nom: "Boîte de friandises pour hiboux et chouettes", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Rats en sucre pour chats", PrixGallions: 1, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Cage dorée pour hibou", PrixGallions: 4, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Brosse à poils de Flanker pour chats", PrixGallions: 2, PrixMornilles: 0, PrixNoises: 0},
		{Nom: "Laisse d'invisibilité pour créatures volantes", PrixGallions: 3, PrixMornilles: 0, PrixNoises: 0},
	}

	for {
		fmt.Println("\n" + Yellow + "=== LA MÉNAGERIE MAGIQUE ===" + Reset)
		fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)
		for i, item := range catalogue {
			fmt.Printf("%2d - %-55s (%2dG %2dM %2dN)\n", i+1, item.Nom, item.PrixGallions, item.PrixMornilles, item.PrixNoises)
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
			nomLower := strings.ToLower(article.Nom)
			estAccessoire := strings.Contains(nomLower, "boîte") ||
				strings.Contains(nomLower, "rats en sucre") ||
				strings.Contains(nomLower, "cage") ||
				strings.Contains(nomLower, "brosse") ||
				strings.Contains(nomLower, "laisse")
 
			if !estAccessoire {
				nomPersonnalise := LireSaisie(reader, "Quel nom souhaitez-vous donner à votre nouvel animal ? : ")
				nomPersonnalise = strings.TrimSpace(nomPersonnalise)

				nomFinal := article.Nom
				if nomPersonnalise != "" {
					nomFinal = fmt.Sprintf("%s (%s)", article.Nom, nomPersonnalise)
				}
				p.AjouterItem(nomFinal, TypeAnimal, article.PrixGallions, 1)
				ClearTerminal()
				fmt.Printf(Green+"Vous avez acheté votre compagnon : %s !\n"+Reset, nomFinal)
			} else {
				p.AjouterItem(article.Nom, TypeDivers, article.PrixGallions, 1)
				ClearTerminal()
				fmt.Printf(Green+"Vous avez acheté : %s !\n"+Reset, article.Nom)
			}
		} else {
			fmt.Println(Red + "Vous n'avez pas assez d'argent !" + Reset)
		}
	}
}