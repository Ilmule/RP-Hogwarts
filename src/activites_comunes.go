package src

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// JeuEchecsSorcier est un mini-jeu tactique fonctionnel en 3 tours
func JeuEchecsSorcier() {
	reader := Lecteur
	fmt.Println(Yellow + "=== ÉCHECS VERSION SORCIER ===" + Reset)
	fmt.Println("Vous jouez contre un camarade. Les pièces s'animent et s'insultent sur le plateau !")

	scoreJoueur := 15
	scoreAdversaire := 15
	rand.Seed(time.Now().UnixNano())

	for tour := 1; tour <= 3; tour++ {
		fmt.Printf("\n--- TOUR %d ---\n", tour)
		fmt.Printf("Votre Roi : %d PV | Roi adverse : %d PV\n", scoreJoueur, scoreAdversaire)
		fmt.Println(" 1 - Attaque brutale (Cavalier et Tour)")
		fmt.Println(" 2 - Défense stricte (Pions en phalange)")
		fmt.Println(" 3 - Piège stratégique (Sacrifice du Fou)")

		choix := strings.TrimSpace(LireSaisie(reader, "Votre stratégie (1-3) : "))
		actionAdverse := rand.Intn(3) + 1 // 1: Attaque, 2: Défense, 3: Piège

		fmt.Println(Cyan + "L'adversaire déplace ses pièces..." + Reset)
		time.Sleep(1 * time.Second)

		// Résolution (Pierre-Feuille-Ciseaux)
		if choix == "1" {
			if actionAdverse == 1 {
				fmt.Println("Choc frontal ! Les pièces se fracassent.")
				scoreJoueur -= 3
				scoreAdversaire -= 3
			} else if actionAdverse == 2 {
				fmt.Println("Votre attaque rebondit sur sa défense, votre Cavalier est détruit !")
				scoreJoueur -= 4
			} else {
				fmt.Println("Vous percez son piège et détruisez sa Tour !")
				scoreAdversaire -= 5
			}
		} else if choix == "2" {
			if actionAdverse == 1 {
				fmt.Println("Sa violente attaque est bloquée, vous contre-attaquez !")
				scoreAdversaire -= 4
			} else if actionAdverse == 2 {
				fmt.Println("Vous défendez tous les deux. Rien ne se passe.")
			} else {
				fmt.Println("Votre défense statique vous rend vulnérable à son piège !")
				scoreJoueur -= 5
			}
		} else if choix == "3" {
			if actionAdverse == 1 {
				fmt.Println("Votre piège est trop lent face à son assaut brutal !")
				scoreJoueur -= 5
			} else if actionAdverse == 2 {
				fmt.Println("Vous contournez sa défense grâce à votre sacrifice et prenez sa Reine !")
				scoreAdversaire -= 5
			} else {
				fmt.Println("Vous tendez un piège tous les deux. Le plateau est confus.")
			}
		} else {
			fmt.Println(Red + "Stratégie incomprise. Vous passez votre tour et perdez des pièces." + Reset)
			scoreJoueur -= 5
		}
	}

	fmt.Println("\n" + Yellow + "=== FIN DE LA PARTIE ===" + Reset)
	if scoreJoueur > scoreAdversaire {
		fmt.Println(Green + "Échec et Mat ! Vous fracassez le Roi adverse. Belle victoire (+10 Exp) !" + Reset)
		JoueurActuel.Exp += 10
	} else if scoreJoueur < scoreAdversaire {
		fmt.Println(Red + "Votre Roi jette sa couronne à terre et se rend. Vous avez perdu." + Reset)
	} else {
		fmt.Println("Égalité parfaite. Il ne reste plus que les deux Rois sur le plateau.")
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour quitter le plateau]"+Reset)
}

// ChangerVetementsDortoir filtre l'inventaire pour ne proposer que des équipements
func ChangerVetementsDortoir() {
	p := JoueurActuel
	reader := Lecteur

	type Vetement struct {
		Nom   string
		Index int
	}

	var penderie []Vetement
	for i, item := range p.Inventaire {
		if item.Type == TypeEquipement {
			penderie = append(penderie, Vetement{Nom: item.Nom, Index: i})
		}
	}

	fmt.Println("\n" + Yellow + "=== VOTRE MALLE (PENDERIE) ===" + Reset)
	fmt.Println("Tenue actuelle : " + Cyan + p.RobeEquipee + Reset)

	if len(penderie) == 0 {
		fmt.Println(Red + "Vous n'avez aucun vêtement ou équipement de rechange dans votre malle." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)
		return
	}

	fmt.Println("Que voulez-vous enfiler ?")
	for i, v := range penderie {
		fmt.Printf(" %d - %s\n", i+1, v.Nom)
	}
	fmt.Println(" Q - Ne rien changer")

	choix := LireSaisie(reader, "\nVotre choix : ")
	if strings.ToLower(strings.TrimSpace(choix)) == "q" {
		return
	}

	idx, err := strconv.Atoi(strings.TrimSpace(choix))
	if err == nil && idx >= 1 && idx <= len(penderie) {
		vetementChoisi := penderie[idx-1]
		p.RobeEquipee = vetementChoisi.Nom
		fmt.Println(Green + "Vous avez revêtu : " + p.RobeEquipee + Reset)
	} else {
		fmt.Println(Red + "Choix invalide." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)
}