package src

import (
	"fmt"
	"strings"
	"time"
)

// CeremonieChoixpeau gère le déroulement de la Répartition des élèves dans la Grande Salle.
func CeremonieChoixpeau() string {
	p := JoueurActuel
	reader := Lecteur

	ClearTerminal()

	// 1. Narration d'introduction
	fmt.Println(Yellow + "==========================================================================" + Reset)
	fmt.Println(Yellow + "               LA CÉRÉMONIE DE LA RÉPARTITION À POUDLARD                  " + Reset)
	fmt.Println(Yellow + "==========================================================================" + Reset)
	fmt.Println()
	fmt.Println("Les grandes portes de la Grande Salle s'ouvrent devant vous.")
	fmt.Println("Des milliers de bougies flottent dans les airs sous un plafond magique étoilé.")
	fmt.Println("Le professeur McGonagall vous avance au centre de la salle, devant un tabouret")
	fmt.Println("sur lequel repose un vieux chapeau rapiécé : le Choixpeau Magique.")
	fmt.Println()
	time.Sleep(2 * time.Second)

	fmt.Println(Cyan + "Le professeur McGonagall : « Quand j'appellerai votre nom, vous vous avancerez ! »" + Reset)
	fmt.Printf(Cyan+"« %s ! »\n\n"+Reset, p.Name)
	time.Sleep(1 * time.Second)

	fmt.Println("Vous vous avancez nerveusement sous le regard de toute l'école et vous vous asseyez.")
	fmt.Println("McGonagall dépose doucement le Choixpeau sur votre tête...")
	fmt.Println()
	time.Sleep(2 * time.Second)

	fmt.Println(Yellow + "Le Choixpeau Magique (murmure à votre oreille) :" + Reset)
	fmt.Println(Yellow + "« Ah... Mmmh... Je vois, je vois. Une tête bien remplie, avec plein de potentiel... »" + Reset)
	fmt.Println(Yellow + "« Mais où vais-je te mettre ? Laisse-moi explorer les recoins de ton esprit... »" + Reset)
	fmt.Println()
	time.Sleep(2 * time.Second)

	// 2. Initialisation des compteurs pour chaque maison
	ptsGryffondor := 0
	ptsSerdaigle := 0
	ptsPoufsouffle := 0
	ptsSerpentard := 0

	// -------------------------------------------------------------------
	// QUESTION 1
	// -------------------------------------------------------------------
	for {
		fmt.Println(DarkGray + "------------------------------------------------------------------" + Reset)
		fmt.Println(Yellow + "Choixpeau Magique : « Face à un obstacle dangereux ou une injustice, quelle est ta première réaction ? »" + Reset)
		fmt.Println(" 1 - Je fonce la tête la première pour défendre les miens, coûte que coûte !")
		fmt.Println(" 2 - J'analyse la situation de manière logique pour trouver le plan parfait.")
		fmt.Println(" 3 - Je m'assure que personne ne soit laissé de côté et j'aide tout le monde.")
		fmt.Println(" 4 - Je cherche la méthode la plus efficace et rusée pour l'utiliser à mon avantage.")

		choix := strings.TrimSpace(LireSaisie(reader, "\nVotre choix (1-4) : "))
		if choix == "1" {
			ptsGryffondor++
			break
		} else if choix == "2" {
			ptsSerdaigle++
			break
		} else if choix == "3" {
			ptsPoufsouffle++
			break
		} else if choix == "4" {
			ptsSerpentard++
			break
		}
		fmt.Println(Red + "« Mmmh... Sois clair dans ta pensée ! Choisis entre 1 et 4. »" + Reset)
	}

	// -------------------------------------------------------------------
	// QUESTION 2
	// -------------------------------------------------------------------
	for {
		fmt.Println(DarkGray + "------------------------------------------------------------------" + Reset)
		fmt.Println(Yellow + "Choixpeau Magique : « Quelle qualité apprécies-tu le plus chez un ami ? »" + Reset)
		fmt.Println(" 1 - La loyauté inébranlable et le sens du travail bien fait.")
		fmt.Println(" 2 - L'ambition, la détermination et le désir de grandeur.")
		fmt.Println(" 3 - La bravoure, l'audace et la loyauté aux idées de justice.")
		fmt.Println(" 4 - La curiosité intellectuelle, l'esprit d'analyse et la sagesse.")

		choix := strings.TrimSpace(LireSaisie(reader, "\nVotre choix (1-4) : "))
		if choix == "1" {
			ptsPoufsouffle++
			break
		} else if choix == "2" {
			ptsSerpentard++
			break
		} else if choix == "3" {
			ptsGryffondor++
			break
		} else if choix == "4" {
			ptsSerdaigle++
			break
		}
		fmt.Println(Red + "« Concentre-toi, la réponse doit être précise ! Choisis entre 1 et 4. »" + Reset)
	}

	// -------------------------------------------------------------------
	// QUESTION 3
	// -------------------------------------------------------------------
	for {
		fmt.Println(DarkGray + "------------------------------------------------------------------" + Reset)
		fmt.Println(Yellow + "Choixpeau Magique : « Tu découvres un grimoire interdit dans la Réserve de la bibliothèque. Que fais-tu ? »" + Reset)
		fmt.Println(" 1 - Je le lis immédiatement en cachette, la connaissance ne devrait pas avoir de limites !")
		fmt.Println(" 2 - Je le garde précieusement, ses connaissances me serviront à devenir puissant.")
		fmt.Println(" 3 - Je le remets au bibliothécaire, respecter les règles préserve la sécurité de tous.")
		fmt.Println(" 4 - Je l'ouvre pour voir s'il y a un défi ou un mystère passionnant à résoudre !")

		choix := strings.TrimSpace(LireSaisie(reader, "\nVotre choix (1-4) : "))
		if choix == "1" {
			ptsSerdaigle++
			break
		} else if choix == "2" {
			ptsSerpentard++
			break
		} else if choix == "3" {
			ptsPoufsouffle++
			break
		} else if choix == "4" {
			ptsGryffondor++
			break
		}
		fmt.Println(Red + "« Réponds clairement, gamin ! Choisis entre 1 et 4. »" + Reset)
	}

	// -------------------------------------------------------------------
	// QUESTION 4
	// -------------------------------------------------------------------
	for {
		fmt.Println(DarkGray + "------------------------------------------------------------------" + Reset)
		fmt.Println(Yellow + "Choixpeau Magique : « Plus tard, quel souvenir aimerais-tu laisser dans l'histoire ? »" + Reset)
		fmt.Println(" 1 - Celui d'un grand sorcier redouté et accompli qui a marqué son époque.")
		fmt.Println(" 2 - Celui d'un héros courageux qui s'est dressé contre le mal.")
		fmt.Println(" 3 - Celui d'un érudit brillant ayant découvert de grands secrets magiques.")
		fmt.Println(" 4 - Celui d'une personne juste, gentille et toujours fidèle à ses proches.")

		choix := strings.TrimSpace(LireSaisie(reader, "\nVotre choix (1-4) : "))
		if choix == "1" {
			ptsSerpentard++
			break
		} else if choix == "2" {
			ptsGryffondor++
			break
		} else if choix == "3" {
			ptsSerdaigle++
			break
		} else if choix == "4" {
			ptsPoufsouffle++
			break
		}
		fmt.Println(Red + "« Ne cherche pas à ruser ! Choisis entre 1 et 4. »" + Reset)
	}

	// 3. Calcul de la maison victorieuse
	maxPoints := ptsGryffondor
	maisonGagnante := "Gryffondor"

	if ptsSerdaigle > maxPoints {
		maxPoints = ptsSerdaigle
		maisonGagnante = "Serdaigle"
	}
	if ptsPoufsouffle > maxPoints {
		maxPoints = ptsPoufsouffle
		maisonGagnante = "Poufsouffle"
	}
	if ptsSerpentard > maxPoints {
		maxPoints = ptsSerpentard
		maisonGagnante = "Serpentard"
	}

	// 4. Gestion des égalités (Choixpeau hésitant)
	maisonsEnEgalite := []string{}
	if ptsGryffondor == maxPoints {
		maisonsEnEgalite = append(maisonsEnEgalite, "Gryffondor")
	}
	if ptsSerdaigle == maxPoints {
		maisonsEnEgalite = append(maisonsEnEgalite, "Serdaigle")
	}
	if ptsPoufsouffle == maxPoints {
		maisonsEnEgalite = append(maisonsEnEgalite, "Poufsouffle")
	}
	if ptsSerpentard == maxPoints {
		maisonsEnEgalite = append(maisonsEnEgalite, "Serpentard")
	}

	// S'il y a plus d'une maison au score maximal, le Choixpeau demande au joueur de départager
	if len(maisonsEnEgalite) > 1 {
		fmt.Println(DarkGray + "------------------------------------------------------------------" + Reset)
		fmt.Println(Yellow + "Choixpeau Magique : « Mmmh... C'est extrêmement difficile... J'hésite grandement entre plusieurs voies pour toi... »" + Reset)
		fmt.Println(Yellow + "« Quelle valeur résonne le plus au fond de ton cœur ? »" + Reset)

		for i, m := range maisonsEnEgalite {
			fmt.Printf(" %d - %s\n", i+1, m)
		}

		for {
			choixDep := strings.TrimSpace(LireSaisie(reader, "\nTranchez vous-même le choix du Choixpeau : "))
			var idx int
			_, err := fmt.Sscan(choixDep, &idx)
			if err == nil && idx >= 1 && idx <= len(maisonsEnEgalite) {
				maisonGagnante = maisonsEnEgalite[idx-1]
				break
			}
			fmt.Println(Red + "Choix invalide. Entrez le numéro de la maison de votre cœur." + Reset)
		}
	}

	// 5. Attribution et Annonce finale
	p.Maison = maisonGagnante

	// Ajustement des statistiques selon la maison
	switch maisonGagnante {
	case "Gryffondor":
		p.Health += 20
		p.MaxHealth += 20
		p.Atk += 2
	case "Serpentard":
		p.Atk += 5
	case "Serdaigle":
		p.Def += 3
		p.Atk += 2
	case "Poufsouffle":
		p.Health += 30
		p.MaxHealth += 30
		p.Def += 2
	}

	fmt.Println("\n" + DarkGray + "------------------------------------------------------------------" + Reset)
	time.Sleep(1 * time.Second)
	fmt.Println(Yellow + "Le Choixpeau s'agite sur votre tête, prend une grande inspiration et crie devant toute la salle :" + Reset)
	time.Sleep(2 * time.Second)

	fmt.Println("\n" + Green + "==========================================================================" + Reset)
	fmt.Printf(Green+"                          « %s ! »\n"+Reset, strings.ToUpper(maisonGagnante))
	fmt.Println(Green + "==========================================================================" + Reset)
	fmt.Println()

	fmt.Printf("La table des %s explose sous les applaudissements et les acclamations !\n", maisonGagnante)
	fmt.Println("Vous retirez le Choixpeau et allez fièrement vous asseoir à la table de votre nouvelle maison.")
	fmt.Println()
	LireSaisie(reader, DarkGray+"Appuyez sur Entrée pour rejoindre le banquet..." + Reset)

	ClearTerminal()
	return maisonGagnante
}