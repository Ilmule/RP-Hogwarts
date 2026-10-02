package src

import (
	"fmt"
	"strings"
	"time"
)

func AfficherInventaire(retour func()) {
	p := JoueurActuel
	reader := Lecteur

	for {
		ClearTerminal()
		fmt.Println(Yellow + "==========================================================" + Reset)
		fmt.Println(Yellow + "                     SAC À DOS                            " + Reset)
		fmt.Println(Yellow + "==========================================================" + Reset)
		fmt.Printf("Joueur : %s | Niveau : %d | Année : %s\n", p.Name, p.Level, p.Title)
		AfficherBourse(*p)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Équipement actuel :")
		fmt.Printf(" - Baguette : %s\n", p.BaguetteEquipee)
		fmt.Printf(" - Balai    : %s\n", p.BalaiEquipe)
		fmt.Printf(" - Robe     : %s\n", p.RobeEquipee)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Objets dans le sac :")

		if len(p.Inventaire) == 0 {
			fmt.Println("Votre sac est complètement vide.")
		} else {
			for i, item := range p.Inventaire {
				fmt.Printf(" %d - %s (x%d) [%s]\n", i+1, item.Nom, item.Quantite, item.Type)
			}
		}

		fmt.Println("----------------------------------------------------------")
		fmt.Println(" Entrez le numéro d'un objet pour interagir avec")
		fmt.Println(" Q - Retourner en arrière")
		fmt.Println()
		fmt.Println()
		fmt.Print("Level botanique : ", p.Levelbotanique)
		fmt.Println()
		fmt.Print("Level divination : ", p.Leveldivination)
		fmt.Println()
		fmt.Print("Level metamorphose : ", p.Levelmetamorphose)
		fmt.Println()
		fmt.Print("Level potions : ", p.Levelpotions)
		fmt.Println()
		fmt.Print("Level sortilege : ", p.Levelsortilege)
		fmt.Println()
		fmt.Println()
		fmt.Println(Yellow + "==========================================================" + Reset)

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "Votre choix : ")))

		if choix == "Q" {
			ClearTerminal()
			if retour != nil {
				retour()
			}
			return
		}

		var idx int
		_, err := fmt.Sscan(choix, &idx)
		if err == nil && idx >= 1 && idx <= len(p.Inventaire) {
			itemChoisi := &p.Inventaire[idx-1]
			ClearTerminal()

			fmt.Printf("=== OBJET : %s ===\n", itemChoisi.Nom)
			fmt.Printf("Type : %s | Quantité : %d\n", itemChoisi.Type, itemChoisi.Quantite)
			fmt.Println("1 - Équiper / Utiliser")
			fmt.Println("2 - Jeter un exemplaire")
			fmt.Println("Q - Retour")

			sousChoix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "Action : ")))

			if sousChoix == "1" {
				switch itemChoisi.Type {
				case TypeEquipement:
					nomLower := strings.ToLower(itemChoisi.Nom)
					if strings.Contains(nomLower, "baguette") {
						p.BaguetteEquipee = itemChoisi.Nom
						fmt.Printf(Green+"Vous avez équipé : %s\n"+Reset, itemChoisi.Nom)
					} else if strings.Contains(nomLower, "balai") || strings.Contains(nomLower, "nimbus") {
						p.BalaiEquipe = itemChoisi.Nom
						fmt.Printf(Green+"Vous avez équipé : %s\n"+Reset, itemChoisi.Nom)
					} else {
						p.RobeEquipee = itemChoisi.Nom
						fmt.Printf(Green+"Vous avez revêtu : %s\n"+Reset, itemChoisi.Nom)
					}
				case TypeConsommable:
					fmt.Printf(Green+"Vous consommez %s. Vos PV augmentent !\n"+Reset, itemChoisi.Nom)
					p.Health += 25
					if p.Health > p.MaxHealth {
						p.Health = p.MaxHealth
					}
					itemChoisi.Quantite--
					if itemChoisi.Quantite <= 0 {
						p.Inventaire = append(p.Inventaire[:idx-1], p.Inventaire[idx:]...)
					}
				default:
					fmt.Println("Cet objet ne peut pas être utilisé activement.")
				}
				time.Sleep(1500 * time.Millisecond)
			} else if sousChoix == "2" {
				itemChoisi.Quantite--
				fmt.Println(Red + "Objet jeté." + Reset)
				if itemChoisi.Quantite <= 0 {
					p.Inventaire = append(p.Inventaire[:idx-1], p.Inventaire[idx:]...)
				}
				time.Sleep(1 * time.Second)
			}
		}
	}
}