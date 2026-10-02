package src

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CuisinesPoudlard gère l'accès aux cuisines et l'interaction avec les elfes de maison
func Soussolscuisine() {
	reader := Lecteur

	// Liste des prénoms d'elfes de maison connus dans la saga Harry Potter
	nomsElfes := []string{"Dobby", "Winky", "Kreattur", "Hokey", "Pitiponquis"}

	for {
		fmt.Println("\n" + Yellow + "=================== LES CUISINES DE POUDLARD ===================" + Reset)
		fmt.Println("Vous vous trouvez dans une immense salle située sous la Grande Salle.")
		fmt.Println("Des dizaines de petits elfes de maison s'activent bruyamment autour")
		fmt.Println("de grandes marmites en cuivre et de fûts de bièrobeurre.")
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println(" 1 - Parler avec un elfe de maison")
		fmt.Println(" 2 - Repartir vers les sous-sols et le Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Yellow + "================================================================================" + Reset)

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch choix {
		case "1":
			ClearTerminal()
			// Choisir un prénom d'elfe au hasard pour cette interaction
			rand.Seed(time.Now().UnixNano())
			nomElfe := nomsElfes[rand.Intn(len(nomsElfes))]
			InteractionElfe(nomElfe)

		case "2":
			ClearTerminal()
			fmt.Println(Yellow + "Vous quittez les odeurs de tartes pour remonter vers le château..." + Reset)
			Halldepoudlard()
			return

		case "INV":
			ClearTerminal()
			AfficherInventaire(Soussolscuisine)
			return

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// InteractionElfe gère la séquence de dialogue détaillée et laisse le choix au joueur
func InteractionElfe(nomElfe string) {
	reader := Lecteur

	// --- ÉTAPE 1 : DIALOGUE D'INTRODUCTION NARRATIF ---
	fmt.Println("\n" + Yellow + "=== RENCONTRE AVEC " + strings.ToUpper(nomElfe) + " L'ELFE DE MAISON ===" + Reset)
	fmt.Printf("Un petit elfe nommé %s, aux yeux aussi grands que des balles de tennis, s'approche de vous\n", nomElfe)
	fmt.Println("en faisant une profonde révérence qui fait traîner ses longues oreilles au sol.")
	fmt.Println()
	fmt.Printf(Cyan+"%s : « Oh ! Bienvenue aux cuisines, jeune maître ! »\n"+Reset, nomElfe)
	fmt.Printf(Cyan+"« %s est si honoré(e) de recevoir un élève de Poudlard ! »\n"+Reset, nomElfe)
	fmt.Println()

	// Pause de lecture
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour continuer à lui parler...]" + Reset)
	ClearTerminal()

	fmt.Println("\n" + Yellow + "=== RENCONTRE AVEC " + strings.ToUpper(nomElfe) + " (SUITE) ===" + Reset)
	fmt.Printf("Vous remarquez que %s porte un vieux torchon tout râpé en guise de tunique.\n", nomElfe)
	fmt.Println("Ses yeux s'attardent avec curiosité sur les affaires que vous portez à l'épaule...")
	fmt.Println()
	fmt.Printf(Cyan+"%s (murmure doucement) : « Les elfes travaillent dur pour le château... »\n"+Reset, nomElfe)
	fmt.Printf(Cyan+"« Mais un vrai vêtement à soi... Oh, ce serait le plus beau des cadeaux... »\n"+Reset)
	fmt.Println()

	// Pause de lecture
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour prendre une décision...]" + Reset)
	ClearTerminal()

	// --- ÉTAPE 2 : LE CHOIX PROPOSÉ AU JOUEUR ---
	for {
		fmt.Println("\n" + Yellow + "=== QUE SOUHAITEZ-VOUS FAIRE AVEC " + strings.ToUpper(nomElfe) + " ? ===" + Reset)
		fmt.Println(" 1 - Donner un vêtement à l'elfe (depuis votre sac)")
		fmt.Println(" 2 - Demander poliment à manger")
		fmt.Println(" 3 - Ne rien faire et repartir sagement aux cuisines")
		fmt.Println()

		choixAction := strings.TrimSpace(LireSaisie(reader, DarkGray+"Votre choix : "+Reset))

		switch choixAction {
		case "1":
			ClearTerminal()
			DonnerVetementMenu(nomElfe)
			return

		case "2":
			ClearTerminal()
			DemanderAManger(nomElfe)
			return

		case "3":
			ClearTerminal()
			fmt.Printf(Yellow+"Vous saluez %s et vous reculez doucement.\n"+Reset, nomElfe)
			return

		default:
			fmt.Println(Red + "Choix invalide. Veuillez choisir 1, 2 ou 3." + Reset)
		}
	}
}

// DemanderAManger gère la demande de nourriture (avec grande chance de refus de l'elfe)
func DemanderAManger(nomElfe string) {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println("\n" + Yellow + "=== VOUS DEMANDEZ À MANGER À " + strings.ToUpper(nomElfe) + " ===" + Reset)
	fmt.Printf("Vous demandez gentiment à %s si vous pouvez avoir un petit quelque chose à grignoter.\n\n", nomElfe)

	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	chanceRefus := rand.Intn(100) // 0 à 99

	if chanceRefus < 75 {
		// 75% de chance de REFUS
		fmt.Printf(Cyan+"%s tortille nerveusement ses doigts et baisse les yeux :\n"+Reset, nomElfe)
		fmt.Printf(Cyan+"« Oh non, non ! %s est tellement désolé(e) ! Le professeur McGonagall et les chefs ont dit que les repas sont servis uniquement à la Grande Salle pendant les heures de banquet ! »\n"+Reset, nomElfe)
		fmt.Printf(Cyan+"« %s ne peut pas enfreindre les règles du château... Pardonnez-moi, jeune maître ! »\n"+Reset, nomElfe)
		fmt.Println()
		fmt.Println(Red + "(L'elfe refuse poliment de vous servir à manger en dehors des heures officielles.)" + Reset)
	} else {
		// 25% de chance d'ACCEPTATION
		fmt.Printf(Cyan+"%s esquisse un grand sourire et s'incline profondément :\n"+Reset, nomElfe)
		fmt.Printf(Cyan+"« Eh bien... juste une petite gourmandise alors ! Mais ne le dites pas aux professeurs ! »\n"+Reset)
		
		plat := "Petit biscuit au gingembre magique (+10 PV)"
		p.AjouterItem(plat, TypeConsommable, 1, 1)

		fmt.Println("\n" + Green + "================================================================" + Reset)
		fmt.Printf(Green+"[SUCCÈS] %s vous glisse en douce : %s !\n"+Reset, nomElfe, plat)
		fmt.Println(Green + "================================================================" + Reset)
	}

	fmt.Println()
	LireSaisie(reader, DarkGray+"Appuyez sur Entrée pour continuer..." + Reset)
	ClearTerminal()
}

// DonnerVetementMenu permet de choisir et donner un vêtement de l'inventaire
func DonnerVetementMenu(nomElfe string) {
	p := JoueurActuel
	reader := Lecteur

	// Filtrer l'inventaire pour ne garder que les vêtements
	type VetementOption struct {
		IndexInventaire int
		Nom             string
		Quantite        int
	}

	var vetementsDispos []VetementOption

	for i, item := range p.Inventaire {
		nomLower := strings.ToLower(item.Nom)
		estVetement := item.Type == TypeEquipement ||
			strings.Contains(nomLower, "robe") ||
			strings.Contains(nomLower, "chapeau") ||
			strings.Contains(nomLower, "gants") ||
			strings.Contains(nomLower, "cape") ||
			strings.Contains(nomLower, "chaussette") ||
			strings.Contains(nomLower, "bonnet") ||
			strings.Contains(nomLower, "gilet") ||
			strings.Contains(nomLower, "manteau") ||
			strings.Contains(nomLower, "bottes") ||
			strings.Contains(nomLower, "echarpe") ||
			strings.Contains(nomLower, "cravate") ||
			strings.Contains(nomLower, "ceinture")

		if estVetement {
			vetementsDispos = append(vetementsDispos, VetementOption{
				IndexInventaire: i,
				Nom:             item.Nom,
				Quantite:        item.Quantite,
			})
		}
	}

	if len(vetementsDispos) == 0 {
		fmt.Println("\n" + Red + "=== AUCUN VÊTEMENT DISPONIBLE ===" + Reset)
		fmt.Println("Vous fouillez votre sac mais vous n'avez aucun vêtement (robe, cape, gants...) à lui offrir.")
		fmt.Printf(Cyan+"\n%s : « Ce n'est pas grave du tout, jeune maître ! »\n"+Reset, nomElfe)
		
		LireSaisie(reader, DarkGray+"\nAppuyez sur Entrée pour revenir aux cuisines..." + Reset)
		ClearTerminal()
		return
	}

	fmt.Printf("\n=== SÉLECTION DU VÊTEMENT À OFFRIR À %s ===\n", strings.ToUpper(nomElfe))
	fmt.Println("Choisissez un vêtement dans votre sac :")
	fmt.Println()

	for i, v := range vetementsDispos {
		fmt.Printf(" %d - Offrir : %s (x%d)\n", i+1, v.Nom, v.Quantite)
	}
	fmt.Println()
	fmt.Println(" Q - Changer d'avis")

	choix := LireSaisie(reader, "\nEntrez le numéro du vêtement ou 'Q' : ")
	if strings.ToLower(strings.TrimSpace(choix)) == "q" {
		ClearTerminal()
		return
	}

	idx, err := strconv.Atoi(strings.TrimSpace(choix))
	if err != nil || idx < 1 || idx > len(vetementsDispos) {
		fmt.Println(Red + "Choix invalide." + Reset)
		return
	}

	vetementChoisi := vetementsDispos[idx-1]
	itemInventaire := &p.Inventaire[vetementChoisi.IndexInventaire]

	// Retrait d'un vêtement de l'inventaire
	itemInventaire.Quantite--
	nomArticleDonne := itemInventaire.Nom

	if itemInventaire.Quantite <= 0 {
		p.Inventaire = append(p.Inventaire[:vetementChoisi.IndexInventaire], p.Inventaire[vetementChoisi.IndexInventaire+1:]...)
	}

	// Réactions de l'elfe libre
	ClearTerminal()
	fmt.Printf(Green+"Vous sortez « %s » et le tendez à %s.\n"+Reset, nomArticleDonne, nomElfe)
	fmt.Println()
	fmt.Printf(Yellow+"%s écarquille de grands yeux ronds et regarde le vêtement avec émotion.\n"+Reset, nomElfe)
	fmt.Println(Yellow + "Une grosse larme coule sur son petit nez crochu..." + Reset)
	fmt.Println()
	fmt.Printf(Cyan+"%s (en poussant un cri aigu) : « UN VÊTEMENT ! »\n"+Reset, nomElfe)
	fmt.Printf(Cyan+"« Le bon jeune maître a donné un vêtement à %s ! %s EST LIBRE ! »\n"+Reset, nomElfe, strings.ToUpper(nomElfe))
	fmt.Println()

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour voir la réaction de l'elfe...]" + Reset)
	ClearTerminal()

	fmt.Println("\n" + Yellow + "=== L'ÉCHANGE DES ELFES ===" + Reset)
	fmt.Printf("%s serre le vêtement contre son cœur, saute sur place avec enthousiasme,\n", nomElfe)
	fmt.Println("puis disparaît soudainement dans un claquement sec (*CRAC*) !")
	fmt.Println()
	fmt.Println("Un instant plus tard, l'elfe réapparaît avec un présent...")
	fmt.Println()

	// Récompense (80% nourriture / 20% info)
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)

	if chance < 80 {
		nourritures := []struct {
			Nom string
			PV  int
		}{
			{Nom: "Tarte aux mélasses et crème anglaise (+25 PV)", PV: 25},
			{Nom: "Grand verre de Bièrobeurre bien moussante (+20 PV)", PV: 20},
			{Nom: "Boîte de Chocogrenouilles des cuisines (+15 PV)", PV: 15},
			{Nom: "Plateau de Pâtisseries dorées fait maison (+30 PV)", PV: 30},
			{Nom: "Jus de Potiron frappé des elfes (+15 PV)", PV: 15},
		}

		platGagne := nourritures[rand.Intn(len(nourritures))]
		p.AjouterItem(platGagne.Nom, TypeConsommable, 2, 1)

		fmt.Println(Green + "==========================================================================" + Reset)
		fmt.Printf(Green+"[RÉCOMPENSE DE L'ELFE] %s vous tend fièrement : %s !\n"+Reset, nomElfe, platGagne.Nom)
		fmt.Printf(Green+"%s : « Prenez ceci, jeune maître ! C'est le meilleur plat des cuisines ! »\n"+Reset, nomElfe)
		fmt.Println(Green + "(L'objet a été ajouté à votre sac à dos)" + Reset)
		fmt.Println(Green + "==========================================================================" + Reset)

	} else {
		infosSecretes := []string{
			"« Psst ! Si vous chatouillez la poire sur le tableau de la nature morte, la porte des cuisines s'ouvre toute seule ! »",
			"« L'elfe sait que le Professeur Rogue cache des ingrédients rares derrière la troisième brique à gauche de son bureau ! »",
			"« Si vous cherchez la Salle sur Demande, marchez trois fois devant la tapisserie de Barnabas le Follet au 7ème étage... »",
			"« Le Peeves l'Esprit Frappeur a très peur du Baron Sanglant ! Citez son nom s'il vous embête dans les couloirs ! »",
			"« Les clés volantes de la réserve ont une aile cassée, c'est la plus ancienne qu'il faut attraper ! »",
		}

		secretGagne := infosSecretes[rand.Intn(len(infosSecretes))]

		fmt.Println(Yellow + "==========================================================================" + Reset)
		fmt.Printf(Yellow+"[SECRET RÉVÉLÉ] %s se penche à votre oreille et murmure :\n"+Reset, nomElfe)
		fmt.Println(Cyan + secretGagne + Reset)
		fmt.Println(Yellow + "==========================================================================" + Reset)
	}

	fmt.Println()
	LireSaisie(reader, DarkGray+"Appuyez sur Entrée pour continuer à explorer les cuisines..." + Reset)
	ClearTerminal()
}