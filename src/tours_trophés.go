package src

import (
	"fmt"
	"strings"
	"time"
)

func Tourstrophes() {
	reader := Lecteur
	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Admirer les trophés ")
	fmt.Println()
	fmt.Println(" 2 - Aller voir les Sabliers ")
	fmt.Println()
	fmt.Println(" Q - Retourner au Hall ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := LireSaisie(reader, DarkGray+"Dans quelle salle de cours voulez-vous vous rendre ? : "+Reset)

	switch Option {
	case "1" : 
		ClearTerminal()
		SalleDesTrophees()
	case "2" :
		ClearTerminal()
		AfficherSabliersMaisons()
	case "Q":
		ClearTerminal()
		Halldepoudlard()
	case "inv":
		ClearTerminal()
		AfficherInventaire(Tourstrophes)
	default:
		ClearTerminal()
		Halldepoudlard()
	}
}

// SalleDesTrophees gère la visite de la salle des trophées et du mémorial de Poudlard.
func SalleDesTrophees() {
	p := JoueurActuel
	reader := Lecteur

	for {
		ClearTerminal()

		// ASCII Art monumental de la Galerie des Trophées
fmt.Println(Yellow + `
  ____________________________________________________________________  
 /   ________________________________________________________________ \
|  |                                                                |  |
|  |` + Reset + Red + `               LA SALLE DES TROPHÉES DE POUDLARD              ` + Reset + Yellow + `  |  |
|  |` + Reset + Cyan + `                     ~ Éternelle Mémoire ~                    ` + Reset + Yellow + `  |  |
|  |________________________________________________________________|  |
 \____________________________________________________________________/ 
` + Reset)

		fmt.Println(Cyan + `
               .-.                                                      .-.
              (   )                                                    (   )
               |=|                                                      |=|
              /   \                                                    /   \
          .-. |   | .-.                                            .-. |   | .-.
         (   )|   |(   )          .----------------.              (   )|   |(   )
          |=| |   | |=|          /  _     _     _   \              |=| |   | |=|
         /   \|   |/   \        |  (_)   (_)   (_)  |             /   \|   |/   \
        |_____|___|_____|       |   .-----------.   |            |_____|___|_____|
        |               |       |  /  ___   ___  \  |            |               |
        |  COUPE DES    |       | |  /   \ /   \  | |            |  COUPE DU     |
        | 4 MAISONS     |       | |  |   | |   |  | |            |  TOURNOI DES  |
        |               |       | |  \___/ \___/  | |            | 3 SORCIERS    |
        |_______________|       |  \_____________/  |            |_______________|
          |           |          \_________________/               |           |
         _|___________|_          |               |               _|___________|_
        [_______________]         | TOURNOI 1994  |              [_______________]
                                  |_______________|
                                    |           |
                                   _|___________|_
                                  [_______________]
` + Reset)

		fmt.Println(Yellow + `
  _____________________________________________________________________________________________________
 |                                                                                                     |
 | [1] Coupe des Quatre Maisons                    [5] Mémorial des Élèves Tombés pour Poudlard        |
 | [2] Trophée du Tournoi des Trois Sorciers       [6] Plaque de la Reconstruction & de l'Unité        |
 | [3] Trophée de Quidditch & Armures Gravées      [7] Médaille de la Paix Inter-Sorcière              |
 | [4] Services Rendus à l'École (Harry & Tom)     [Q] Revenir au Hall de Poudlard                     |
 |_____________________________________________________________________________________________________|` + Reset)

		fmt.Printf("\nSorcier en visite : %s (%s) | Pression de l'air : Électrique et solennelle\n\n", p.Name, p.Maison)

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Choisissez un trophée ou une plaque à examiner (1-7 ou 'Q') : "+Reset)))

		switch choix {
		case "1":
			ClearTerminal()
			fmt.Println(Yellow + "=== LA COUPE DES QUATRE MAISONS ===" + Reset)
			fmt.Println(`
       .-''''-.
      /  ____  \
     |  /    \  |
     | |      | |   "Accordée chaque fin d'année à la Maison ayant
     |  \____/  |    accumulé le plus de sabliers de rubis, d'émeraudes,
      \        /     de saphirs et de diamants."
       '-....-'
          ||
        __||__
       [______]`)
			fmt.Println("Les gravures récentes montrent une alternance féroce entre Gryffondor, Serpentard, Serdaigle et Poufsouffle.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "2":
			ClearTerminal()
			fmt.Println(Cyan + "=== LE TROPHÉE DU TOURNOI DES TROIS SORCIERS ===" + Reset)
			fmt.Println(`
          /\
         /  \
        / /\ \
       / /  \ \       "Une coupe en argent massif étincelante,
      / / /\ \ \      taillée comme du cristal brut.
     / / /  \ \ \     Elle rappelle la victoire tragique de 1994-1995
    /_/_/____\_\_\    et le souvenir immortel de Cedric Diggory."
        |    |
        |____|
       (______)`)
			fmt.Println("En posant la main près du socle, vous ressentez une légère fraîcheur magique, trace du Portoloin d'autrefois.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "3":
			ClearTerminal()
			fmt.Println(Red + "=== LES TROPHÉES DE QUIDDITCH ET ARMURES gravées ===" + Reset)
			fmt.Println(`
       ======
     /  ____  \      "Sur l'un des écussons dorés, le nom de James Potter est gravé
    |  /    \  |     comme Attrapeur d'exception de Gryffondor.
     \ \____/ /      À côté repose le blason de Minerva McGonagall,
      '-....-'       joueuse redoutable des années 1950."`)
			fmt.Println("Les armures alignées le long du mur s'inclinent légèrement sur votre passage dans un grincement de métal poli.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "4":
			ClearTerminal()
			fmt.Println(Yellow + "=== LES DISTINCTIONS POUR SERVICES RENDUS À L'ÉCOLE ===" + Reset)
			fmt.Println(`
   ____________________________________________________________________
  |                                                                    |
  |  - HARRY JAMES POTTER ET RONALD BILLIUS WEASLEY (1993)             |
  |    Pour avoir sauvé la Chambre des Secrets et terrasse le Monstre. |
  |                                                                    |
  |  - TOM JEDUSOR (1943) [Plaque altérée et partiellement masquée]    |
  |    "Distinction attribuée pour fausse dénonciation..."             |
  |____________________________________________________________________|`)
			fmt.Println("La plaque de Tom Jedusor a perdu son éclat doré et semble couverte d'une ombre terne ineffaçable.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "5":
			ClearTerminal()
			fmt.Println(Green + "=== MÉMORIAL ÉTERNEL DES ÉLÈVES TOMBÉS POUR POUDLAARD ===" + Reset)
			fmt.Println(Yellow + `
                      🕯️   🕊️   🕯️
            .-------------------------------.
           /   IN MEMORIAM - 2 MAI 1998      \
          |                                   |
          |  "Aux courageux élèves qui ont    |
          |   dressé leurs baguettes contre   |
          |   les ténèbres pour protéger      |
          |   Poudlard et l'avenir de la      |
          |   magie. Leurs noms brillent      |
          |   à jamais dans ces pierres."     |
          |                                   |
          |  - Fred Weasley                   |
          |  - Lavender Brown                 |
          |  - Colin Creevey                  |
          |  ... et tous les héros anonymes   |
          |     des Quatre Maisons.           |
           \_________________________________/` + Reset)
			fmt.Println(Cyan + "Une lumière douce et dorée émane continuellement du marbre blanc." + Reset)
			fmt.Println("Des fleurs fraîches de lys et de pensées apparaissent magiquement chaque matin au pied du monument.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour vous recueillir et continuer]"+Reset)

		case "6":
			ClearTerminal()
			fmt.Println(Yellow + "=== TROPHÉE DE LA RECONSTRUCTION ET DE L'UNITÉ ===" + Reset)
			fmt.Println(`
             /=========================\
            |   L'UNION DES QUATRE      |
            |                           |
            |  "Décerné aux volontaires,|
            |   élèves, professeurs et  |
            |   créatures magiques qui  |
            |   ont rebâti les tours    |
            |   et les remparts de      |
            |   Poudlard pierre par     |
            |   pierre après la bataille|
             \=========================/`)
			fmt.Println("Les blasons des quatre maisons sont fondus ensemble dans un même morceau de bronze enchanteur.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "7":
			ClearTerminal()
			fmt.Println(Cyan + "=== LA MÉDAILLE DE LA PAIX INTER-SORCIÈRE ===" + Reset)
			fmt.Println(`
               .---.
              /  _  \
             |  / \  |
             |  \_/  |    "Forcée en l'honneur du retour de la Paix
              \     /     et de la réconciliation entre le Ministère,
               '---'      Poudlard et les créatures du monde magique."
                 ||
              .-'||'-.
             '--------'`)
			fmt.Println("Cette médaille commémore l'abolition des lois discriminatoires envers les nés-mouldus et les êtres magiques.")
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "Q":
			ClearTerminal()
			fmt.Println(Yellow + "Vous jetez un dernier regard solennel aux trophées avant de repasser les lourdes portes..." + Reset)
			time.Sleep(1 * time.Second)
			Tourstrophes()
			return

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}



// AfficherSabliersMaisons affiche les sabliers des points des quatre maisons de Poudlard
func AfficherSabliersMaisons() {
	p := JoueurActuel
	reader := Lecteur

	ClearTerminal()

	fmt.Println(Yellow + "==========================================================================" + Reset)
	fmt.Println(Yellow + "            LES SABLIERS DES POINTS DES QUATRE MAISONS                    " + Reset)
	fmt.Println(Yellow + "==========================================================================" + Reset)
	fmt.Println()
	fmt.Println("Face aux fameux trophés des élèves de Poudlard se dresse une imposante rangée de sabliers")
	fmt.Println("en verre sertis d'or. Chaque grain de pierre précieuse représente les points")
	fmt.Println("gagnés ou perdus par les élèves au cours de l'année scolaire.")
	fmt.Println()

	// Entêtes colorées des quatre maisons
	fmt.Printf("   %s%-15s%s   %s%-15s%s   %s%-15s%s   %s%-15s%s\n",
		Red, "GRYFFONDOR", Reset,
		Green, "SERPENTARD", Reset,
		Cyan, "SERDAIGLE", Reset,
		Yellow, "POUFSOUFFLE", Reset)

	fmt.Printf("   %s%-15s%s   %s%-15s%s   %s%-15s%s   %s%-15s%s\n\n",
		Red, "(Rubis)", Reset,
		Green, "(Émeraudes)", Reset,
		Cyan, "(Saphirs)", Reset,
		Yellow, "(Diamants)", Reset)

	// Motifs des sabliers (72 caractères de large = 4 blocs de 18 caractères)
	lignes := []string{
		"     .-------.         .-------.         .-------.         .-------.    ",
		"    /         \\       /         \\       /         \\       /         \\   ",
		"   |   o   o   |     |   o   o   |     |   o   o   |     |   o   o   |  ",
		"    \\    o    /       \\    o    /       \\    o    /       \\    o    /   ",
		"     )   o   (         )   o   (         )   o   (         )   o   (    ",
		"    /  o   o  \\       /  o   o  \\       /  o   o  \\       /  o   o  \\   ",
		"   |  o o o o  |     |  o o o o  |     |  o o o o  |     |  o o o o  |  ",
		"    '---------'       '---------'       '---------'       '---------'   ",
	}

	// Affichage ligne par ligne en colorant chaque sablier séparément
	for _, l := range lignes {
		p1 := l[0:18]
		p2 := l[18:36]
		p3 := l[36:54]
		p4 := l[54:]

		fmt.Printf("%s%s%s%s%s%s%s%s\n",
			Red, p1,
			Green, p2,
			Cyan, p3,
			Yellow, p4+Reset)
	}

	fmt.Println()
	fmt.Println(DarkGray + "--------------------------------------------------------------------------" + Reset)
	fmt.Printf("Votre maison actuelle : %s%s%s\n\n", Yellow, p.Maison, Reset)

	LireSaisie(reader, DarkGray+"Appuyez sur Entrée pour revenir au Hall de Poudlard..." + Reset)
	ClearTerminal()
	Tourstrophes()
}