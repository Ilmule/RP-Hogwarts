package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// CoursSortileges gère l'accès au cours de sortilèges du Professeur Flitwick
func Courssortilège() {
	p := JoueurActuel
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	for {
		ClearTerminal()
		// Calcul de l'année en fonction du niveau (1-10 = 1ère année, 11-20 = 2ème année, etc.)
		annee := (p.Levelsortilege-1)/10 + 1
		if annee > 7 {
			annee = 7
		}

		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Println(Cyan + "               SALLE DE COURS DE SORTILÈGES               " + Reset)
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Printf("Santé : %d/%d PV | Exp sortilège : %d | Niveau : %d (Année %d)\n", p.Health, p.MaxHealth, p.Expsortilege, p.Levelsortilege, annee)
		fmt.Println("----------------------------------------------------------")
		fmt.Println("La salle est baignée de lumière. Des piles de grimoires flottent")
		fmt.Println("légèrement au-dessus des pupitres. Sur une haute pile de gros livres")
		fmt.Println("au fond, le minuscule Professeur Flitwick agite sa baguette avec entrain.")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Écouter les conseils du Professeur Flitwick (Dialogue)")
		fmt.Println(" 2 - Pratiquer les sorts de 1ère année (Niveaux 1 à 10)")
		fmt.Println(" 3 - Pratiquer les sorts de 2ème année (Niveaux 11 à 20)")
		fmt.Println(" 4 - Pratiquer les sorts avancés (Niveaux 21 et plus)")
		fmt.Println(" 5 - Retourner au Hall de Poudlard")
		fmt.Println()
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")
		fmt.Println(Cyan + "==========================================================" + Reset)

		Option := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch Option {
		case "1":
			ClearTerminal()
			DialoguesFlitwick()
		case "2":
			ClearTerminal()
			MenuPremiereAnnee()
		case "3":
			ClearTerminal()
			MenuDeuxiemeAnnee()
		case "4":
			ClearTerminal()
			MenuAnneeSuperieure()
		case "5":
			ClearTerminal()
			fmt.Println(Yellow + "Vous saluez Flitwick d'une inclinaison de tête et quittez la salle..." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard()
			return
		case "INV":
			ClearTerminal()
			AfficherInventaire(Courssortilège)
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// DialoguesFlitwick gère les encouragements du professeur
func DialoguesFlitwick() {
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	dialogues := []string{
		"« N'oubliez pas le mouvement du poignet ! C'est la rotation qui fait tout le charme du sortilège ! »",
		"« Remarquez bien la prononciation : ce n'est pas *Leviosá*, mais bien *Levio-sà* ! »",
		"« La magie réside autant dans l'intention que dans la précision du geste, mes chers élèves ! »",
		"« Cinq points pour votre maison si vous parvenez à faire léviter cette plume du premier coup ! »",
	}

	phrase := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Cyan + "=== PROFESSEUR FILIUS FLITWICK ===" + Reset)
	fmt.Println("Flitwick sautille sur sa pile de livres en vous apercevant.")
	fmt.Println()
	fmt.Printf(Cyan+"Flitwick : %s\n"+Reset, phrase)
	fmt.Println()

	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour reprendre vos exercices]"+Reset)
}

// MenuPremiereAnnee regroupe les sorts de niveau 1 à 10
func MenuPremiereAnnee() {
	p := JoueurActuel
	reader := Lecteur

	for {
		ClearTerminal()
		fmt.Println(Yellow + "=== COURS DE 1ÈRE ANNÉE (Niveaux 1 à 10) ===" + Reset)
		fmt.Printf("Votre niveau actuel : %d\n\n", p.Levelsortilege)
		fmt.Println(" 1 - Lumos (Niveau 1 requis) - Illumine les endroits sombres")
		fmt.Println(" 2 - Wingardium Leviosa (Niveau 3 requis) - Fait léviter les objets")
		fmt.Println(" 3 - Alohomora (Niveau 6 requis) - Déverrouille les serrures simples")
		fmt.Println(" R - Retour au menu principal des sortilèges")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "Votre choix : ")))

		switch choix {
		case "1":
			ClearTerminal()
			LancerSortLumos(p)
		case "2":
			ClearTerminal()
			LancerSortWingardium(p)
		case "3":
			ClearTerminal()
			LancerSortAlohomora(p)
		case "R":
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// MenuDeuxiemeAnnee regroupe les sorts de niveau 11 à 20
func MenuDeuxiemeAnnee() {
	p := JoueurActuel
	reader := Lecteur

	for {
		ClearTerminal()
		fmt.Println(Yellow + "=== COURS DE 2ÈME ANNÉE (Niveaux 11 à 20) ===" + Reset)
		fmt.Printf("Votre niveau actuel : %d\n\n", p.Levelsortilege)
		fmt.Println(" 1 - Accio (Niveau 11 requis) - Attire un objet à distance")
		fmt.Println(" 2 - Riddikulus (Niveau 14 requis) - Repousse les Épouvantards")
		fmt.Println(" R - Retour au menu principal des sortilèges")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "Votre choix : ")))

		switch choix {
		case "1":
			ClearTerminal()
			LancerSortAccio(p)
		case "2":
			ClearTerminal()
			LancerSortRiddikulus(p)
		case "R":
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// MenuAnneeSuperieure regroupe les sorts de niveau 21 et plus
func MenuAnneeSuperieure() {
	p := JoueurActuel
	reader := Lecteur

	for {
		ClearTerminal()
		fmt.Println(Yellow + "=== COURS AVANCÉ (Niveaux 21+) ===" + Reset)
		fmt.Printf("Votre niveau actuel : %d\n\n", p.Levelsortilege)
		fmt.Println(" 1 - Expsortilegeecto Patronum (Niveau 21 requis) - Invoque un Patronus protecteur")
		fmt.Println(" R - Retour au menu principal des sortilèges")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "Votre choix : ")))

		switch choix {
		case "1":
			ClearTerminal()
			LancerSortPatronum(p)
		case "R":
			return
		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			time.Sleep(1 * time.Second)
		}
	}
}

// --- MINI-JEUX DE SORTILÈGES ---

func LancerSortLumos(p *Player) {
	reader := Lecteur
	const niveauRequis = 1

	if p.Levelsortilege < niveauRequis {
		fmt.Println(Red + "Votre niveau est insuffisant ! Vous devez être au moins niveau " + fmt.Sprint(niveauRequis) + "." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== APPRENTISSAGE : LUMOS ===" + Reset)
	fmt.Println("Pointez votre baguette et prononcez l'incantation avec la bonne intensité.")
	fmt.Println("1 - Chuchoter « Lumos » doucement")
	fmt.Println("2 - Prononcer clairement « Lumos ! »")
	fmt.Println("3 - Crier de toutes vos forces « LUMOS !! »")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre méthode : "))
	time.Sleep(1 * time.Second)

	if choix == "2" {
		fmt.Println(Green + "Parfait ! Le bout de votre baguette s'illumine d'un éclat brillant et stable !" + Reset)
		fmt.Println("Flitwick applaudit : « Excellent dosage ! » (+10 Expsortilege)")
		p.Expsortilege += 10
	} else if choix == "3" {
		fmt.Println(Yellow + "Trop intense ! Le bout de votre baguette brille tellement fort qu'il aveugle vos voisins !" + Reset)
		fmt.Println("Flitwick sourit : « Un peu trop d'enthousiasm, mais l'effet est là. » (+5 Expsortilege)")
		p.Expsortilege += 5
	} else {
		fmt.Println(Red + "Trop faible... Rien ne se passe au bout de votre baguette." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func LancerSortWingardium(p *Player) {
	reader := Lecteur
	const niveauRequis = 3

	if p.Levelsortilege < niveauRequis {
		fmt.Println(Red + "Votre niveau est insuffisant ! Vous devez être au moins niveau " + fmt.Sprint(niveauRequis) + " pour ce sort." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== APPRENTISSAGE : WINGARDIUM LEVIOSA ===" + Reset)
	fmt.Println("Une plume repose sur votre pupitre. Effectuez le mouvement : pique et secoue !")
	fmt.Println("Quelle prononciation choisissez-vous ?")
	fmt.Println("1 - Wing-gar-dium Levi-o-sa (Accentuation correcte)")
	fmt.Println("2 - Wing-gar-di-um Levios-á (Accentuation erronée)")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "La plume frémit, s'élève gracieusement dans les airs et flotte à hauteur de visage !" + Reset)
		fmt.Println("Flitwick s'écrie : « Regardez tous ! C'est une réussite parfaite ! » (+15 Expsortilege)")
		p.Expsortilege += 15
	} else {
		fmt.Println(Red + "La plume se retourne brusquement et retombe lourdement. Ron Weasley n'est pas loin pour se moquer." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func LancerSortAlohomora(p *Player) {
	reader := Lecteur
	const niveauRequis = 6

	if p.Levelsortilege < niveauRequis {
		fmt.Println(Red + "Votre niveau est insuffisant ! Vous devez être au moins niveau " + fmt.Sprint(niveauRequis) + " pour ce sort." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== APPRENTISSAGE : ALOHOMORA ===" + Reset)
	fmt.Println("Flitwick place un petit coffret en bois fermé par un verrou magique devant vous.")
	fmt.Println("Quel mouvement de baguette effectuez-vous ?")
	fmt.Println("1 - Une boucle en S rapide")
	fmt.Println("2 - Une spirale vers la droite suivie d'un point")
	fmt.Println("3 - Un tracé en forme de flèche vers le haut")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre mouvement : "))
	time.Sleep(1 * time.Second)

	if choix == "2" {
		fmt.Println(Green + "CLIC ! Le verrou du coffret saute net et le couvercle s'ouvre tout seul !" + Reset)
		fmt.Println("Vous gagnez +20 Expsortilege et empochez 3 Gallions trouvés à l'intérieur !")
		p.Expsortilege += 20
		p.Gallions += 3
	} else {
		fmt.Println(Red + "Le coffret vibre, émet un bourdonnement agacé mais reste hermétiquement fermé." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func LancerSortAccio(p *Player) {
	reader := Lecteur
	const niveauRequis = 11

	if p.Levelsortilege < niveauRequis {
		fmt.Println(Red + "Votre niveau est insuffisant ! Vous devez être au moins niveau " + fmt.Sprint(niveauRequis) + " (2ème année)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== APPRENTISSAGE : ACCIO ===" + Reset)
	fmt.Println("Un manuel de sortilèges est posé à l'autre bout de la pièce.")
	fmt.Println("Visualisez l'objet et criez le sort :")
	fmt.Println("1 - « ACCIO LIVRE ! » avec assurance")
	fmt.Println("2 - « Accio... » timidement en hésitant")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre incantation : "))
	time.Sleep(1 * time.Second)

	if choix == "1" {
		fmt.Println(Green + "Le livre fend les airs en trombe et vient atterrir directement dans votre paume tendue !" + Reset)
		fmt.Println("Flitwick applaudit : « Impressionnant ! Quel réflexe ! » (+25 Expsortilege)")
		p.Expsortilege += 25
	} else {
		fmt.Println(Red + "Le livre tremble à l'autre bout de la salle, glisse de deux centimètres... et s'arrête." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func LancerSortRiddikulus(p *Player) {
	reader := Lecteur
	const niveauRequis = 14

	if p.Levelsortilege < niveauRequis {
		fmt.Println(Red + "Votre niveau est insuffisant ! Vous devez être au moins niveau " + fmt.Sprint(niveauRequis) + "." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== APPRENTISSAGE : RIDDIKULUS ===" + Reset)
	fmt.Println("Un faux Épouvantard surgit d'une armoire magique et prend la forme de votre pire peur !")
	fmt.Println("En riant un bon coup, en quoi le transformez-vous ?")
	fmt.Println("1 - En un énorme ballon gonflable qui éclate en confettis")
	fmt.Println("2 - En un canard en plastique géant chaussé de patins à roulettes")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre choix (1 ou 2) : "))
	time.Sleep(1 * time.Second)

	if choix == "1" || choix == "2" {
		fmt.Println(Green + "Le monstre se tort, prend l'apparence comique que vous avez imaginée et se désintègre dans un éclat de rire !" + Reset)
		fmt.Println("Votre maîtrise du courage face à la peur vous rapporte +30 Expsortilege !")
		p.Expsortilege += 30
	} else {
		fmt.Println(Red + "Vous hésitez... L'Épouvantard en profite pour redoubler d'effrayantes grimaces." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
}

func LancerSortPatronum(p *Player) {
	reader := Lecteur
	const niveauRequis = 21

	if p.Levelsortilege < niveauRequis {
		fmt.Println(Red + "Votre niveau est insuffisant ! Vous devez être au moins niveau " + fmt.Sprint(niveauRequis) + " (Niveau avancé)." + Reset)
		LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée]"+Reset)
		return
	}

	fmt.Println(Cyan + "=== APPRENTISSAGE : EXPsortilegeECTO PATRONUM ===" + Reset)
	fmt.Println("La pièce s'assombrit légèrement. Concentrez-vous sur votre souvenir le plus pur et le plus heureux.")
	fmt.Println("Quel souvenir invoquez-vous ?")
	fmt.Println("1 - Votre premier vol réussi sur un balai magique")
	fmt.Println("2 - Un moment de partage inoubliable avec vos meilleurs amis au festin")
	fmt.Println("3 - Le jour de l'obtention de votre lettre pour Poudlard")

	choix := strings.TrimSpace(LireSaisie(reader, "Votre souvenir (1-3) : "))
	time.Sleep(2 * time.Second)

	if choix == "1" || choix == "2" || choix == "3" {
		fmt.Println(Green + "Une lumière argentée, aveuglante et pure jaillit du bout de votre baguette !" + Reset)
		fmt.Println("Un magnifique animal de fumée argentée prend forme et gambade joyeusement autour de la salle.")
		fmt.Println("Flitwick retient ses larmes : « C'est... c'est un Patronus corporel d'une puissance rare ! (+50 Expsortilege) »")
		p.Expsortilege += 50
	} else {
		fmt.Println(Red + "Votre esprit est embrumé par le doute. Seuls de faibles volutes de fumée grise s'échappent de votre baguette." + Reset)
	}
	LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour reprendre souffle]"+Reset)
}