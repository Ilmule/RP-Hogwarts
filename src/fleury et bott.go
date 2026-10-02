package src

import (
	"fmt"
	"strings"
)
 
func LancerMagasinfleury() {
	p := JoueurActuel
	reader := Lecteur

	catalogue := []ItemBoutique{
// *=== MANUELS OBLIGATOIRES ===*

{Nom: "Livre : Le Livre des sorts et enchantements (Niveau 1) - Miranda Fauconnette", PrixGallions: 0, PrixMornilles: 12, PrixNoises: 15},
{Nom: "Livre : Histoire de la magie - Bathilda Tourdesac", PrixGallions: 0, PrixMornilles: 16, PrixNoises: 10},
{Nom: "Livre : Magie théorique - Adalbert Lasornette", PrixGallions: 0, PrixMornilles: 11, PrixNoises: 20},
{Nom: "Livre : Manuel de métamorphose à l'usage des débutants - Emeric G. Changé", PrixGallions: 0, PrixMornilles: 13, PrixNoises: 5},
{Nom: "Livre : Mille herbes et champignons magiques - Phyllida Augurolle", PrixGallions: 0, PrixMornilles: 14, PrixNoises: 15},
{Nom: "Livre : Potions magiques - Arsenius Beaulitron", PrixGallions: 0, PrixMornilles: 18, PrixNoises: 10},
{Nom: "Livre : Vie et habitat des animaux fantastiques - Norbert Dragonneau", PrixGallions: 1, PrixMornilles: 2, PrixNoises: 15},
{Nom: "Livre : Forces obscures : comment s'en protéger - Quentin Jentremble", PrixGallions: 1, PrixMornilles: 11, PrixNoises: 20},

// *=== NIVEAUX SUPÉRIEURS ===*

{Nom: "Livre : Le Livre des sorts et enchantements (Niveaux 2 à 7)", PrixGallions: 1, PrixMornilles: 8, PrixNoises: 15},
{Nom: "Livre : Le Monstrueux Livre des Monstres (Attention aux doigts !)", PrixGallions: 2, PrixMornilles: 5, PrixNoises: 20},
{Nom: "Livre : Le Levé du voile sur l'avenir - Cassandra Vablatsky", PrixGallions: 1, PrixMornilles: 12, PrixNoises: 10},
{Nom: "Livre : Manuel d'apprentissage de la métamorphose intermédiaire", PrixGallions: 1, PrixMornilles: 16, PrixNoises: 5},
{Nom: "Livre : Potions avancées - Libatius Borage", PrixGallions: 2, PrixMornilles: 14, PrixNoises: 20},
{Nom: "Livre : Élaboration de sortilèges avancés", PrixGallions: 2, PrixMornilles: 7, PrixNoises: 15},

// *=== GILDEROY LOCKHART ===*

{Nom: "Livre : Flâneries avec les goules - Gilderoy Lockhart", PrixGallions: 1, PrixMornilles: 5, PrixNoises: 10},
{Nom: "Livre : Promenades avec les loups-garous - Gilderoy Lockhart", PrixGallions: 1, PrixMornilles: 7, PrixNoises: 5},
{Nom: "Livre : Randonnées avec les banshees - Gilderoy Lockhart", PrixGallions: 1, PrixMornilles: 6, PrixNoises: 15},
{Nom: "Livre : Une année avec le Yéti - Gilderoy Lockhart", PrixGallions: 1, PrixMornilles: 8, PrixNoises: 10},
{Nom: "Livre : Mon autobiographie : Moi le Magicien - Gilderoy Lockhart", PrixGallions: 2, PrixMornilles: 3, PrixNoises: 20},

// *=== LORE & DIVERS ===*

{Nom: "Livre : Le Quidditch à travers les âges - Kennilworthy Whisp", PrixGallions: 1, PrixMornilles: 3, PrixNoises: 15},
{Nom: "Livre : Les Contes de Beedle le Bard (Édition originale en runes)", PrixGallions: 1, PrixMornilles: 15, PrixNoises: 20},
{Nom: "Livre : Présages de mort : que faire quand vous sentez que le pire arrive", PrixGallions: 1, PrixMornilles: 4, PrixNoises: 10},
{Nom: "Livre : Les Arbres généalogiques des familles de Sang-Pur", PrixGallions: 3, PrixMornilles: 2, PrixNoises: 15},
{Nom: "Livre : Guide du jeune sorcier pour l'auto-défense rapide", PrixGallions: 1, PrixMornilles: 5, PrixNoises: 10},
{Nom: "Livre : Comment ensorceler vos balais et vos voisins", PrixGallions: 1, PrixMornilles: 6, PrixNoises: 20},
{Nom: "Livre : Histoire des Poussées de Boutons et de la Peste du Dragon", PrixGallions: 0, PrixMornilles: 15, PrixNoises: 5},
{Nom: "Livre : Les Sorts invisibles et la Magie d'illusion", PrixGallions: 2, PrixMornilles: 1, PrixNoises: 10},
{Nom: "Livre : Dépasser les limites de la Métamorphose humaine", PrixGallions: 2, PrixMornilles: 12, PrixNoises: 15},
{Nom: "Livre : Grand livre des malédictions et contre-malédictions", PrixGallions: 2, PrixMornilles: 4, PrixNoises: 20},
{Nom: "Livre : L'Art des Potions et des Poisons invisibles", PrixGallions: 2, PrixMornilles: 13, PrixNoises: 10},
	}

	for {
		fmt.Println("\n" + Yellow + "=== FLEURY ET BOTT - LIBRAIRIE ===" + Reset)
		fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)
		for i, item := range catalogue {
			fmt.Printf("%2d - %-75s (%2dG %2dM %2dN)\n", i+1, item.Nom, item.PrixGallions, item.PrixMornilles, item.PrixNoises)
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
			p.AjouterItem(article.Nom, TypeLivre, article.PrixGallions, 1)
			ClearTerminal()
			fmt.Printf(Green+"Vous avez acheté : %s !\n"+Reset, article.Nom)
		} else {
			fmt.Println(Red + "Vous n'avez pas assez d'argent !" + Reset)
		}
	}
}