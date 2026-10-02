package src

import (
	"fmt"
	"strings"
	"time"
)

func Lessallescommunes() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Red + "---------------Le grand hall---------------" + Reset)
	fmt.Println()
	fmt.Println(" 1 - Salle commune de Gryffondor ")
	fmt.Println()
	fmt.Println(" 2 - Salle commune de Serpentard ")
	fmt.Println()
	fmt.Println(" 3 - Salle commune de Serdaigle ")
	fmt.Println()
	fmt.Println(" 4 - Salle commune de Poufsouffle ")
	fmt.Println()
	fmt.Println(Red + "-------------------------------------" + Reset)

	Option := strings.TrimSpace(strings.ToLower(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

	switch Option {
	case "1":
		ClearTerminal()
		if p.Maison == "Gryffondor" {
			Sallecommunesgryffondor()
		} else {
			fmt.Println(Red + "=== ACCÈS REFUSÉ ===" + Reset)
			fmt.Println("La Grosse Dame du portrait vous dévisage : « Vous n'avez pas le mot de passe et vous n'êtes pas de Gryffondor ! Filez ! »")
			time.Sleep(2 * time.Second)
			ClearTerminal()
			Lessallescommunes()
		}
	case "2":
		ClearTerminal()
		if p.Maison == "Serpentard" {
			Sallecommuneserpentard()
		} else {
			fmt.Println(Red + "=== ACCÈS REFUSÉ ===" + Reset)
			fmt.Println("Le mur de pierre humide reste immobile. Vous entendez des rires moqueurs de l'autre côté.")
			time.Sleep(2 * time.Second)
			ClearTerminal()
			Lessallescommunes()
		}
	case "3":
		ClearTerminal()
		if p.Maison == "Serdaigle" {
			Sallecommunesserdaigle()
		} else {
			fmt.Println(Red + "=== ACCÈS REFUSÉ ===" + Reset)
			fmt.Println("Le heurtoir en forme d'aigle refuse de vous poser une énigme : « Seuls les Serdaigle peuvent tenter leur chance. »")
			time.Sleep(2 * time.Second)
			ClearTerminal()
			Lessallescommunes()
		}
	case "4":
		ClearTerminal()
		if p.Maison == "Poufsouffle" {
			Sallecommunespoufsouffle()
		} else {
			fmt.Println(Red + "=== ACCÈS REFUSÉ ===" + Reset)
			fmt.Println("Vous tapotez le mauvais tonneau et vous vous retrouvez aspergé de vinaigre ! « Accès refusé ! »")
			time.Sleep(2 * time.Second)
			ClearTerminal()
			Lessallescommunes()
		}
	case "inv":
		ClearTerminal()
		AfficherInventaire(Lessallescommunes)
	default:
		ClearTerminal()
		Halldepoudlard()
	}
}