package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// CoursPotions gère l'accès au cours de potions avec le Professeur Rogue
func Courspotions() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		// Calcul de l'année en fonction du niveau (1-10 = 1ère année, 11-20 = 2ème année, etc.)
		annee := (p.Levelpotions-1)/10 + 1
		if annee > 7 {
			annee = 7
		}

		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Println(Green + "               SALLE DE CLASSE DES POTIONS                " + Reset)
		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp potions : %d | Gallions : %d | Niveau : %d (Année %d)\n", p.Health, p.MaxHealth, p.Exppotions, p.Gallions, p.Levelpotions, annee)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("L'atmosphère dans les cachots est glaciale. Des chaudrons fument")
		fmt.Println("lentement sur les paillasses en pierre. Le Professeur Rogue se tient")
		fmt.Println("silencieusement dans l'ombre, observant chaque mouvement avec mépris.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Écouter les avertissements du Professeur Rogue (Dialogue)")
		fmt.Println(" 2 - 1ère Année : Potion de Soin des Furoncles (Niveau 1+)")
		fmt.Println(" 3 - 2ème Année : Philtre d'Apaisement (Niveau 11+)")
		fmt.Println(" 4 - 3ème Année : Potion d'Aiguisage de la Pensée (Niveau 21+)")
		fmt.Println(" 5 - 5ème Année : Potion de Polynectar (Niveau 41+) [Très complexe]")
		fmt.Println(" 6 - Retourner au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Green + "==========================================================" + Reset)

		Option := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch Option {
		case "1":
			ClearTerminal()
			DialoguesRoguePotions()
		case "2":
			ClearTerminal()
			TenterPotionFuroncles(p)
		case "3":
			ClearTerminal()
			TenterPhiltreApaisement(p)
		case "4":
			ClearTerminal()
			TenterPotionAiguisage(p)
		case "5":
			ClearTerminal()
			TenterPotionPolynectar(p)
		case "6":
			ClearTerminal()
			fmt.Println(Yellow + "Vous saluez Rogue en évitant son regard et quittez les cachots..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard()
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Courspotions)
			return
		default:
			ClearTerminal()
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// DialoguesRoguePotions gère les remarques sarcastiques de Rogue
func DialoguesRoguePotions() {
	reader := Lecteur
	dialogues := []string{
		"« Les potions exigent une précision absolue. Un mauvais coup de baguette, un gramme de trop, et c'est la catastrophe... »",
		"« Inutile deagiter vos chaudrons avec autant d'agitation. La bave de crapaud ne s'apprivoise pas par la force. »",
		"« Je ne tolérerai aucune légèreté dans mon cachot. Moins encore de la part des incompétents notoires. »",
	}
	phrase := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Green + "=== PROFESSEUR SEVERUS ROGUE ===" + Reset)
	fmt.Println("Rogue s'approche de votre table à pas feutrés et croise les bras.")
	fmt.Println()
	fmt.Printf(Cyan+"Rogue : %s\n"+Reset, phrase)
	fmt.Println()

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour reprendre le travail]"+Reset)
}

// --- FONCTION DE VÉRIFICATION DES INGRÉDIENTS AVEC COÛT POUR L'APOTHICAIRE ---

func VerifierEtProposerAchat(p *Player, nomIngredientReq string, coutGallions int) bool {
	reader := Lecteur

	// Vérifie si le joueur possède l'ingrédient dans son inventaire
	possedeIngredient := false
	for _, item := range p.Inventaire {
		if strings.Contains(strings.ToLower(item.Nom), strings.ToLower(nomIngredientReq)) {
			possedeIngredient = true
			break
		}
	}

	if possedeIngredient {
		return true // Le joueur a l'ingrédient, on peut continuer
	}

	// S'il ne l'a pas, on lui indique le montant nécessaire pour aller l'acheter chez l'Apothicaire
	fmt.Println(Red + "=== INGRÉDIENTS MANQUANTS ===" + Reset)
	fmt.Printf("Il vous manque l'ingrédient essentiel : '%s' !\n", nomIngredientReq)
	fmt.Printf("Pour réaliser cette potion, rendez-vous d'abord chez l'Apothicaire au Chemin de Traverse.\n")
	fmt.Printf("Il vous faut au moins %d Gallions pour l'acheter[cite: 3, 5].\n", coutGallions)
	fmt.Printf("Vous possédez actuellement : %d Gallions.\n\n", p.Gallions)

	if p.Gallions >= coutGallions {
		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "Voulez-vous dépenser instantanément "+fmt.Sprint(coutGallions)+" Gallions pour acheter l'ingrédient et lancer le cours ? (O/N) : ")))
		if choix == "O" {
			p.Gallions -= coutGallions
			p.AjouterItem(nomIngredientReq, TypeConsommable, coutGallions, 1)
			fmt.Println(Green + "Achat rapide effectué ! Vous disposez désormais de l'ingrédient." + Reset)
			time.Sleep(1 * time.Second)
			return true
		}
	} else {
		fmt.Println(Red + "Vous n'avez pas assez de Gallions pour acheter cet ingrédient. Allez faire un tour à Gringotts !" + Reset)
	}

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour revenir au choix des potions]"+Reset)
	return false
}

// --- NIVEAU 1 : POTION DE SOIN DES FURONCLES ---

func TenterPotionFuroncles(p *Player) {
	reader := Lecteur
	if p.Levelpotions < 1 {
		fmt.Println(Red + "Niveau insuffisant (Niveau 1 requis)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Sangsues fraîches (Coût chez l'Apothicaire : 2 Gallions)
	if !VerifierEtProposerAchat(p, "Sangsues fraîches", 2) {
		return
	}

	fmt.Println(Cyan + "=== PRÉPARATION : POTION DE SOIN DES FURONCLES ===" + Reset)
	fmt.Println("Étape 1 : Écraser les cornes de serpent et ajouter les sangsues dans le chaudron en étain.")
	fmt.Println("1 - Chauffer à feu doux et remuer 3 fois dans le sens des aiguilles d'une montre")
	fmt.Println("2 - Laisser bouillir à feu vif sans remuer")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Succès ! Le liquide dégage une douce vapeur rose. Potion parfaitement réussie (+10 Exppotions) !" + Reset)
		p.Exppotions += 10
		p.AjouterItem("Potion de Soin des Furoncles", TypeConsommable, 3, 1)
	} else {
		fmt.Println(Red + "Erreur de dosage ! Le chaudron émet un sifflement et se couvre de petites cloques verdâtres." + Reset)
		p.Health -= 5
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 2 : PHILTRE D'APAISEMENT ---

func TenterPhiltreApaisement(p *Player) {
	reader := Lecteur
	if p.Levelpotions < 11 {
		fmt.Println(Red + "Niveau insuffisant ! Vous devez être en 2ème année (Niveau 11 minimum)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Poudre de Pierre de Lune (Coût chez l'Apothicaire : 6 Gallions)
	if !VerifierEtProposerAchat(p, "Poudre de Pierre de Lune", 6) {
		return
	}

	fmt.Println(Cyan + "=== PRÉPARATION : PHILTRE D'APAISEMENT ===" + Reset)
	fmt.Println("Cette potion complexe calme l'agitation et dissipe les soucis.")
	fmt.Println("Combien de gouttes d'extrait de belladone ajoutez-vous au mélange frémissant ?")
	fmt.Println("1 - Exactement 3 gouttes")
	fmt.Println("2 - Une giclée généreuse au pif")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Le liquide devient d'un bleu argenté translucide. Rogue hoche imperceptiblement la tête (+20 Exppotions) !" + Reset)
		p.Exppotions += 20
		p.AjouterItem("Philtre d'Apaisement", TypeConsommable, 10, 1)
	} else {
		fmt.Println(Red + "SPOUTCH ! La potion devient visqueuse et toxique. Vous inhalez des fumées urticantes (-10 PV)." + Reset)
		p.Health -= 10
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 3 : POTION D'AIGUISAGE DE LA PENSÉE ---

func TenterPotionAiguisage(p *Player) {
	reader := Lecteur
	if p.Levelpotions < 21 {
		fmt.Println(Red + "Niveau insuffisant ! Vous devez être en 3ème année (Niveau 21 minimum)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Chrysope séchée (Coût chez l'Apothicaire : 4 Gallions)
	if !VerifierEtProposerAchat(p, "Chrysope séchée", 4) {
		return
	}

	fmt.Println(Cyan + "=== PRÉPARATION : POTION D'AIGUISAGE DE LA PENSÉE ===" + Reset)
	fmt.Println("Mini-jeu de précision : Quel rythme de coupe appliquez-vous aux racines de sopophoro ?")
	fmt.Println("1 - Coupes fines et nettes avec le couteau d'argent")
	fmt.Println("2 - Écraser brutalement avec le pilon en granit")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Le jus s'écoule parfaitement. La potion prend une teinte jaune lumineux (+35 Exppotions) !" + Reset)
		p.Exppotions += 35
		p.AjouterItem("Potion d'Aiguisage de la Pensée", TypeConsommable, 15, 1)
	} else {
		fmt.Println(Red + "Mauvaise méthode ! Le jus s'oxyde instantanément et dégage une odeur d'œuf pourri." + Reset)
		p.Health -= 15
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 5 : POTION DE POLYNECTAR ---

func TenterPotionPolynectar(p *Player) {
	reader := Lecteur
	if p.Levelpotions < 41 {
		fmt.Println(Red + "Niveau insuffisant ! Cette recette de 5ème année exige le niveau 41 minimum." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Peau de Serpent du Cap (Coût chez l'Apothicaire : 8 Gallions)[cite: 1, 2]
	if !VerifierEtProposerAchat(p, "Peau de Serpent du Cap", 8) {
		return
	}

	fmt.Println(Red + "=== PRÉPARATION TRÈS COMPLEXE : POTION DE POLYNECTAR ===" + Reset)
	fmt.Println("Le grimoire indique un processus de brassage long de plusieurs semaines condensé ici.")
	fmt.Println("Quelle est l'ultime étape avant d'incorporer le cheveu de la personne ciblée ?")
	fmt.Println("1 - Ajouter les cuillères de corne de bicorne pulvérisée au dernier moment")
	fmt.Println("2 - Faire bouillir le tout à feu ardent avec des peaux de serpent entières")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(2 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Miraculeux ! La potion bouillonne bruyamment avant de devenir d'une consistance boueuse et dorée." + Reset)
		fmt.Println("Vous obtenez une fiole de Polynectar authentique (+60 Exppotions) !")
		p.Exppotions += 60
		p.AjouterItem("Potion de Polynectar (Imparfaite)", TypeConsommable, 40, 1)
	} else {
		fmt.Println(Red + "Échec cuisant ! Le chaudron exppotionslose dans un geyser de boue noire corrosive (-25 PV)." + Reset)
		p.Health -= 25
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour nettoyer les dégâts]"+Reset)
}