package src

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Toursastronomie() {
	p := JoueurActuel
	reader := Lecteur

	// Initialisation du générateur aléatoire
	rand.Seed(time.Now().UnixNano())

	// Liste des 10 ciels étoilés ASCII
	ciels := []string{
		`
       .        .      *        .        .      *     .      .
     *   .   +      .       .     *   .      .       .     *
       .      _____       .        .        .      .    +
     .   +   /     \   .       *      +       .      *       .
            |   c   |     .       .       .        .
      *     \_____/   .      +        .        .      .    *
    .    .       .        .       .       .        *    .       .
`,
		`
        *   .       .   *    .        .       *      .     +
          +    .  (   )   .      *      +       .     *   .
        .    *   (     )      .      .       .      .
      *    .    (_______)   .    +       *      .     .     *
        .   +     .     .      .    .        .      +    .
`,
		`
      *   .     .        .     *        .        .    *     .
        .    *  \          .       +       .      .    +
      +   .      \   *         .        .       .     *
        .     *   \____   .        *        .       .        .
     *    .     .  \   \     +         .        *    .    *
`,
		`
        *--------*          .       *       .        .    +
         \        \     .       .       +       *    .      .
          *        *       *        .        .       .   *
           \        .   .      .       *        .     +
            *----------*          .        .       .      .
`,
		`
      .      *       .        .     *      .        .    *
        *   .     .     +      .        .     *    .      +
      +     .   /||\    .    *     +      .       .    .
        .      / || \      .      .      .      *    .   *
      *   .   /__||__\   .    +      .       .      .
`,
		`
        .     .       *        .        .       .    *     .
      *    .     .--.      .      +        *      .     +
        +      .-'    '-.     .       .        .    .
      .     * (  (O)    )   .     *       +      .    *
        .      '-.____.-'      .       .        .    *   .
`,
		`
      .     *       .        .       *       .      .    +
        +       .     .---.      .       +       *   .     .
      *    .    .    /     \    .    .       .      .   *
        .     +     |   O   |      .      *       .      .
      .    *    .    \     /    *    +       .      .    +
`,
		`
      *      .      .        .       *       .      .    *
        .   *    .   .  @ .    .       +       *    .     .
      +    .   .   @  @   .      .       .       .    +
        .    *   .   @  .      *       .       .    +   .
      *    .       .     .         .       *        .      *
`,
		`
      .     *      .~~~~~~.       .       *      .   .    *
        +      .  (  *  .  )   .     +       .     *    .
      *    .     (  . +   . )     .      .       .     +
        .    *    '~~~~~~~~'   *      .       +      .   .
`,
		`
          \  |  /        .       *       .       .   *    .
        --  *  --    .       +       .       *     .    +
          /  |  \        *       .       .       .    .
      .     .       .        .       +       .      *    *
        *       +       .        .       .       *    .
`,
	}

	for {
		ClearTerminal()

		// Choix aléatoire d'un ciel à chaque affichage
		cielActuel := ciels[rand.Intn(len(ciels))]

		fmt.Println(Yellow + "==========================================================================" + Reset)
		fmt.Println(Yellow + "                        LA TOUR D'ASTRONOMIE                              " + Reset)
		fmt.Println(Yellow + "==========================================================================" + Reset)
		fmt.Println("Vous vous tenez au sommet de la plus haute tour du château.")
		fmt.Println("L'air est frais et le ciel nocturne s'étend au-dessus de Poudlard.")

		// Affichage du ciel en cyan magique
		fmt.Println(Cyan + cielActuel + Reset)

		fmt.Println(DarkGray + "--------------------------------------------------------------------------" + Reset)
		fmt.Printf("Élève : %s | Maison : %s\n", p.Name, p.Maison)
		fmt.Println(DarkGray + "--------------------------------------------------------------------------" + Reset)
		fmt.Println()
		fmt.Println(" 1 - Observer une autre partie du ciel")
		fmt.Println(" 2 - Régler le télescope en cuivre (Mini-jeu)")
		fmt.Println(" 3 - Redescendre au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir le sac à dos (Inventaire)")
		fmt.Println(Yellow + "==========================================================================" + Reset)

		Option := strings.TrimSpace(strings.ToLower(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch Option {
		case "1":
			// Relance la boucle pour tirer un nouveau ciel
			continue

		case "2":
			ClearTerminal()
			ReglerTelescopeMiniJeu()

		case "3":
			ClearTerminal()
			Halldepoudlard()
			return

		case "inv":
			ClearTerminal()
			AfficherInventaire(Toursastronomie)
			return

		default:
			ClearTerminal()
		}
	}
}

// ReglerTelescopeMiniJeu lance le mini-jeu de calibration astronomique
func ReglerTelescopeMiniJeu() {
	p := JoueurActuel
	reader := Lecteur

	rand.Seed(time.Now().UnixNano())

	// Liste des astres à découvrir
	type Astre struct {
		Nom         string
		Description string
		XP          int
	}

	astres := []Astre{
		{Nom: "Les Anneaux de Saturne", Description: "Un spectacle éblouissant de roches et de glace dorée.", XP: 15},
		{Nom: "La Comète d'Icare", Description: "Une traînée de lumière argentée traversant la voûte céleste.", XP: 20},
		{Nom: "La Nébuleuse du Phénix", Description: "Un nuage de poussière magique aux teintes pourpres et écarlates.", XP: 25},
		{Nom: "La Constellation du Dragon", Description: "Les étoiles de la tête du dragon brillaient d'un éclat vert émeraude.", XP: 15},
		{Nom: "La Lune de Jupiter (Io)", Description: "Un petit disque volcanique parfaitement visible à côté de la géante gazeuse.", XP: 20},
	}

	astreCible := astres[rand.Intn(len(astres))]

	// Valeurs cibles cachées
	focalCible := rand.Intn(10) + 1 // 1 à 10
	angleCible := rand.Intn(10) + 1 // 1 à 10
	filtreCible := rand.Intn(5) + 1 // 1 à 5

	// Valeurs actuelles du télescope
	focalActuel := 1
	angleActuel := 1
	filtreActuel := 1

	essaisMax := 5
	essaisRestants := essaisMax

	ClearTerminal()
	fmt.Println(Yellow + "==========================================================================" + Reset)
	fmt.Println(Yellow + "               MINI-JEU : CALIBRATION DU TÉLESCOPE EN CUIVRE               " + Reset)
	fmt.Println(Yellow + "==========================================================================" + Reset)
	fmt.Println("Vous vous installez derrière le grand télescope en cuivre poli de la Tour.")
	fmt.Println("Un corps céleste rare traverse le ciel ce soir ! Ajustez les trois molettes")
	fmt.Println("pour obtenir une image parfaitement nette avant que l'astre ne disparaisse.")
	fmt.Println()

	for essaisRestants > 0 {
		fmt.Println(DarkGray + "--------------------------------------------------------------------------" + Reset)
		fmt.Printf("Réglages actuels -> Focale : %d/10 | Inclinaison : %d/10 | Filtre : %d/5\n", focalActuel, angleActuel, filtreActuel)
		fmt.Printf("Essais restants : "+Red+"%d/%d\n"+Reset, essaisRestants, essaisMax)
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println(" 1 - Ajuster la molette de Mise au Point (Focale)")
		fmt.Println(" 2 - Ajuster la molette d'Inclinaison (Angle)")
		fmt.Println(" 3 - Ajuster la molette du Filtre de Lumière")
		fmt.Println(" 4 - " + Green + "Regarder dans l'objectif (Valider l'observation)" + Reset)
		fmt.Println(" Q - Abandonner le télescope")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"\nChoix : "+Reset)))

		switch choix {
		case "1":
			saisie := LireSaisie(reader, "Entrez la nouvelle Focale (1-10) : ")
			val, err := strconv.Atoi(strings.TrimSpace(saisie))
			if err == nil && val >= 1 && val <= 10 {
				focalActuel = val
				ClearTerminal()
				fmt.Println(Green + "[✓] Molette de focale ajustée." + Reset)
			} else {
				ClearTerminal()
				fmt.Println(Red + "Valeur invalide (choisissez entre 1 et 10)." + Reset)
			}

		case "2":
			saisie := LireSaisie(reader, "Entrez le nouvel Angle d'inclinaison (1-10) : ")
			val, err := strconv.Atoi(strings.TrimSpace(saisie))
			if err == nil && val >= 1 && val <= 10 {
				angleActuel = val
				ClearTerminal()
				fmt.Println(Green + "[✓] Molette d'inclinaison ajustée." + Reset)
			} else {
				ClearTerminal()
				fmt.Println(Red + "Valeur invalide (choisissez entre 1 et 10)." + Reset)
			}

		case "3":
			saisie := LireSaisie(reader, "Entrez le numéro du Filtre (1-5) : ")
			val, err := strconv.Atoi(strings.TrimSpace(saisie))
			if err == nil && val >= 1 && val <= 5 {
				filtreActuel = val
				ClearTerminal()
				fmt.Println(Green + "[✓] Filtre de lumière ajusté." + Reset)
			} else {
				ClearTerminal()
				fmt.Println(Red + "Valeur invalide (choisissez entre 1 et 5)." + Reset)
			}

		case "4":
			ClearTerminal()
			fmt.Println(Cyan + "Vous collez votre œil contre la lentille froide en verre poli..." + Reset)
			time.Sleep(1 * time.Second)

			// Vérification du succès
			estFocaleOk := focalActuel == focalCible
			estAngleOk := angleActuel == angleCible
			estFiltreOk := filtreActuel == filtreCible

			if estFocaleOk && estAngleOk && estFiltreOk {
				// Victoire
				p.Exp += astreCible.XP
				fmt.Println("\n" + Green + "==========================================================================" + Reset)
				fmt.Println(Green + "                   [★] OBSERVATION PARFAITE RÉUSSIE ! [★]" + Reset)
				fmt.Println(Green + "==========================================================================" + Reset)
				fmt.Printf(Yellow+"Vous avez découvert : %s !\n"+Reset, astreCible.Nom)
				fmt.Println(astreCible.Description)
				fmt.Printf(Cyan+"\n[Bonus] Vous gagnez +%d Points d'Expérience Astronomique !\n"+Reset, astreCible.XP)
				fmt.Println(Green + "==========================================================================" + Reset)

				// Ajout d'une note dans l'inventaire si souhaité
				p.AjouterItem("Carte céleste : "+astreCible.Nom, TypeDivers, 2, 1)

				LireSaisie(reader, DarkGray+"\nAppuyez sur Entrée pour vous relever du télescope..."+Reset)
				ClearTerminal()
				return
			}

			// Indices si échec de la tentative
			essaisRestants--
			fmt.Println("\n" + Yellow + "=== INDICES DE LA LENTILLE ===" + Reset)

			if focalActuel < focalCible {
				fmt.Println(Red + "• Focale : L'image est floue et déformée (focale trop basse)." + Reset)
			} else if focalActuel > focalCible {
				fmt.Println(Red + "• Focale : L'image est trop zoomée et floue (focale trop haute)." + Reset)
			} else {
				fmt.Println(Green + "• Focale : [PARFAIT] La netteté de l'image est excellente !" + Reset)
			}

			if angleActuel < angleCible {
				fmt.Println(Red + "• Inclinaison : Vous pointez trop bas vers l'horizon." + Reset)
			} else if angleActuel > angleCible {
				fmt.Println(Red + "• Inclinaison : Vous pointez trop haut dans le ciel." + Reset)
			} else {
				fmt.Println(Green + "• Inclinaison : [PARFAIT] L'astre est centré dans l'objectif !" + Reset)
			}

			if filtreActuel < filtreCible {
				fmt.Println(Red + "• Filtre : La lumière est trop éblouissante (filtre trop faible)." + Reset)
			} else if filtreActuel > filtreCible {
				fmt.Println(Red + "• Filtre : Le champ de vision est trop sombre (filtre trop fort)." + Reset)
			} else {
				fmt.Println(Green + "• Filtre : [PARFAIT] La luminosité est idéale !" + Reset)
			}
			fmt.Println()

		case "Q":
			ClearTerminal()
			fmt.Println(Yellow + "Vous abandonnez les réglages du télescope." + Reset)
			return

		default:
			ClearTerminal()
			fmt.Println(Red + "Option invalide." + Reset)
		}
	}

	// Échec après épuisement des essais
	fmt.Println(Red + "==========================================================================" + Reset)
	fmt.Println(Red + "L'astre a poursuivi sa course et s'est perdu derrière les nuages..." + Reset)
	fmt.Println("Vous n'avez pas réussi à calibrer le télescope à temps.")
	fmt.Println(Red + "==========================================================================" + Reset)
	LireSaisie(reader, DarkGray+"\nAppuyez sur Entrée pour continuer..."+Reset)
	ClearTerminal()
}
