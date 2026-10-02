package src

import (
	"fmt"
	"time"
)

// ExpRequise définit le palier d'expérience nécessaire pour atteindre le niveau suivant.
func ExpRequise(niveau int) int {
	return niveau * 50
}

// ==========================================
// 1. EXPÉRIENCE GLOBALE ET ANNÉE D'ÉTUDES
// ==========================================

// AjouterExperience ajoute de l'XP global au joueur (détermine l'année d'études).
func AjouterExperience(gain int) {
	p := JoueurActuel
	p.Exp += gain
	fmt.Printf(Green+"Vous gagnez %d points d'expérience globale ! (Exp total: %d)\n"+Reset, gain, p.Exp)

	for p.Exp >= ExpRequise(p.Level) {
		p.Exp -= ExpRequise(p.Level)
		LevelUp(p)
	}
}

func LevelUp(p *Player) {
	p.Level++
	
	p.MaxHealth += 10
	p.Health = p.MaxHealth
	p.Atk += 2
	p.Def += 1

	ClearTerminal()
	fmt.Println(Yellow + "==========================================================" + Reset)
	fmt.Printf(Yellow+"   FÉLICITATIONS ! Vous atteignez le niveau global %d !\n"+Reset, p.Level)
	fmt.Println(Yellow + "==========================================================" + Reset)
	fmt.Printf(Green+"Vos statistiques augmentent : Santé Max (+10) | Attaque (+2) | Défense (+1)\n"+Reset)
	
	MettreAJourAnnee(p)
	time.Sleep(3 * time.Second)
}

func MettreAJourAnnee(p *Player) {
	annee := ((p.Level - 1) / 10) + 1 

	switch annee {
	case 1:
		p.Title = "Première année"
	case 2:
		p.Title = "Deuxième année"
	case 3:
		p.Title = "Troisième année"
	case 4:
		p.Title = "Quatrième année"
	case 5:
		p.Title = "Cinquième année (Préparation B.U.S.E)"
	case 6:
		p.Title = "Sixième année"
	case 7:
		p.Title = "Septième année (Préparation A.S.P.I.C)"
	default:
		p.Title = "Sorcier diplômé"
	}
	
	fmt.Printf(Cyan+"Votre rang académique : %s\n"+Reset, p.Title)
}

// ==========================================
// 2. EXPÉRIENCE PAR MATIÈRE
// ==========================================

// AjouterExperienceBotanique gère la progression spécifique au cours de Chourave
func AjouterExperienceBotanique(gain int) {
	p := JoueurActuel
	p.Expbotanique += gain
	fmt.Printf(Green+"Vous gagnez %d points d'expérience en Botanique !\n"+Reset, gain)

	for p.Expbotanique >= ExpRequise(p.Levelbotanique) {
		p.Expbotanique -= ExpRequise(p.Levelbotanique)
		p.Levelbotanique++
		fmt.Printf(Yellow+"Félicitations ! Vous atteignez le niveau %d en Botanique !\n"+Reset, p.Levelbotanique)
		time.Sleep(1 * time.Second)
	}
}

// AjouterExperienceDivination gère la progression spécifique au cours de Trelawney
func AjouterExperienceDivination(gain int) {
	p := JoueurActuel
	p.Expdivination += gain
	fmt.Printf(Green+"Vous gagnez %d points d'expérience en Divination !\n"+Reset, gain)

	for p.Expdivination >= ExpRequise(p.Leveldivination) {
		p.Expdivination -= ExpRequise(p.Leveldivination)
		p.Leveldivination++
		fmt.Printf(Yellow+"Félicitations ! Vous atteignez le niveau %d en Divination !\n"+Reset, p.Leveldivination)
		time.Sleep(1 * time.Second)
	}
}

// AjouterExperienceMetamorphose gère la progression spécifique au cours de McGonagall
func AjouterExperienceMetamorphose(gain int) {
	p := JoueurActuel
	p.Expmetamorphose += gain
	fmt.Printf(Green+"Vous gagnez %d points d'expérience en Métamorphose !\n"+Reset, gain)

	for p.Expmetamorphose >= ExpRequise(p.Levelmetamorphose) {
		p.Expmetamorphose -= ExpRequise(p.Levelmetamorphose)
		p.Levelmetamorphose++
		fmt.Printf(Yellow+"Félicitations ! Vous atteignez le niveau %d en Métamorphose !\n"+Reset, p.Levelmetamorphose)
		time.Sleep(1 * time.Second)
	}
}

// AjouterExperiencePotions gère la progression spécifique au cours de Rogue
func AjouterExperiencePotions(gain int) {
	p := JoueurActuel
	p.Exppotions += gain
	fmt.Printf(Green+"Vous gagnez %d points d'expérience en Potions !\n"+Reset, gain)

	for p.Exppotions >= ExpRequise(p.Levelpotions) {
		p.Exppotions -= ExpRequise(p.Levelpotions)
		p.Levelpotions++
		fmt.Printf(Yellow+"Félicitations ! Vous atteignez le niveau %d en Potions !\n"+Reset, p.Levelpotions)
		time.Sleep(1 * time.Second)
	}
}

// AjouterExperienceSortilege gère la progression spécifique au cours de Flitwick
func AjouterExperienceSortilege(gain int) {
	p := JoueurActuel
	p.Expsortilege += gain
	fmt.Printf(Green+"Vous gagnez %d points d'expérience en Sortilèges !\n"+Reset, gain)

	for p.Expsortilege >= ExpRequise(p.Levelsortilege) {
		p.Expsortilege -= ExpRequise(p.Levelsortilege)
		p.Levelsortilege++
		fmt.Printf(Yellow+"Félicitations ! Vous atteignez le niveau %d en Sortilèges !\n"+Reset, p.Levelsortilege)
		time.Sleep(1 * time.Second)
	}
}