package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Infirmerie gère l'accès à l'infirmerie de Poudlard avec Madame Pomfresh
func Toursinfirmerie() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Println(Cyan + "                    L'INFIRMERIE DE POUDLARD              " + Reset)
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("L'endroit est d'une propreté immaculée, baigné d'une lumière blanche.")
		fmt.Println("Plusieurs lits alignés accueillent des élèves convalescents.")
		fmt.Println("Madame Pomfresh va et vient entre les rangées, une fiole à la main.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Parler à Madame Pomfresh (Soins complets & Remèdes)")
		fmt.Println(" 2 - Rendre visite aux élèves alités (Dialogues)")
		fmt.Println(" 3 - Mini-jeu : Gober le Remède de Madame Pomfresh")
		fmt.Println(" 4 - Retourner vers le Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Cyan + "==========================================================" + Reset)

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch choix {
		case "1":
			ClearTerminal()
			MadamePomfreshSoins()
		case "2":
			ClearTerminal()
			VisiterPatients()
		case "3":
			ClearTerminal()
			MiniJeuRemede()
		case "4":
			ClearTerminal()
			fmt.Println(Yellow + "Vous quittez l'infirmerie sur la pointe des pieds..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard()
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Toursinfirmerie)
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// MadamePomfreshSoins soigne le joueur et propose des dialogues stricts de l'infirmière
func MadamePomfreshSoins() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Cyan + "=== MADAME POMFRESH ===" + Reset)
	fmt.Println("Madame Pomfresh vous aperçoit et fronce immédiatement les sourcils.")
	fmt.Println("Elle pose ses mains sur ses hanches :")
	fmt.Println(Yellow + "« Alors, qu'est-ce qui ne va pas cette fois ? Un match de Quidditch raté ? Une potion qui a explosé au visage ? »" + Reset)
	time.Sleep(1 * time.Second)

	if p.Health < p.MaxHealth {
		p.Health = p.MaxHealth
		fmt.Println(Green + "\nElle soupire, vous force à avaler une gorgée de potion Pousse-os et un philtre revigorant." + Reset)
		fmt.Println(Green + "Vos PV sont restaurés au maximum ! « Et essayez d'être plus prudent la prochaine fois ! »" + Reset)
	} else {
		fmt.Println(Green + "\nElle vous tâte le pouls, hausse les épaules et sourit : " +
			"« Vous êtes en parfaite santé ! Ne restez pas à traîner ici si vous n'êtes pas blessé. »" + Reset)
	}

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)
}

// VisiterPatients génère des dialogues avec des élèves alités
func VisiterPatients() {
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	patients := []string{
		"Un élève de Gryffondor murmure : « Je crois que j'ai avalé de travers une pastille de vomissement des jumeaux Weasley... »",
		"Une élève de Serdaigle lit un livre de sortilèges en portant d'immenses mitaines en fourrure : « J'ai raté mon enchantement de réchauffement. Mes mains sont gelées pour la journée. »",
		"Un joueur de l'équipe de Quidditch gémit : « Un souaffle m'a atterri en plein visage pendant l'entraînement... J'ai le nez complètement déplacé. »",
	}

	patientChoisi := patients[rand.Intn(len(patients))]

	fmt.Println(Cyan + "=== LES LITS DE L'INFIRMERIE ===" + Reset)
	fmt.Println("Vous marchez discrètement entre les lits pour voir vos camarades souffrir en silence.")
	fmt.Println()
	fmt.Println(Yellow + patientChoisi + Reset)
	fmt.Println()

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour quitter les lits des patients]"+Reset)
}

// MiniJeuRemede : Avaler un remède corsé de Madame Pomfresh
func MiniJeuRemede() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Cyan + "=== LE REMÈDE DE MADAME POMFRESH ===" + Reset)
	fmt.Println("Madame Pomfresh vous tend une petite fiole contenant un liquide fumant et verdâtre.")
	fmt.Println("« C'est pour stimuler votre magie interne. Avalez d'un trait, sans respirer ! »")
	time.Sleep(1 * time.Second)

	fmt.Println("\n1 - Avaler d'un coup sec")
	fmt.Println("2 - Hésiter et flairer le contenu")
	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix : "))

	if choix == "1" {
		fmt.Println(Yellow + "\nVous fermez les yeux et buvez la mixture... De la fumée commence à sortir de vos oreilles !" + Reset)
		gain := rand.Intn(20) + 10
		p.Exp += gain
		fmt.Printf(Green+"Votre corps est traversé d'une puissante énergie magique ! Vous gagnez %d points d'expérience !\n"+Reset, gain)
	} else {
		fmt.Println(Red + "\nVous hésitez trop longtemps. Madame Pomfresh vous arrache la fiole des mains d'un air agacé :" + Reset)
		fmt.Println(Red + "« Du caractère, bon sang ! Ce n'est pas du jus de citrouille ! » (Aucun effet)" + Reset)
	}

	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour reprendre vos esprits]"+Reset)
}