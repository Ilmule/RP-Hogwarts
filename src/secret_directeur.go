package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// BureauDumbledore gère la visite du bureau directorial de Poudlard
func Secretdirecteur() {
	p := JoueurActuel
	reader := Lecteur

	for {
		fmt.Println("\n" + Yellow + "==========================================================" + Reset)
		fmt.Println(Yellow + "                  LE BUREAU DU DIRECTEUR                  " + Reset)
		fmt.Println(Yellow + "==========================================================" + Reset)
		
		fmt.Printf("Santé : %d/%d PV | Exp : %d\n", p.Health, p.MaxHealth, p.Exp)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Vous vous trouvez dans une vaste pièce circulaire pleine de bruits étranges.")
		fmt.Println("Sur de petites tables à pieds grêles, des instruments en argent tourbillonnent.")
		fmt.Println("Le bureau est temporairement inoccupé par le directeur, mais Neville Londubat,")
		fmt.Println("qui semble chercher un livre précis, se tient près de l'immense bibliothèque.")

		fmt.Println(" 1 - Discuter avec Neville Londubat")
		fmt.Println(" 2 - Observer les tableaux des anciens directeurs")
		fmt.Println(" 3 - Plonger son visage dans la Pensine (Mini-événement)")
		fmt.Println(" 4 - Caresser Fumsec le phénix (Restaurer ses PV)")
		fmt.Println(" 5 - Prendre un bonbon au citron dans la coupelle")
		fmt.Println(" R - Redescendre l'escalier en colimaçon (Retour aux Tours)")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			DiscuterNeville()
		case "2":
			ClearTerminal()
			ObserverTableauxDirecteurs()
		case "3":
			ClearTerminal()
			UtiliserPensine()
		case "4":
			ClearTerminal()
			CaresserFumsec()
		case "5":
			ClearTerminal()
			MangerBonbonCitron()
		case "inv" :
			ClearTerminal()
			AfficherInventaire(Secretdirecteur)
		case "R":
			ClearTerminal()
			fmt.Println(Yellow + "Vous montez sur la gargouille de pierre qui tourne sur elle-même pour vous descendre..." + Reset)
			time.Sleep(1 * time.Second)
			// Remplacer par la fonction de retour vers ta zone de navigation (ex: ToursAstronomie() ou Halldepoudlard())
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// DiscuterNeville lance un dialogue aléatoire avec Neville Londubat
func DiscuterNeville() {
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	dialogues := []string{
		"« Oh, salut ! Je cherchais un vieil ouvrage sur les Branchiflores que le professeur Chourave m'a conseillé. Ne touche à rien, l'épée de Gryffondor est très coupante... »",
		"« Harry m'a raconté comment il a détruit le journal de Jedusor ici même. C'est fou quand on y pense, non ? »",
		"« Trevor s'est encore échappé. J'espérais qu'il ne se soit pas caché sous le bureau du directeur... »",
		"« Tu as vu le Choixpeau magique sur l'étagère ? Parfois, j'ai l'impression qu'il me fixe en se disant qu'il a eu raison de me mettre à Gryffondor. »",
	}

	choixDialogue := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Yellow + "=== DISCUSSION AVEC NEVILLE LONDUBAT ===" + Reset)
	fmt.Println("Neville sursaute légèrement en vous voyant approcher et lâche presque un parchemin.")
	fmt.Println()
	fmt.Println(Cyan + "Neville : " + choixDialogue + Reset)
	fmt.Println()
	
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour continuer]"+Reset)
}

// ObserverTableauxDirecteurs permet d'interagir avec les portraits endormis
func ObserverTableauxDirecteurs() {
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	fmt.Println(Yellow + "=== LES PORTRAITS DES ANCIENS DIRECTEURS ===" + Reset)
	fmt.Println("Les murs sont recouverts de portraits de sorciers et de sorcières qui dorment paisiblement dans leurs cadres.")
	time.Sleep(1 * time.Second)

	chance := rand.Intn(100)
	if chance < 30 {
		fmt.Println(Cyan + "\nPhineas Nigellus Black ouvre un œil paresseux et vous toise avec mépris :" + Reset)
		fmt.Println(Cyan + "« Encore un élève qui traîne là où il ne devrait pas... De mon temps, vous seriez déjà en retenue à récurer les cachots ! »" + Reset)
	} else if chance < 60 {
		fmt.Println(Cyan + "\nLe portrait d'Armando Dippet ronfle bruyamment, une petite bulle de salive éclatant à chaque respiration." + Reset)
	} else {
		fmt.Println(Cyan + "\nUne ancienne directrice à l'air sévère ajuste ses lunettes en demi-lune et vous fait un clin d'œil complice avant de faire semblant de dormir." + Reset)
	}

	fmt.Println()
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour vous éloigner des portraits]"+Reset)
}

// UtiliserPensine génère un souvenir aléatoire immersif
func UtiliserPensine() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	fmt.Println(Cyan + "=== LA PENSINE ===" + Reset)
	fmt.Println("Une bassine en pierre peu profonde est posée sur un piédestal. Elle est remplie d'une substance argentée,")
	fmt.Println("qui n'est ni un liquide ni un gaz. Vous penchez votre visage vers la surface...")
	time.Sleep(2 * time.Second)

	souvenirs := []string{
		"Vous voyez un jeune Rogue, l'air sombre, marcher dans la cour de Poudlard avec Lily Evans. Le vent souffle dans leurs robes.",
		"La Grande Salle apparaît, floue. Vous entendez Dumbledore crier 'HARRY POTTER !' pendant que la Coupe de Feu crépite en arrière-plan.",
		"Vous êtes dans une caverne sombre entourée d'eau noire. Un Dumbledore affaibli boit une potion émeraude dans une vasque en gémissant.",
		"Vous apercevez Tom Jedusor, jeune et charmeur, demandant des informations sur les Horcruxes au professeur Slughorn.",
	}

	souvenir := souvenirs[rand.Intn(len(souvenirs))]

	fmt.Println("\n" + Blue + "Le bureau disparaît... Vous êtes aspiré dans un tourbillon d'argent." + Reset)
	time.Sleep(1 * time.Second)
	fmt.Println(Blue + souvenir + Reset)
	time.Sleep(2 * time.Second)
	fmt.Println("\n" + Blue + "D'un coup sec, vous relevez la tête en haletant. Vous êtes de retour dans le bureau." + Reset)
	
	// Petite récompense pour avoir vu un souvenir caché
	fmt.Println(Green + "La sagesse de ce souvenir vous octroie +10 Exp !" + Reset)
	p.Exp += 10

	fmt.Println()
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour reprendre vos esprits]"+Reset)
}

// CaresserFumsec restaure les PV grâce aux larmes du phénix
func CaresserFumsec() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Red + "=== FUMSEC LE PHÉNIX ===" + Reset)
	fmt.Println("Sur un perchoir en or repose un oiseau magnifique de la taille d'un cygne, avec des plumes rouges et or.")
	fmt.Println("Il vous observe avec ses petits yeux noirs et brillants.")
	time.Sleep(1 * time.Second)

	if p.Health < p.MaxHealth {
		fmt.Println(Yellow + "\nFumsec remarque que vous êtes blessé. Il penche la tête et laisse tomber une larme perlée sur vous." + Reset)
		p.Health = p.MaxHealth
		fmt.Println(Green + "Les propriétés curatives incroyables des larmes de phénix vous soignent instantanément ! (PV restaurés au maximum)" + Reset)
	} else {
		fmt.Println(Yellow + "\nVous caressez doucement ses plumes chaudes. Fumsec pousse un petit trille musical qui vous emplit de courage et de paix." + Reset)
	}

	fmt.Println()
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour continuer]"+Reset)
}

// MangerBonbonCitron offre un petit bonus amusant
func MangerBonbonCitron() {
	p := JoueurActuel
	reader := Lecteur

	fmt.Println(Yellow + "=== LA COUPELLE DE BONBONS ===" + Reset)
	fmt.Println("Une coupelle en cristal est remplie de petites douceurs jaunes, les fameux 'Sherbet Lemons' chéris par Dumbledore.")
	fmt.Println("Vous en prenez un et le mettez dans votre bouche.")
	time.Sleep(1 * time.Second)

	fmt.Println(Green + "\nLe goût acidulé explose sur votre langue ! C'est délicieux et très revigorant." + Reset)
	
	// Petit soin bonus
	if p.Health < p.MaxHealth {
		soin := 5
		p.Health += soin
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
		fmt.Printf(Green + "Vous récupérez %d PV !\n" + Reset, soin)
	}

	fmt.Println()
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour continuer]"+Reset)
}