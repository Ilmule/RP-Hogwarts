package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Cachots gère l'exploration de la salle de classe des Potions dans les sous-sols
func Soussolscachots() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Println(Green + "                  LES CACHOTS (POTIONS)                   " + Reset)
		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp : %d | Niveau : %d\n", p.Health, p.MaxHealth, p.Exp, p.Level)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Il fait un froid glacial ici. Les murs de pierre suintent d'humidité.")
		fmt.Println("Des animaux conservés flottent dans des bocaux en verre tout autour de la pièce.")
		fmt.Println("Au fond, le Professeur Rogue corrige des parchemins d'un air lugubre.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Parler au Professeur Rogue")
		fmt.Println(" 2 - Fouiller discrètement les bocaux et étagères")
		fmt.Println(" 3 - S'installer à un chaudron (Mini-jeu de Potion)")
		fmt.Println(" 4 - Remonter les escaliers froids (Retour)")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Green + "==========================================================" + Reset)

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch choix {
		case "1":
			ClearTerminal()
			DialoguesRogue()
		case "2":
			ClearTerminal()
			InspecterBocaux()
		case "3":
			ClearTerminal()
			MiniJeuPotion()
		case "4":
			ClearTerminal()
			fmt.Println(Yellow + "Vous remontez prudemment vers les étages supérieurs..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard() 
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Soussolscachots)
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// DialoguesRogue génère des remarques sarcastiques de Severus Rogue
func DialoguesRogue() {
	p := JoueurActuel
	reader := Lecteur
	
	dialogues := []string{
		"« Qu'y a-t-il, %s ? Votre esprit est-il trop lent pour comprendre les instructions écrites au tableau ? »",
		"« Je pourrais vous apprendre à mettre la gloire en bouteille... mais je doute fort que vous en ayez les capacités. »",
		"« Ne respirez pas si fort près de mon bureau. Vous dissipez les effluves de mon Polynectar. »",
		"« Cinq points en moins pour %s... Juste parce que votre présence m'insupporte aujourd'hui. »",
	}

	fmt.Println(Green + "=== PROFESSEUR ROGUE ===" + Reset)
	fmt.Println("Vous vous approchez du bureau. Rogue lève lentement ses yeux noirs et froids vers vous.")
	time.Sleep(1 * time.Second)
	
	phrase := dialogues[rand.Intn(len(dialogues))]
	fmt.Printf("\n"+Cyan+"Rogue : "+phrase+Reset+"\n", p.Name)
	
	if strings.Contains(phrase, "Cinq points en moins") {
		fmt.Printf(Red+"(Rogue vient vraiment de pénaliser %s pour rien...) \n"+Reset, p.Maison)
	}

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour reculer sans faire de bruit]"+Reset)
}

// InspecterBocaux permet de looter des ingrédients de potions
func InspecterBocaux() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Green + "=== ÉTAGÈRES DES CACHOTS ===" + Reset)
	fmt.Println("Vous inspectez les bocaux remplis de liquides troubles et de créatures visqueuses.")
	time.Sleep(1 * time.Second)

	chance := rand.Intn(100)
	if chance < 30 {
		fmt.Println(Red + "Vous faites grincer un bocal. Rogue se retourne brusquement : « Que faites-vous là ?! »" + Reset)
		fmt.Println("Vous fuyez avant de perdre des points !")
	} else if chance < 70 {
		fmt.Println(Yellow + "Rien d'utile ici, à part un vieux rat confit dans du vinaigre." + Reset)
	} else {
		nomIngredient := "Yeux de scarabées (Poignée)"
		if chance > 90 {
			nomIngredient = "Peau de Serpent du Cap"
		}
		fmt.Println(Cyan + "Vous trouvez un ingrédient non étiqueté qui traîne derrière un chaudron !" + Reset)
		fmt.Printf(Green+"Vous avez récupéré : %s.\n"+Reset, nomIngredient)
		p.AjouterItem(nomIngredient, TypeConsommable, 2, 1)
	}

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)
}

// MiniJeuPotion est un test de mémorisation court et punitif
func MiniJeuPotion() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Green + "=== PRÉPARATION DE LA POTION DE GUÉRISON ===" + Reset)
	fmt.Println("Vous allumez un feu sous un chaudron en cuivre. Le manuel indique 3 étapes précises.")
	fmt.Println("Si vous vous trompez, le mélange risque de vous exploser au visage.")
	time.Sleep(2 * time.Second)

	// Étape 1
	fmt.Println("\nÉtape 1 : Le liquide frémit.")
	fmt.Println("1 - Ajouter du jus de Flobberworm")
	fmt.Println("2 - Ajouter des crocs de serpent écrasés")
	choix1 := strings.TrimSpace(LireSaisie(reader, "Votre choix : "))

	if choix1 != "2" {
		explosionPotion(p, "Le liquide vire au vert acide et dégage une fumée toxique !")
		return
	}
	fmt.Println(Cyan + "Le mélange prend une belle teinte bleue. Parfait." + Reset)

	// Étape 2
	fmt.Println("\nÉtape 2 : Température de cuisson.")
	fmt.Println("1 - Laisser bouillir à feu vif")
	fmt.Println("2 - Éteindre le feu et remuer doucement")
	choix2 := strings.TrimSpace(LireSaisie(reader, "Votre choix : "))

	if choix2 != "1" {
		explosionPotion(p, "La potion s'épaissit comme du béton. Le chaudron se fissure !")
		return
	}
	fmt.Println(Cyan + "Des étincelles roses s'échappent du chaudron. Vous êtes sur la bonne voie." + Reset)

	// Étape 3
	fmt.Println("\nÉtape 3 : La touche finale.")
	fmt.Println("1 - Ajouter des piquants de porc-épic")
	fmt.Println("2 - Ajouter un bézoard")
	choix3 := strings.TrimSpace(LireSaisie(reader, "Votre choix : "))

	// Rogue précise toujours d'enlever le chaudron du feu avant les piquants de porc-épic !
	if choix3 != "1" {
		explosionPotion(p, "La réaction chimique est désastreuse !")
		return
	}

	fmt.Println(Green + "\nSuccès ! La potion devient limpide." + Reset)
	fmt.Println("Vous remplissez une fiole de ce précieux liquide. (+15 Exp)")
	p.Exp += 15
	p.AjouterItem("Potion de Guérison (+30 PV)", TypeConsommable, 5, 1)

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour nettoyer votre plan de travail]"+Reset)
}

// explosionPotion gère l'échec du mini-jeu
func explosionPotion(p *Player, message string) {
	reader := Lecteur
	degats := 15
	fmt.Println(Red + "\n" + message + Reset)
	fmt.Println(Red + "BOUM ! Le chaudron explose !" + Reset)
	
	p.Health -= degats
	if p.Health < 1 {
		p.Health = 1
	}
	
	fmt.Printf("Vous êtes couvert de suie et perdez %d PV. Rogue vous fixe d'un air méprisant.\n", degats)
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour fuir l'atelier]"+Reset)
}