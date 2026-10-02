package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Toursdivination gère le cours de divination dans la tour isolée
func Toursdivination() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		fmt.Println(Red + "==========================================================" + Reset)
		fmt.Println(Red + "               LA SALLE DE DIVINATION                     " + Reset)
		fmt.Println(Red + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Expdivination : %d | Niveau : %d\n", p.Health, p.MaxHealth, p.Expdivination, p.Leveldivination)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("L'atmosphère est étouffante, chauffée à outrance par un grand feu de bois.")
		fmt.Println("La pièce est remplie de petites tables rondes, de coussins moelleux et")
		fmt.Println("d'une odeur entêtante de parfum bon marché qui donne le tournis.")
		fmt.Println("Le Professeur Trelawney émerge de la pénombre, drapée de châles scintillants.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Écouter les sinistres prédictions du Professeur Trelawney")
		fmt.Println(" 2 - Mini-jeu : La Tasse de Thé ")
		fmt.Println(" 3 - Mini-jeu : La Boule de Cristal ")
		fmt.Println(" 4 - Mini-jeu : L'Heure Planétaire ")
		fmt.Println(" 5 - Retourner au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Red + "==========================================================" + Reset)

		Option := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch Option {
		case "1":
			ClearTerminal()
			DialoguesTrelawney()
		case "2":
			ClearTerminal()
			MiniJeuTasseDeThe()
		case "3":
			ClearTerminal()
			MiniJeuBouleDeCristal()
		case "4":
			ClearTerminal()
			MiniJeuAstrologie()
		case "5":
			ClearTerminal()
			fmt.Println(Yellow + "Vous quittez la touffeur de la classe en redescendant l'échelle de corde..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard()
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Toursdivination)
			return
		default:
			ClearTerminal()
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// DialoguesTrelawney gère les discussions énigmatiques et théâtrales avec la professeure
func DialoguesTrelawney() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	dialogues := []string{
		"« Oh... mon pauvre enfant... Je vois autour de votre aura une ombre très, très sombre... Le Ravier ! Oui, le grand Ravier de la fatalité ! »",
		"« Vous êtes né sous l'influence de Saturne, mon cher... Cela indique une résistance aux confitures... ou à la mort subite. L'un ou l'autre. »",
		"« Chaque fois que treize personnes dînent ensemble, la première à se lever est la première à mourir ! Ne l'oubliez jamais ! »",
		"« Votre aura vacille, %s... Je décèle en vous un destin... extraordinaire... ou alors un simple rhume. »",
	}

	phrase := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Red + "=== PROFESSEUR SYBILLE TRELAWNEY ===" + Reset)
	fmt.Println("Trelawney cligne des yeux derrière ses immenses lunettes qui grossissent ses prunelles à l'excès.")
	fmt.Println()
	fmt.Printf(Red+"Trelawney : "+phrase+"\n"+Reset, p.Name)
	fmt.Println()

	// Gain d'expdivinationérience en écoutant le cours
	gainExpdivination := 5
	p.Expdivination += gainExpdivination
	fmt.Printf(Green+"En écoutant ses sombres prophéties, vous gagnez %d points d'expdivinationérience !\n"+Reset, gainExpdivination)

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour vous éloigner de son regard perçant]"+Reset)
}

// MiniJeuTasseDeThe : Interprétation des feuilles de thé au fond de la tasse
func MiniJeuTasseDeThe() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Red + "=== MINI-JEU : LA TASSE DE THÉ ===" + Reset)
	fmt.Println("Vous finissez votre tasse de thé brûlant comme l'exige le cours, puis vous tournez")
	fmt.Println("la tasse trois fois de la main gauche avant de l'observer à l'envers.")
	fmt.Println("Quelle forme croyez-vous distinguer dans les restes de feuilles sombres au fond ?")
	time.Sleep(1 * time.Second)

	fmt.Println("\n 1 - Un Faucon (Indique un ennemi ou un défi imminent)")
	fmt.Println(" 2 - Une Masse d'armes / Un Bâton (Indique une agression ou un coup dur)")
	fmt.Println(" 3 - Un Chien (Traditionnellement... le Ravier, un signe de mort imminente)")
	fmt.Println(" 4 - Une Tasse fendue (Indique un chagrin ou un échec proche)")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre interprétation (1-4) : "))

	fmt.Println(Cyan + "\nVous montrez le fond de votre tasse au Professeur Trelawney..." + Reset)
	time.Sleep(1 * time.Second)

	switch choix {
	case "1":
		fmt.Println(Green + "Trelawney pousse un cri perçant : « Un faucon ! Précisons : un adversaire vous guette ! Vous devez affronter vos peurs ! »" + Reset)
		gain := 12
		p.Expdivination += gain
		fmt.Printf(Green+"Cette clairvoyance aiguise votre esprit. Vous gagnez +%d Expdivination !\n"+Reset, gain)

	case "2":
		fmt.Println(Yellow + "Trelawney hoche gravement la tête : « Une massue... Des épreuves physiques s'annoncent. Protégez vos points de vie ! »" + Reset)
		soin := 15
		p.Health += soin
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
		gain := 8
		p.Expdivination += gain
		fmt.Printf(Green+"La prise de conscience vous galvanise (+%d PV, +%d Expdivination) !\n"+Reset, soin, gain)

	case "3":
		fmt.Println(Red + "Trelawney défaille presque en portant ses mains à sa gorge : « Le... Le RAVIER ! Mon Dieu, le chien ! La mort ! Oh, c'est terrible ! »" + Reset)
		fmt.Println("La tension dramatique vous stresse un peu. Vous perdez 5 PV par pure panique.")
		p.Health -= 5
		if p.Health < 1 {
			p.Health = 1
		}
		gain := 5
		p.Expdivination += gain
		fmt.Printf(Green+"Malgré la peur, vous apprenez de cette leçon (+%d Expdivination).\n"+Reset, gain)

	case "4":
		fmt.Println(Red + "Trelawney souffle : « Une tasse fendue... Des larmes et des projets brisés. Mais restez stoïque, l'avenir peut changer. »" + Reset)
		gainGallions := rand.Intn(5) + 2
		p.Gallions += gainGallions
		gainExpdivination := 8
		p.Expdivination += gainExpdivination
		fmt.Printf(Yellow+"En rangeant votre tasse, vous trouvez par hasard %d Gallions oubliés sur la table (+%d Expdivination).\n"+Reset, gainGallions, gainExpdivination)

	default:
		fmt.Println(Red + "Trelawney soupire : « Vous n'avez pas l'Œil de l'interne, mon pauvre ami... Tout cela n'est que bouillie de feuilles. »" + Reset)
		gain := 3
		p.Expdivination += gain
		fmt.Printf(Green+"Même dans l'erreur, vous apprenez (+%d Expdivination).\n"+Reset, gain)
	}

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour terminer la lecture de tasse]"+Reset)
}

// MiniJeuBouleDeCristal : Tenter de deviner ce qui se cache dans la brume de la sphère
func MiniJeuBouleDeCristal() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Red + "=== MINI-JEU : LA BOULE DE CRISTAL ===" + Reset)
	fmt.Println("Vous fixez intensément la grande sphère de verre transparente posée sur votre table.")
	fmt.Println("Des volutes de fumée argentée tourbillonnent à l'intérieur...")
	fmt.Println("Concentrez-vous. Que voulez-vous essayer de deviner en premier ?")
	time.Sleep(1 * time.Second)

	fmt.Println("\n 1 - Le résultat de votre prochain examen à Poudlard")
	fmt.Println(" 2 - La météo de demain dans le parc")
	fmt.Println(" 3 - Un secret caché dans le château")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre concentration (1-3) : "))

	fmt.Println(Cyan + "\nLa brume s'éclaircit un instant dans le cristal..." + Reset)
	time.Sleep(1 * time.Second)

	succes := rand.Intn(2) == 0

	if choix == "1" {
		if succes {
			fmt.Println(Green + "Flash ! Le cristal vous montre une copie marquée d'un 'O' éclatant (Optimal) !" + Reset)
			gain := 15
			p.Expdivination += gain
			fmt.Printf(Green+"Votre confiance en vous monte en flèche (+%d Expdivination).\n"+Reset, gain)
		} else {
			fmt.Println(Red + "Flash ! Le cristal montre une copie maculée d'encre rouge... Aïe." + Reset)
			gain := 5
			p.Expdivination += gain
			fmt.Printf(Green+"Vous apprenez de vos erreurs virtuelles (+%d Expdivination).\n"+Reset, gain)
		}
	} else if choix == "2" {
		fmt.Println(Yellow + "Le cristal se voile d'un gris menaçant : de grosses giboulées et du vent glacial s'abattront sur le Quidditch." + Reset)
		soin := 10
		p.Health += soin
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
		gain := 10
		p.Expdivination += gain
		fmt.Printf(Cyan+"Savoir anticiper le mauvais temps vous met à l'abri (+%d PV de repos mental, +%d Expdivination).\n"+Reset, soin, gain)
	} else if choix == "3" {
		if succes {
			fmt.Println(Green + "Vision fulgurante ! Le cristal dessine brièvement les contours d'une porte dérobée derrière une tapisserie au 7ème étage..." + Reset)
			gain := 20
			p.Expdivination += gain
			fmt.Printf(Green+"Découverte majeure (+%d Expdivination) !\n"+Reset, gain)
		} else {
			fmt.Println(Red + "Le cristal reste désespérément opaque. Rien d'autre que le reflet de votre propre nez inquiet." + Reset)
			gain := 5
			p.Expdivination += gain
			fmt.Printf(Green+"Tentative audacieuse (+%d Expdivination).\n"+Reset, gain)
		}
	} else {
		fmt.Println(Red + "Vous fixez le vide. Vos yeux pleurent à cause de la chaleur de la pièce." + Reset)
		gain := 2
		p.Expdivination += gain
		fmt.Printf(Green+"Effort noté (+%d Expdivination).\n"+Reset, gain)
	}

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour vous frotter les yeux]"+Reset)
}

// MiniJeuAstrologie : Calcul des influences célestes et des planètes
func MiniJeuAstrologie() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Red + "=== MINI-JEU : CALCULS ASTROLOGIQUES ===" + Reset)
	fmt.Println("Le Professeur Trelawney vous demande de tracer votre carte du ciel sur un parchemin")
	fmt.Println("en alignant l'angle d'incidence de Mars avec la constellation du Centaure.")
	time.Sleep(1 * time.Second)

	fmt.Println("\nQuelle position choisissez-vous pour votre planétarium miniature ?")
	fmt.Println(" 1 - Aligner Mars sur l'ascendant du Bélier (Voie martienne)")
	fmt.Println(" 2 - Placer Vénus en conjonction avec la Lune (Voie douce)")
	fmt.Println(" 3 - Isoler Neptune dans les brumes de l'océan astral (Voie mystique)")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix d'alignement (1-3) : "))

	fmt.Println(Cyan + "\nVous ajustez les petits anneaux de laiton doré de votre instrument..." + Reset)
	time.Sleep(1 * time.Second)

	switch choix {
	case "1":
		fmt.Println(Green + "Étincelle ! Les planètes s'imbriquent dans un tintement mélodieux de clochettes !" + Reset)
		gain := 12
		p.Expdivination += gain
		fmt.Printf(Green+"Votre force martienne est stimulée (+%d Expdivination).\n"+Reset, gain)
	case "2":
		fmt.Println(Cyan + "Une douce lueur opaline émane de votre feuille de calculs. Votre esprit trouve une paix absolue." + Reset)
		p.Health = p.MaxHealth
		gain := 10
		p.Expdivination += gain
		fmt.Printf(Green+"Vos points de vie sont entièrement restaurés grâce à l'harmonie céleste (+%d Expdivination) !\n"+Reset, gain)
	case "3":
		fmt.Println(Red + "Le planétarium se bloque net en émettant un petit nuage de fumée violette." + Reset)
		fmt.Println("Trelawney s'exclame : « Fascinant... Un mystère enveloppé dans une énigme ! »")
		gain := 15
		p.Expdivination += gain
		fmt.Printf(Green+"Mystère expdivinationloré avec succès (+%d Expdivination) !\n"+Reset, gain)
	default:
		fmt.Println(Red + "Vos calculs sont erronés. Les planètes semblent se moquer de vous depuis le plafond." + Reset)
		gain := 4
		p.Expdivination += gain
		fmt.Printf(Green+"Tentative enregistrée (+%d Expdivination).\n"+Reset, gain)
	}

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour ranger vos instruments]"+Reset)
}