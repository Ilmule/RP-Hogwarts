package src

import (
	"fmt"
	"strings"
	"time"
)

// ChambreDesSecrets gère l'exploration de la salle cachée de Salazar Serpentard
func Secretchambredessecrets() {
	p := JoueurActuel
	reader := Lecteur

	ClearTerminal()

	// 1. Vérification du Fourchelang (On simule la présence d'une compétence ou d'un livre dans l'inventaire)
	saitFourchelang := false
	for _, item := range p.Inventaire {
		if strings.Contains(strings.ToLower(item.Nom), "fourchelang") {
			saitFourchelang = true
			break
		}
	}

	if !saitFourchelang {
		fmt.Println(Red + "=== ACCÈS IMPOSSIBLE ===" + Reset)
		fmt.Println("Vous vous tenez devant le lavabo orné d'un petit serpent gravé sur le robinet.")
		fmt.Println("Vous essayez de lui parler, mais seuls des sifflements ridicules sortent de votre bouche.")
		fmt.Println(DarkGray + "Il vous manque la capacité de parler le Fourchelang (à apprendre à la bibliothèque)." + Reset)
		time.Sleep(3 * time.Second)
		ClearTerminal()
		// Retourne dans les toilettes de Mimi Geignarde
		Secretsmimigeignarde() 
		return
	}

	fmt.Println(Green + "Vous sifflez une phrase gutturale en Fourchelang..." + Reset)
	time.Sleep(1 * time.Second)
	fmt.Println(Yellow + "Le lavabo s'écarte dans un grincement sourd, révélant un immense tuyau sombre." + Reset)
	fmt.Println("Vous glissez à l'intérieur et atterrissez sur un tas d'ossements de petits animaux.")
	time.Sleep(2 * time.Second)

	for {
		fmt.Println("\n" + Green + "==========================================================" + Reset)
		fmt.Println(Green + "                 LA CHAMBRE DES SECRETS                   " + Reset)
		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Println("La pièce est immense, soutenue par des piliers de pierre enlacés de serpents sculptés.")
		fmt.Println("L'atmosphère est glaciale et une eau noire tapisse le sol.")
		fmt.Println("Au fond, l'imposante statue du visage de Salazar Serpentard vous domine.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Explorer les recoins sombres de la Chambre")
		fmt.Println(" 2 - S'approcher du gigantesque squelette du Basilic")
		fmt.Println(" 3 - Remonter par le tuyau (Retour)")
		fmt.Println(" INV - Ouvrir votre inventaire")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			fmt.Println(Cyan + "=== EXPLORATION DE LA CHAMBRE ===" + Reset)
			fmt.Println("Vous marchez prudemment dans l'eau glacée. Le silence est de plomb.")
			fmt.Println("Vous remarquez d'anciennes mues de serpent géantes abandonnées dans les recoins,")
			fmt.Println("mais elles tombent en poussière dès que vous les touchez. L'endroit est mort.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "2":
			ClearTerminal()
			InteragirSqueletteBasilic()

		case "3":
			ClearTerminal()
			fmt.Println(Yellow + "Grâce à un sortilège de lévitation, vous parvenez à remonter le tuyau..." + Reset)
			time.Sleep(1 * time.Second)
			Secretsmimigeignarde()
			return

		case "INV":
			ClearTerminal()
			AfficherInventaire(Secretchambredessecrets)
			return

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// InteragirSqueletteBasilic gère la récolte limitée des crocs
func InteragirSqueletteBasilic() {
	p := JoueurActuel
	reader := Lecteur
	limiteCrocs := 5 // Le joueur ne peut posséder que 5 crocs maximum en même temps

	fmt.Println(Red + "=== LE SQUELETTE DU BASILIC ===" + Reset)
	fmt.Println("Devant la statue de Serpentard gît la gigantesque carcasse du Basilic, terrassé jadis.")
	fmt.Println("Ses os immenses sont encore intacts. Dans sa mâchoire, plusieurs crocs longs comme des épées")
	fmt.Println("sont encore solidement attachés, gorgés d'un venin fossilisé d'une valeur inestimable.")
	fmt.Println("----------------------------------------------------------")
	fmt.Println(" 1 - Tenter d'arracher un croc de Basilic (Valeur de revente élevée)")
	fmt.Println(" R - Reculer")

	choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nAction : ")))

	if choix == "R" {
		ClearTerminal()
		return
	}

	if choix == "1" {
		// Vérification du nombre de crocs que le joueur possède déjà
		crocsPossedes := 0
		for _, item := range p.Inventaire {
			if item.Nom == "Croc de Basilic" {
				crocsPossedes = item.Quantite
				break
			}
		}

		if crocsPossedes >= limiteCrocs {
			fmt.Println(Red + "\nVous avez déjà récupéré suffisamment de crocs." + Reset)
			fmt.Println("Les autres sont trop endommagés ou trop profondément enfoncés dans la mâchoire.")
			fmt.Println(DarkGray + "(Limite atteinte : " + fmt.Sprint(limiteCrocs) + " crocs maximum. Vendez-les d'abord !)" + Reset)
		} else {
			fmt.Println(Yellow + "\nVous tirez de toutes vos forces sur l'un des crocs... CRAC !" + Reset)
			fmt.Println(Green + "Vous avez récupéré un Croc de Basilic ! Faites attention à ne pas vous piquer." + Reset)
			
			// Ajout dans l'inventaire avec une forte valeur (ex: 50 Gallions à la revente)
			p.AjouterItem("Croc de Basilic", TypeDivers, 50, 1)
		}
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)
		ClearTerminal()
	} else {
		fmt.Println(Red + "Choix invalide." + Reset)
	}
}