package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Coursbotanique gère l'accès au cours de botanique du Professeur Chourave dans les serres
func Coursbotanique() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		// Calcul de l'année en fonction du niveau (1-10 = 1ère année, 11-20 = 2ème année, etc.)
		annee := (p.Levelbotanique-1)/10 + 1
		if annee > 7 {
			annee = 7
		}

		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Println(Green + "               SERRES DE BOTANIQUE DE POUDLARD            " + Reset)
		fmt.Println(Green + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp botanique : %d | Gallions : %d | Niveau botanique : %d (Année %d)\n", p.Health, p.MaxHealth, p.Expbotanique, p.Gallions, p.Levelbotanique, annee)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Une douce odeur de terre humide, de compost et de plantes exotiques")
		fmt.Println("flotte dans l'air chaud des serres de verre. Le Professeur Chourave")
		fmt.Println("s'affaire autour de bacs de terreau armée de gros gants de protection.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Écouter les conseils du Professeur Chourave (Dialogue)")
		fmt.Println(" 2 - 1ère Année : Rempotage des Mandragores (Niveau 1+)")
		fmt.Println(" 3 - 2ème Année : Récolte du Bulbe Enflammé (Niveau 11+)")
		fmt.Println(" 4 - 3ème Année : Entretien du Mimbulus Mimbletonia (Niveau 21+)")
		fmt.Println(" 5 - 5ème Année : Manipulation des Filets du Diable (Niveau 41+) [Dangereux]")
		fmt.Println(" 6 - Retourner au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Green + "==========================================================" + Reset)

		Option := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Dans quelle salle de cours voulez-vous vous rendre ? : "+Reset)))

		switch Option {
		case "1":
			ClearTerminal()
			DialoguesChourave()
		case "2":
			ClearTerminal()
			TenterMandragores(p)
		case "3":
			ClearTerminal()
			TenterBulbeEnflamme(p)
		case "4":
			ClearTerminal()
			TenterMimbulus(p)
		case "5":
			ClearTerminal()
			TenterFiletsDuDiable(p)
		case "6":
			ClearTerminal()
			fmt.Println(Yellow + "Vous saluez le Professeur Chourave et quittez les serres..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard()
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Coursbotanique)
			return
		default:
			ClearTerminal()
			Halldepoudlard()
		}
	}
}

// DialoguesChourave gère les encouragements bienveillants de la professeure
func DialoguesChourave() {
	reader := Lecteur
	dialogues := []string{
		"« N'oubliez jamais de porter vos gants en cuir de dragon dans la serre numéro 3 ! La sécurité avant tout, mes petits choux. »",
		"« Les plantes ont des sentiments, elles ressentent la peur comme la joie. Il faut leur parler avec douceur. »",
		"« Un bon botaniste sait écouter la terre autant que les feuilles. Prenez votre temps. »",
	}
	phrase := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Green + "=== PROFESSEUR POMONA CHOURAVE ===" + Reset)
	fmt.Println("Chourave essuie un peu de terre sur son front et vous sourit chaleureusement.")
	fmt.Println()
	fmt.Printf(Cyan+"Chourave : %s\n"+Reset, phrase)
	fmt.Println()

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour reprendre le jardinage]"+Reset)
}

// --- NIVEAU 1 : REMPOTAGE DES MANDRAGORES ---

func TenterMandragores(p *Player) {
	reader := Lecteur
	if p.Levelbotanique < 1 {
		fmt.Println(Red + "Niveau insuffisant (Niveau 1 requis)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Gants protecteurs en cuir de dragon (ou terreau spécial / engrais - Coût chez l'Apothicaire : 2 Gallions)
	if !VerifierEtProposerAchat(p, "Engrais magique pour Mandragores", 2) {
		return
	}

	fmt.Println(Cyan + "=== COURS 1 : REMPOTAGE DES MANDRAGORES ===" + Reset)
	fmt.Println("La jeune mandragore crie de toutes ses forces dans son pot ! Vous devez faire vite.")
	fmt.Println("1 - Agripper fermement les feuilles d'une main ferme et enfouir d'un coup sec")
	fmt.Println("2 - Hésiter, lâcher les feuilles et boucher vos oreilles en panique")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Bien joué ! La mandragore est rempotée sans un cri de trop. Chourave vous félicite (+10 Expbotanique) !" + Reset)
		p.Expbotanique += 10
		p.AjouterItem("Racine de Mandragore séchée", TypeConsommable, 5, 1)
	} else {
		fmt.Println(Red + "Son cri perçant vous étourdit et vous assomme un instant (-5 PV)." + Reset)
		p.Health -= 5
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 2 : RÉCOLTE DU BULBE ENFLAMMÉ ---

func TenterBulbeEnflamme(p *Player) {
	reader := Lecteur
	if p.Levelbotanique < 11 {
		fmt.Println(Red + "Niveau insuffisant ! Vous devez être en 2ème année (Niveau 11 minimum)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Gants en peau de Moke (Coût chez l'Apothicaire : 5 Gallions)
	if !VerifierEtProposerAchat(p, "Gants en peau de Moke", 5) {
		return
	}

	fmt.Println(Cyan + "=== COURS 2 : RÉCOLTE DU BULBE ENFLAMMÉ ===" + Reset)
	fmt.Println("Le bulbe orange vif palpite et dégage une chaleur intense. Comment récupérez-vous son jus ?")
	fmt.Println("1 - Utiliser délicatement le couteau à lame d'argent pour percer la coque")
	fmt.Println("2 - Tirer dessus à mains nues en espérant que ça tienne")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Précision chirurgicale ! Le jus incandescent s'écoule dans votre fiole (+20 Expbotanique) !" + Reset)
		p.Expbotanique += 20
		p.AjouterItem("Jus de Bulbe Enflammé", TypeConsommable, 8, 1)
	} else {
		fmt.Println(Red + "BOUM ! Le bulbe éclate en gerbes d'étincelles brûlantes sur vos doigts (-10 PV)." + Reset)
		p.Health -= 10
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 3 : ENTRETIEN DU MIMBULUS MIMBLETONIA ---

func TenterMimbulus(p *Player) {
	reader := Lecteur
	if p.Levelbotanique < 21 {
		fmt.Println(Red + "Niveau insuffisant ! Vous devez être en 3ème année (Niveau 21 minimum)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Sachet d'Aconit (Coût chez l'Apothicaire : 3 Gallions)
	if !VerifierEtProposerAchat(p, "Sachet d'Aconit", 3) {
		return
	}

	fmt.Println(Cyan + "=== COURS 3 : ENTRETIEN DU MIMBULUS MIMBLETONIA ===" + Reset)
	fmt.Println("La plante cactus-like frémit. Quel soin lui prodiguez-vous pour nourrir ses furoncles de pus ?")
	fmt.Println("1 - Appliquer l'engrais d'Aconit en massant délicatement les pustules")
	fmt.Println("2 - Arroser abondamment d'eau glacée d'un coup sec")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Le Mimbulus ronronne presque de plaisir et dégage un parfum agréable (+35 Expbotanique) !" + Reset)
		p.Expbotanique += 35
		p.AjouterItem("Échantillon de Pus de Mimbulus", TypeConsommable, 12, 1)
	} else {
		fmt.Println(Red + "SPOUTCH ! La plante panique et vous asperge entièrement de pus d'empestine nauséabond (-15 PV)." + Reset)
		p.Health -= 15
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

// --- NIVEAU 5 : MANIPULATION DES FILETS DU DIABLE ---

func TenterFiletsDuDiable(p *Player) {
	reader := Lecteur
	if p.Levelbotanique < 41 {
		fmt.Println(Red + "Niveau insuffisant ! Cette plante de 5ème année exige le niveau 41 minimum." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	// Ingrédient requis : Potion de Soin ou Torche magique (Coût chez l'Apothicaire : 10 Gallions)
	if !VerifierEtProposerAchat(p, "Kit anti-constriction magique", 10) {
		return
	}

	fmt.Println(Red + "=== COURS AVANCÉ : LES FILETS DU DIABLE ===" + Reset)
	fmt.Println("Les lianes sombres vous enserrent déjà les poignets et tentent de vous étouffer ! Que faites-vous ?")
	fmt.Println("1 - Rester parfaitement calme et détendre ses muscles pour faire reculer les lianes")
	fmt.Println("2 - Paniquer et tirer de toutes ses forces sur les lianes épineuses")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(2 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Maîtrise absolue ! En cessant de lutter, les lianes se relâchent et vous libèrent en s'inclinant (+60 Expbotanique)." + Reset)
		p.Expbotanique += 60
		p.AjouterItem("Fibre de Filet du Diable rare", TypeDivers, 30, 1)
	} else {
		fmt.Println(Red + "Erreur fatale ! Plus vous luttez, plus les lianes se resserrent violemment autour de vous (-25 PV)." + Reset)
		p.Health -= 25
		if p.Health < 1 {
			p.Health = 1
		}
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour reprendre votre souffle]"+Reset)
}