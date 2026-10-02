package src

import (
	"fmt"
	"strings"
)
 
func LancerMagasinfarces() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
{Nom: "Bonbons à vomir (Puking Pastilles)", PrixGallions: 0, PrixMornilles: 8, PrixNoises: 12},
{Nom: "Nugat Fainéant (Fainting Fancies)", PrixGallions: 0, PrixMornilles: 10, PrixNoises: 5},
{Nom: "Petits-fours Flessel (Fever Fudge)", PrixGallions: 0, PrixMornilles: 7, PrixNoises: 18},
{Nom: "Pâté au saignement de nez (Nosebleed Nougat)", PrixGallions: 0, PrixMornilles: 9, PrixNoises: 20},

// *=== FEUX D'ARTIFICE & EXPLOSIFS ===*

{Nom: "Feux d'artifice Incendro : Fusée étincelante", PrixGallions: 0, PrixMornilles: 15, PrixNoises: 10},
{Nom: "Feux d'artifice Incendro : Dragon de feu volant", PrixGallions: 1, PrixMornilles: 8, PrixNoises: 15},
{Nom: "Pétard étincelant multicolore", PrixGallions: 0, PrixMornilles: 6, PrixNoises: 20},
{Nom: "Boîte de feux d'artifice autonettoyants", PrixGallions: 2, PrixMornilles: 12, PrixNoises: 10},

// *=== MAGIE DE RUE & DÉFENSE ===*

{Nom: "Poudre d'Obscurité du Pérou", PrixGallions: 0, PrixMornilles: 14, PrixNoises: 25},
{Nom: "Chapeau Bouclier (Shield Hat)", PrixGallions: 1, PrixMornilles: 5, PrixNoises: 10},
{Nom: "Cape Bouclier (Shield Cloak)", PrixGallions: 2, PrixMornilles: 3, PrixNoises: 15},
{Nom: "Gants Bouclier (Shield Gloves)", PrixGallions: 0, PrixMornilles: 13, PrixNoises: 5},
{Nom: "Leurre détonateur (Decoy Detonator)", PrixGallions: 0, PrixMornilles: 11, PrixNoises: 20},
{Nom: "Marécage portatif (Portable Swamp)", PrixGallions: 1, PrixMornilles: 9, PrixNoises: 10},

// *=== FARCES & ATTRAPES ===*

{Nom: "Baguette farceuse (Se transforme en poireau)", PrixGallions: 0, PrixMornilles: 16, PrixNoises: 5},
{Nom: "Baguette farceuse géante", PrixGallions: 0, PrixMornilles: 19, PrixNoises: 15},
{Nom: "Plume à auto-correction (Modèle défectueux)", PrixGallions: 0, PrixMornilles: 8, PrixNoises: 20},
{Nom: "Plume Réponse-Automatique", PrixGallions: 0, PrixMornilles: 14, PrixNoises: 10},
{Nom: "Plume à Encre Incolore", PrixGallions: 0, PrixMornilles: 7, PrixNoises: 25},
{Nom: "Télescope frappeur", PrixGallions: 0, PrixMornilles: 15, PrixNoises: 5},
{Nom: "Corde à nœud coulissant magique", PrixGallions: 0, PrixMornilles: 12, PrixNoises: 15},
{Nom: "Gobelets à fond amovible", PrixGallions: 0, PrixMornilles: 5, PrixNoises: 20},
{Nom: "Savon de crapaud", PrixGallions: 0, PrixMornilles: 4, PrixNoises: 15},

// *=== FRIANDISES PIÉGÉES ===*

{Nom: "Pralines d'Amour (Love Potions)", PrixGallions: 0, PrixMornilles: 17, PrixNoises: 10},
{Nom: "Crème de Canari", PrixGallions: 0, PrixMornilles: 13, PrixNoises: 20},
{Nom: "Savons en sucre", PrixGallions: 0, PrixMornilles: 5, PrixNoises: 15},
{Nom: "Diablotins en poivre", PrixGallions: 0, PrixMornilles: 8, PrixNoises: 25},

// *=== RÊVES ÉVEILLÉS ===*

{Nom: "Sortilège de Rêve Éveillé (Daydream Charms)", PrixGallions: 0, PrixMornilles: 16, PrixNoises: 20},
{Nom: "Poudre de Camouflage / Teint Parfait", PrixGallions: 0, PrixMornilles: 10, PrixNoises: 15},
{Nom: "Anti-Boutons Garanti", PrixGallions: 0, PrixMornilles: 6, PrixNoises: 10},
{Nom: "Boursouflet Rose (Pygmy Puff)", PrixGallions: 1, PrixMornilles: 2, PrixNoises: 20},
	}

	for {
		fmt.Println("\n" + Yellow + "=== WEASLEY & WEASLEY - FARCES ===" + Reset)
		fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)
		for i, item := range catalogue {
			fmt.Printf("%2d - %-60s (%2dG %2dM %2dN)\n", i+1, item.Nom, item.PrixGallions, item.PrixMornilles, item.PrixNoises)
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
			p.AjouterItem(article.Nom, TypeDivers, article.PrixGallions, 1)
			ClearTerminal()
			fmt.Printf(Green+"Vous avez acheté : %s !\n"+Reset, article.Nom)
		} else {
			fmt.Println(Red + "Vous n'avez pas assez d'argent !" + Reset)
		}
	}
}