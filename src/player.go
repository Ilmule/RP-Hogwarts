package src

import (
	"fmt"
)

type Player struct {
	Name            	string
	Maison          	string
	Animal          	string
	Nomanimal       	string
	Level           	int
	Levelbotanique		int
	Leveldivination		int
	Levelmetamorphose	int
	Levelpotions		int
	Levelsortilege		int
	Health          	int
	MaxHealth       	int
	Atk             	int
	Def             	int
	Gallions        	int
	Mornilles       	int
	Noises          	int
	CoffreGallions  	int
	CoffreMornilles 	int
	CoffreNoises    	int
	Exp             	int
	Expdivination       int
	Expmetamorphose     int
	Expbotanique        int
	Exppotions          int
	Expsortilege       	int
	Title           	string

	// Champs d'inventaire et d'équipement
	Inventaire      []Item
	BaguetteEquipee string
	BalaiEquipe     string
	RobeEquipee     string
}

func CreatePlayerFromInput() Player {
	reader := Lecteur

	fmt.Println("=== CRÉATION DU PERSONNAGE ===")

	nom := LireSaisie(reader, "Entrez votre nom : ")

	p := Player{
		Name:            nom,
		Maison:         	 "Non attribuée",
		// Animal:         	 "Hibou",
		// Nomanimal:      	 "Hedwige",
		Level:          	 1,
		Levelbotanique:		 1,
		Leveldivination:	 1,
		Levelmetamorphose:	 1,
		Levelpotions:		 1,
		Levelsortilege:		 1,
		Health:         	 100,
		MaxHealth:      	 100,
		Atk:            	 10,
		Def:            	 5,
		Gallions:       	 80,
		Mornilles:      	 10,
		Noises:         	 10,
		Exp:            	 0,
		Title:          	 "Première année",
		Inventaire:     	 []Item{},
		BaguetteEquipee:	 "Baguette de test",
		RobeEquipee:    	 "Robe de travail (noire)",
	}

	// // --- FOURNITURES OBLIGATOIRES ET TICKET DE DÉPART POUR VOS TESTS ---

	// // 1. Uniformes & Équipements
	// p.AjouterItem("Robe de travail (noire)", TypeEquipement, 3, 3)
	// p.AjouterItem("Chapeau pointu (noir)", TypeEquipement, 1, 1)
	// p.AjouterItem("Gants protecteurs en cuir de dragon", TypeEquipement, 2, 1)
	// p.AjouterItem("Cape d'hiver avec attaches d'argent", TypeEquipement, 5, 1)

	// // 2. Matériel de classe
	// p.AjouterItem("Baguette en houx, 11 pouces, plume de phénix", TypeEquipement, 7, 1)
	// p.AjouterItem("Chaudron en étain (taille 2)", TypeDivers, 5, 1)
	// p.AjouterItem("Boîte de fioles en verre", TypeDivers, 2, 1)
	// p.AjouterItem("Balance en cuivre", TypeDivers, 3, 1)

	// // 3. Manuels scolaires de 1ère année
	// p.AjouterItem("Livre : Le Livre des sorts et enchantements (Niveau 1) - Miranda Fauconnette", TypeLivre, 1, 1)
	// p.AjouterItem("Livre : Histoire de la magie - Bathilda Tourdesac", TypeLivre, 2, 1)
	// p.AjouterItem("Livre : Magie théorique - Adalbert Lasornette", TypeLivre, 1, 1)
	// p.AjouterItem("Livre : Manuel de métamorphose à l'usage des débutants - Emeric G. Changé", TypeLivre, 1, 1)
	// p.AjouterItem("Livre : Mille herbes et champignons magiques - Phyllida Augurolle", TypeLivre, 1, 1)
	// p.AjouterItem("Livre : Potions magiques - Arsenius Beaulitron", TypeLivre, 2, 1)
	// p.AjouterItem("Livre : Vie et habitat des animaux fantastiques - Norbert Dragonneau", TypeLivre, 2, 1)
	// p.AjouterItem("Livre : Forces obscures : comment s'en protéger - Quentin Jentremble", TypeLivre, 3, 1)

	// // 4. Compagnon (Animal par défaut)
	// p.AjouterItem("Hibou (Hedwige)", TypeAnimal, 5, 1)

	// // 5. Billet pour la voie 9 ¾
	// p.AjouterItem("Ticket pour le Poudlard Express (Voie 9 ¾)", TypeDivers, 0, 1)

	JoueurActuel = &p

	return p
}

// AjouterItem ajoute un objet dans le sac du joueur ou augmente sa quantité si déjà présent
func (p *Player) AjouterItem(nom string, typeItem TypeItem, prix int, quantite int) {
	for i := range p.Inventaire {
		if p.Inventaire[i].Nom == nom {
			p.Inventaire[i].Quantite += quantite
			return
		}
	}

	p.Inventaire = append(p.Inventaire, Item{
		Nom:        nom,
		Type:       typeItem,
		PrixValeur: prix,
		Quantite:   quantite,
	})
}

func AfficherBourse(p Player) {
	fmt.Println(Yellow + "=== BOURSE DE GAINS (Gringotts) ===" + Reset)
	fmt.Printf("🪙 Gallions (Or)  : %d\n", p.Gallions)
	fmt.Printf("🥈 Mornilles (Argent): %d\n", p.Mornilles)
	fmt.Printf("🥉 Noises (Bronze)  : %d\n", p.Noises)
}

// Payer vérifie si le joueur a assez d'argent au total et fait l'appoint automatiquement
func (p *Player) Payer(prixG int, prixM int, prixN int) bool {
	// 1 Gallion = 17 Mornilles, 1 Mornille = 29 Noises -> 1 Gallion = 493 Noises
	totalNoisesJoueur := (p.Gallions * 493) + (p.Mornilles * 29) + p.Noises
	totalNoisesCout := (prixG * 493) + (prixM * 29) + prixN

	// Si le joueur est trop pauvre au total, on refuse l'achat
	if totalNoisesJoueur < totalNoisesCout {
		return false
	}

	// On soustrait le coût global au portefeuille du joueur
	reste := totalNoisesJoueur - totalNoisesCout

	// On recalcule proprement les pièces (le "rendu de monnaie")
	p.Gallions = reste / 493
	reste = reste % 493
	p.Mornilles = reste / 29
	p.Noises = reste % 29

	return true
}