package src

import (
	"fmt"
	"strings"
	"time"
)

// GrandeSalle gère le lieu de rassemblement, le banquet et la redirection vers le Hall.
func GrandeSalle() {
	p := JoueurActuel
	reader := Lecteur

	// 🔒 VÉRIFICATION OBLIGATOIRE DE LA RÉPARTITION
	if p.Maison == "" || p.Maison == "Non attribuée" {
		p.Maison = CeremonieChoixpeau()
	}

	for {
		fmt.Println("\n" + Yellow + "==========================================================" + Reset)
		fmt.Println(Yellow + "                    LA GRANDE SALLE                       " + Reset)
		fmt.Println(Yellow + "==========================================================" + Reset)

		fmt.Printf("Maison : " + Cyan + "%s" + Reset + " | Titre : %s\n", p.Maison, p.Title)
		fmt.Printf("Santé : %d/%d PV\n", p.Health, p.MaxHealth)
		fmt.Println("----------------------------------------------------------")

		fmt.Println("Les quatre grandes tables des maisons sont garnies de plats étincelants.")
		fmt.Println("Au-dessus de vos têtes, le plafond magique reflète le ciel extérieur.")

		fmt.Println()

		fmt.Println(" 1 - Manger au banquet avec les élèves de " + p.Maison)
		fmt.Println(" 2 - Retourner dans le Hall de Poudlard (Naviguer dans le château)")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Yellow + "==========================================================" + Reset)

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch choix {
		case "1":
			ClearTerminal()
			MangerAuBanquet()

		case "2":
			ClearTerminal()
			fmt.Println(Yellow + "Vous quittez les grandes tables pour retourner dans le Hall..." + Reset)
			Halldepoudlard()
			return

		case "INV":
			ClearTerminal()
			AfficherInventaire(GrandeSalle)
			return

		default:
			fmt.Println(Red + "Choix invalide. Veuillez sélectionner une option disponible." + Reset)
		}
	}
}

// MangerAuBanquet simule le repas avec les camarades de maison et restaure les PV du joueur
func MangerAuBanquet() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println("\n" + Yellow + "=== BANQUET DE LA MAISON " + strings.ToUpper(p.Maison) + " ===" + Reset)
	fmt.Println("Vous prenez place à la grande table aux côtés de vos camarades.")
	time.Sleep(1 * time.Second)

	// Dialogues immersifs adaptés selon la maison du joueur
	switch p.Maison {
	case "Gryffondor":
		fmt.Println(Red + "Un camarade de Gryffondor : « Passe-moi le jus de potiron ! T'as vu le dernier entraînement de Quidditch ? On va tous les écraser cette année ! »" + Reset)
	case "Serpentard":
		fmt.Println(Green + "Un camarade de Serpentard : « Installe-toi. Les autres maisons pensent encore pouvoir rivaliser pour la Coupe des Quatre Maisons... Ridicule. »" + Reset)
	case "Serdaigle":
		fmt.Println(Cyan + "Un camarade de Serdaigle : « Tu as révisé ton cours de Métamorphose ? Le professeur McGonagall va interroger tout le monde cet après-midi ! »" + Reset)
	case "Poufsouffle":
		fmt.Println(Yellow + "Un camarade de Poufsouffle : « Goûte ces tartes aux mélasses, elles sont succulentes ! Servons-nous à manger avant que les assiettes ne se vident ! »" + Reset)
	default:
		fmt.Println("Des plats dorés remplis de dinde rôtie, de gratin et de friandises apparaissent magiquement sous vos yeux.")
	}

	fmt.Println("\nVous mangez copieusement et profitez de la chaleur du repas...")
	time.Sleep(2 * time.Second)

	// Restauration complète de la santé du joueur
	p.Health = p.MaxHealth
	fmt.Println(Green + "\n[✓] Le repas vous réconforte ! Vos points de vie sont entièrement restaurés (" + fmt.Sprintf("%d/%d PV", p.Health, p.MaxHealth) + ")." + Reset)
	fmt.Println()

	LireSaisie(reader, DarkGray+"Appuyez sur Entrée pour vous relever de la table..." + Reset)
	ClearTerminal()
}