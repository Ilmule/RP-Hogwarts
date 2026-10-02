package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// CoursMetamorphose gère l'accès au cours de métamorphose du Professeur McGonagall
func Coursmétamorphose() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		// Calcul de l'année en fonction du niveau (1-10 = 1ère année, 11-20 = 2ème année, etc.)
		annee := (p.Levelmetamorphose-1)/10 + 1
		if annee > 7 {
			annee = 7
		}

		fmt.Println(Red + "==========================================================" + Reset)
		fmt.Println(Red + "               SALLE DE CLASSE DE MÉTAMORPHOSE            " + Reset)
		fmt.Println(Red + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp métamorphose : %d | Niveau : %d (Année %d)\n", p.Health, p.MaxHealth, p.Expmetamorphose, p.Levelmetamorphose, annee)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("La salle est vaste, baignée d'une lumière claire. Sur chaque pupitre")
		fmt.Println("reposent des objets hétéroclites en attente d'être transformés.")
		fmt.Println("Le Professeur Minerva McGonagall arpente les rangs d'un pas ferme,")
		fmt.Println("veillant à ce qu'aucun élève ne bâcle ses mouvements de baguette.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Écouter les exigences du Professeur McGonagall (Dialogue)")
		fmt.Println(" 2 - 1ère Année : Transformer une allumette en aiguille (Niveau 1+)")
		fmt.Println(" 3 - 2ème Année : Transformer un scarabée en bouton (Niveau 11+)")
		fmt.Println(" 4 - 3ème Année : Sortilège de Commutation (Niveau 21+)")
		fmt.Println(" 5 - Retourner au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Red + "==========================================================" + Reset)

		Option := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch Option {
		case "1":
			ClearTerminal()
			DialoguesMcGonagall()
		case "2":
			ClearTerminal()
			TenterAllumetteAiguille(p)
		case "3":
			ClearTerminal()
			TenterScarabeeBouton(p)
		case "4":
			ClearTerminal()
			TenterCommutation(p)
		case "5":
			ClearTerminal()
			fmt.Println(Yellow + "Vous saluez le professeur et quittez la salle de métamorphose..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard()
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Coursmétamorphose)
			return
		default:
			ClearTerminal()
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// DialoguesMcGonagall gère les remarques rigoureuses et professionnelles de la directrice adjointe
func DialoguesMcGonagall() {
	reader := Lecteur
	dialogues := []string{
		"« La métamorphose est l'une des magies les plus complexes et dangereuses que vous apprendrez à Poudlard. Ne l'prenez pas à la légère. »",
		"« Un mouvement de poignet imprécis, et vous risquez de vous retrouver avec une demi-tasse à la place d'une gerboise ! »",
		"« Concentrez-vous sur l'objet cible et visualisez sa nouvelle forme dans les moindres détails. »",
		"« La rigueur fait les grands sorciers. Redressez-vous et appliquez-vous, s'il vous plaît ! »",
	}
	phrase := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Red + "=== PROFESSEUR MINERVA MCGONAGALL ===" + Reset)
	fmt.Println("McGonagall s'arrête devant votre pupitre, plisse les yeux par-dessus ses carrés de lunettes.")
	fmt.Println()
	fmt.Printf(Cyan+"McGonagall : %s\n"+Reset, phrase)
	fmt.Println()

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour reprendre vos exercices]"+Reset)
}

// --- NIVEAU 1 : ALLUMETTE EN AIGUILLE ---

func TenterAllumetteAiguille(p *Player) {
	reader := Lecteur
	if p.Levelmetamorphose < 1 {
		fmt.Println(Red + "Niveau insuffisant (Niveau 1 requis)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== EXERCICE 1 : ALLUMETTE EN AIGUILLE ===" + Reset)
	fmt.Println("Une simple allumette en bois repose devant vous. Vous devez la piquer et la transformer en aiguille pointue.")
	fmt.Println("Quel mouvement de baguette effectuez-vous ?")
	fmt.Println("1 - Un coup sec et précis en forme de pointe vers le bas")
	fmt.Println("2 - Une grande boucle ample et circulaire")
	fmt.Println("3 - Tapoter l'allumette trois fois rapidement")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre méthode (1-3) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Brillant ! L'allumette frémit, s'affine et devient une aiguille en argent étincelante !" + Reset)
		fmt.Println("McGonagall hoche la tête avec satisfaction : « Très propre. 10 points pour votre maison. (+10 Expmetamorphose) »")
		p.Expmetamorphose += 10
	} else if choix == "3" {
		fmt.Println(Yellow + "Résultat partiel : L'allumette a bien changé de matière, mais elle conserve la forme d'un bout de bois tordu." + Reset)
		fmt.Println("McGonagall : « C'est un début, mais l'aiguille ne cousra pas grand-chose. (+5 Expmetamorphose) »")
		p.Expmetamorphose += 5
	} else {
		fmt.Println(Red + "Échec ! Le mouvement était trop large. L'allumette prend feu toute seule et se consume en cendres." + Reset)
		p.Health -= 3
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 2 : SCARABÉE EN BOUTON ---

func TenterScarabeeBouton(p *Player) {
	reader := Lecteur
	if p.Levelmetamorphose < 11 {
		fmt.Println(Red + "Niveau insuffisant ! Vous devez être en 2ème année (Niveau 11 minimum)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== EXERCICE 2 : SCARABÉE EN BOUTON DE MANCHETTE ===" + Reset)
	fmt.Println("Un petit scarabée vert bouge frénétiquement dans une petite boîte en bois.")
	fmt.Println("Comment prononcez-vous la formule d'impulsion ?")
	fmt.Println("1 - « Vera Verto ! » d'une voix claire et assurée")
	fmt.Println("2 - « Vera... vert-o ? » en hésitant à mi-chemin")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Le scarabée se fige, se solidifie et se transforme en un splendide bouton de nacre sculpté !" + Reset)
		fmt.Println("McGonagall sourit légèrement : « Excellent travail de précision. (+20 Expmetamorphose) »")
		p.Expmetamorphose += 20
	} else {
		fmt.Println(Red + "Catastrophe ! Le scarabée panique, change brièvement de couleur avant de se remettre à courir partout sur votre pupitre." + Reset)
		p.Health -= 8
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 3 : SORTILÈGE DE COMMUTATION ---

func TenterCommutation(p *Player) {
	reader := Lecteur
	if p.Levelmetamorphose < 21 {
		fmt.Println(Red + "Niveau insuffisant ! Vous devez être en 3ème année (Niveau 21 minimum)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== EXERCICE AVANCÉ : COMMUTATION DE LA MATIÈRE ===" + Reset)
	fmt.Println("L'exercice consiste à permuter les propriétés physiques de deux objets distincts (une plume et une pierre).")
	fmt.Println("Quelle stratégie de concentration adoptez-vous ?")
	fmt.Println("1 - Diviser mentalement votre attention en deux faisceaux égaux vers chaque objet")
	fmt.Println("2 - Forcer l'échange par une simple décharge d'énergie brute")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(2 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Remarquable ! La plume devient lourde comme la pierre, tandis que la pierre s'envole au premier souffle d'air !" + Reset)
		fmt.Println("McGonagall : « Une maîtrise de la commutation digne d'un futur ASPIC ! (+40 Expmetamorphose) »")
		p.Expmetamorphose += 40
	} else {
		fmt.Println(Red + "Surcharge magique ! L'énergie rebondit, pulvérisant la pierre en poussière et collant la plume à votre front." + Reset)
		p.Health -= 15
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour nettoyer le bureau]"+Reset)
}