package src

import (
	"fmt"
	"strings"
)

func Chaudronbaveur() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le Chaudron Baveur---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Aller a King's cross ")
	fmt.Println()
	fmt.Println(" 2 - Aller dans le chemin de traverse ")
	fmt.Println()
	fmt.Println(" 3 - Demander à Hagrid son ticket pour le chemin de traverse ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)

	switch Option {
	case "1":
		ClearTerminal()
		Kingscross()
	case "2":
		ClearTerminal()
		Chemindetraverse()
	case "3" :
		ClearTerminal()
		Veriffournitures()
	case "leave":
		ClearTerminal()
		Leave()
	case "fournitures":
		ClearTerminal()
		Printlistefourniture()
		Chaudronbaveur()
	case "inv":
		ClearTerminal()
		AfficherInventaire(Chaudronbaveur)
	default:
		ClearTerminal()
		Chaudronbaveur()
	}
}

func Veriffournitures() bool {
	p := JoueurActuel

	// Compteurs pour les vêtements et le matériel
	nbRobesNoires := 0
	aChapeau := false
	aGants := false
	aCape := false
	aChaudron := false
	aFioles := false
	aBalance := false
	aBaguette := false

	// Compteurs et vérifications des 8 livres obligatoires (auteurs/mots-clés)
	livresObligatoires := map[string]bool{
		"fauconnette": false, // Le Livre des sorts et enchantements (Niveau 1)
		"tourdesac":   false, // Histoire de la magie
		"lasornette":  false, // Magie théorique
		"changé":      false, // Manuel de métamorphose à l'usage des débutants
		"augurolle":   false, // Mille herbes et champignons magiques
		"beaulitron":  false, // Potions magiques
		"dragonneau":  false, // Vie et habitat des animaux fantastiques
		"jentremble":  false, // Forces obscures : comment s'en protéger
	}

	// Compteurs pour la règle : UN chat OU UN hibou/chouette OU UN crapaud
	nbHiboux := 0
	nbChats := 0
	nbCrapauds := 0

	// Parcours complet de l'inventaire du joueur
	for _, item := range p.Inventaire {
		nomLower := strings.ToLower(item.Nom)

		// 1. Uniformes et Équipements
		if strings.Contains(nomLower, "robe de travail (noire)") || strings.Contains(nomLower, "robe de travail noire") {
			nbRobesNoires += item.Quantite
		}
		if strings.Contains(nomLower, "chapeau pointu (noir)") || strings.Contains(nomLower, "chapeau pointu noir") {
			aChapeau = true
		}
		if strings.Contains(nomLower, "gants protecteurs en cuir de dragon") || strings.Contains(nomLower, "gants en cuir de dragon") {
			aGants = true
		}
		if strings.Contains(nomLower, "cape d'hiver avec attaches d'argent") || strings.Contains(nomLower, "cape d'hiver") {
			aCape = true
		}

		// 2. Matériel de classe
		if strings.Contains(nomLower, "chaudron en étain") {
			aChaudron = true
		}
		if strings.Contains(nomLower, "fioles en verre") {
			aFioles = true
		}
		if strings.Contains(nomLower, "balance en cuivre") {
			aBalance = true
		}
		if strings.Contains(nomLower, "baguette") {
			aBaguette = true
		}

		// 3. Vérification des Livres
		for auteur := range livresObligatoires {
			if strings.Contains(nomLower, auteur) {
				livresObligatoires[auteur] = true
			}
		}

		// 4. Décompte des animaux autorisés
		if strings.Contains(nomLower, "hibou") || strings.Contains(nomLower, "chouette") || strings.Contains(nomLower, "petit duc") {
			nbHiboux += item.Quantite
		}
		if strings.Contains(nomLower, "chat") {
			nbChats += item.Quantite
		}
		if strings.Contains(nomLower, "crapaud") {
			nbCrapauds += item.Quantite
		}
	}

	toutesFournituresPresentes := true

	fmt.Println("\n" + Yellow + "=== HAGRID VÉRIFIE TON SAC DE FOURNITURES ===" + Reset)
	fmt.Println("Hagrid : « Voyons un peu c'que t'as dégotté sur le Chemin de Traverse... »")

	if nbRobesNoires < 3 {
		fmt.Printf(Red+"[X] « T'as seulement %d/3 robes noires ! Tu vas pas te balader en chemise à Poudlard, quand même ! »\n"+Reset, nbRobesNoires)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Bon, t'as tes trois robes noires. C'est déjà ça ! »" + Reset)
	}

	if !aChapeau {
		fmt.Println(Red + "[X] « Il te manque le chapeau pointu noir ! C'est obligatoire pour les cérémonies ! »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Ah, le chapeau pointu est là. M'en parle pas, j'entre jamais la tête dedans moi... »" + Reset)
	}

	if !aGants {
		fmt.Println(Red + "[X] « T'as oublié les gants en cuir de dragon ! Tu veux te faire griller les doigts en Botanique ? »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Des gants en cuir de dragon, du solide ça ! Pratique pour manipuler des créatures un peu nerveuses. »" + Reset)
	}

	if !aCape {
		fmt.Println(Red + "[X] « T'as pas pris la cape d'hiver avec les attaches d'argent ? Ça caille sévère autour du lac en décembre ! »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « La cape d'hiver est bien au chaud dans ton sac. »" + Reset)
	}

	if !aChaudron {
		fmt.Println(Red + "[X] « T'as pas de chaudron en étain ? Le professeur Rogue va pas te rater si t'as rien pour mélanger tes baves de crapaud ! »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Chaudron en étain taille 2, parfait pour les cours de Potions. »" + Reset)
	}

	if !aFioles {
		fmt.Println(Red + "[X] « Il te manque les fioles en verre ! Tu comptais mettre tes échantillons dans tes poches ? »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Tes fioles en verre sont bien emballées. »" + Reset)
	}

	if !aBalance {
		fmt.Println(Red + "[X] « T'as pas pris la balance en cuivre ! Tu vas rater tous tes dosages de poudre de scarabée ! »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Balance en cuivre, nikel. »" + Reset)
	}

	if !aBaguette {
		fmt.Println(Red + "[X] « Attends un peu... T'es pas allé chez Ollivander ?! Tu peux pas aller à Poudlard sans baguette magique, gamin ! »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Une vraie baguette de chez Ollivander ! Y'a rien de tel ! »" + Reset)
	}

	// Vérification des manuels scolaires
	livresManquants := 0
	for _, possede := range livresObligatoires {
		if !possede {
			livresManquants++
		}
	}
	if livresManquants > 0 {
		fmt.Printf(Red+"[X] « Oula, il te manque encore %d livre(s) de la liste ! File chez Fleury et Bott rattraper ça ! »\n"+Reset, livresManquants)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Tous tes bouquins de première année sont là ! Sacrément lourd ce sac, hein ? »" + Reset)
	}

	// Vérification de la règle de l'animal autorisé
	typesAnimauxPossedes := 0
	if nbHiboux > 0 {
		typesAnimauxPossedes++
	}
	if nbChats > 0 {
		typesAnimauxPossedes++
	}
	if nbCrapauds > 0 {
		typesAnimauxPossedes++
	}

	totalAnimauxValides := nbHiboux + nbChats + nbCrapauds

	if typesAnimauxPossedes == 0 {
		fmt.Println(Red + "[X] « T'as pas pris de compagnon ? Il te faut soit UN hibou, soit UN chat, soit UN crapaud. Va faire un tour à la Ménagerie Magique ! »" + Reset)
		toutesFournituresPresentes = false
	} else if typesAnimauxPossedes > 1 || totalAnimauxValides > 1 {
		fmt.Println(Red + "[X] « Holà ! Tu te crois dans ma cabane ou quoi ?! C'est UN SEUL animal autorisé : soit un chat, soit un hibou, soit un crapaud. Pas toute une ménagerie ! »" + Reset)
		toutesFournituresPresentes = false
	} else {
		fmt.Println(Green + "[✓] « Un bon petit compagnon pour le voyage. Dumbledore sera content ! »" + Reset)
	}

	fmt.Println("\n-------------------------------------------------------------------")

	if toutesFournituresPresentes {
		fmt.Println(Green + "Hagrid (en tapant dans le dos du joueur) : « Par les barbes de Merlin, t'as absolument tout ! T'es paré pour l'aventure ! »" + Reset)

		// Attribuer le billet si le joueur ne l'a pas déjà
		dejaTicket := false
		for _, item := range p.Inventaire {
			if strings.Contains(strings.ToLower(item.Nom), "ticket") && strings.Contains(strings.ToLower(item.Nom), "9") {
				dejaTicket = true
				break
			}
		}

		if !dejaTicket {
			p.AjouterItem("Ticket pour le Poudlard Express (Voie 9 ¾)", TypeDivers, 0, 1)
			fmt.Println(Yellow + "\nHagrid fouille dans sa grande veste en peau de taupe et te tend un billet doré :" + Reset)
			fmt.Println(Green + "« Tiens ! Voilà ton ticket pour le train. C'est à la gare de King's Cross, voie 9 ¾. Tu perds pas ça, d'accord ? On se voit à Poudlard ! »" + Reset)
		} else {
			fmt.Println(Yellow + "Hagrid : « T'as déjà ton billet pour la voie 9 ¾ bien au chaud dans tes poches ! N'oublie pas de prendre le train à King's Cross. »" + Reset)
		}
	} else {
		fmt.Println(Red + "Hagrid : « Allons, allons ! Il te manque encore du matériel. Fais vite le tour des magasins avant qu'on rate le train ! »" + Reset)
	}

	reader := Lecteur
    LireSaisie(reader, "\nAppuyez sur Entrée pour continuer...")
    ClearTerminal()
	Chaudronbaveur()

    return toutesFournituresPresentes
}
