package src

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// BainsDesPrefets gère l'exploration et les activités de la salle de bain des préfets
func Soussolsprefets() {
	p := JoueurActuel
	reader := Lecteur

	// Initialisation du générateur aléatoire pour le mini-jeu
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Println(Cyan + "               LA SALLE DE BAIN DES PRÉFETS               " + Reset)
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("La pièce est entièrement en marbre blanc, baignée d'une lumière douce.")
		fmt.Println("Au centre se trouve un bassin rectangulaire de la taille d'une piscine,")
		fmt.Println("entouré d'une centaine de robinets en or incrustés de pierres précieuses.")
		fmt.Println("Un magnifique vitrail représentant une sirène endormie orne le mur.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Faire couler un bain relaxant (Restaurer tous ses PV)")
		fmt.Println(" 2 - Jouer à 'L'Attrape-Bulle Magique' (Mini-Jeu)")
		fmt.Println(" 3 - Sortir et retourner au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			PrendreBain()
		case "2":
			ClearTerminal()
			MiniJeuBulles()
		case "3":
			ClearTerminal()
			fmt.Println(Yellow + "Vous remettez votre cape et retournez dans les couloirs..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard() // Remplace par le menu d'où tu accèdes aux bains si besoin
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Soussolsprefets) // Retourne ici après fermeture de l'inventaire[cite: 1]
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// PrendreBain soigne complètement le joueur
func PrendreBain() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Cyan + "=== BAIN RELAXANT ===" + Reset)
	fmt.Println("Vous ouvrez plusieurs robinets en or au hasard. Des jets d'eau chaude parfumée,")
	fmt.Println("des nuages de mousse épaisse et des bulles multicolores remplissent le bassin.")
	time.Sleep(2 * time.Second)
	
	fmt.Println("\nVous plongez dans l'eau merveilleusement chaude... C'est incroyablement relaxant.")
	fmt.Println("La sirène du vitrail ouvre un œil, vous fait un clin d'œil, puis se rendort.")
	time.Sleep(2 * time.Second)

	p.Health = p.MaxHealth
	fmt.Println(Green + "\nToutes vos courbatures disparaissent. Vous êtes en pleine forme ! (PV restaurés au maximum)" + Reset)
	
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour sortir de l'eau et vous rhabiller]"+Reset)
}

// MiniJeuBulles est un jeu de hasard et de réflexe avec des récompenses (ou des pièges)
func MiniJeuBulles() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Cyan + "=== JEU : L'ATTRAPE-BULLE MAGIQUE ===" + Reset)
	fmt.Println("Vous tournez un robinet étrange en forme de gargouille. Au lieu d'eau,")
	fmt.Println("il se met à cracher des bulles magiques géantes qui s'envolent rapidement vers le plafond !")
	fmt.Println("Vous n'avez le temps d'en attraper qu'une seule avant qu'elles n'éclatent.")
	time.Sleep(2 * time.Second)

	// Liste des bulles possibles
	bullesPossibles := []string{
		"Bulle Émeraude", 
		"Bulle Saphir", 
		"Bulle Rubis", 
		"Bulle Dorée", 
		"Bulle Noire d'Encre",
	}

	// Mélange des bulles
	rand.Shuffle(len(bullesPossibles), func(i, j int) {
		bullesPossibles[i], bullesPossibles[j] = bullesPossibles[j], bullesPossibles[i]
	})

	// On ne propose que les 3 premières pour laisser planer le mystère
	bullesProposees := bullesPossibles[:3]

	fmt.Println("\nTrois bulles s'envolent devant vous :")
	for i, b := range bullesProposees {
		fmt.Printf(" %d - Attraper la %s\n", i+1, b)
	}
	fmt.Println(" R - Ne rien attraper et reculer prudemment")

	choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nLaquelle attrapez-vous ? (1-3) : ")))

	if choix == "R" {
		fmt.Println("\nVous regardez les bulles éclater au plafond dans un petit *pop* inoffensif.")
		time.Sleep(1 * time.Second)
		return
	}

	idx, err := strconv.Atoi(choix)
	if err != nil || idx < 1 || idx > 3 {
		fmt.Println(Red + "\nVous avez hésité trop longtemps... Les bulles ont éclaté au plafond !" + Reset)
		time.Sleep(2 * time.Second)
		return
	}

	// Résolution
	bulleChoisie := bullesProposees[idx-1]
	fmt.Printf(Yellow+"\nVous sautez et attrapez la %s au vol ! *POP*\n"+Reset, bulleChoisie)
	time.Sleep(1 * time.Second)

	switch bulleChoisie {
	case "Bulle Émeraude":
		gain := 15
		p.Exp += gain
		fmt.Printf(Green+"Un nuage de fumée parfumée vous enveloppe. Votre esprit s'éclaircit (+%d Exp) !\n"+Reset, gain)
	
	case "Bulle Saphir":
		soin := 30
		p.Health += soin
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
		fmt.Printf(Cyan+"Une eau cristalline glacée vous éclabousse. Vous êtes vivifié (+%d PV) !\n"+Reset, soin)
	
	case "Bulle Rubis":
		degats := 10
		p.Health -= degats
		if p.Health < 1 {
			p.Health = 1 // On ne meurt pas en jouant avec des bulles
		}
		fmt.Printf(Red+"Aïe ! La bulle explose comme un pétard brûlant ! Vous perdez %d PV.\n"+Reset, degats)
	
	case "Bulle Dorée":
		gain := rand.Intn(4) + 2
		p.Gallions += gain
		fmt.Printf(Yellow+"Jackpot ! Une pluie de %d Gallions tombe de la bulle ! Vous les ramassez vite.\n"+Reset, gain)
	
	case "Bulle Noire d'Encre":
		fmt.Println(DarkGray + "SPOUTCH ! La bulle éclate et vous recouvre le visage d'encre de seiche puante..." + Reset)
		fmt.Println("La sirène du vitrail se met à rire. Vous êtes bon pour un autre bain.")
	}

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)
}