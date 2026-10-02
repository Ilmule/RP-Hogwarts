package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ConfigSalleDemande définit la structure d'une variante de la Salle sur Demande
type ConfigSalleDemande struct {
	Nom          string
	Description  string
	LootPossible bool
	LootNom      string
	LootType     TypeItem // Utilise les types définis dans types.go[cite: 4]
	LootPrix     int      // Valeur de revente ou prix estimé
}

// SalleSurDemande gère la génération aléatoire et l'exploration de la salle
func Secretssallesurdemande() {
	p := JoueurActuel
	reader := Lecteur

	rand.Seed(time.Now().UnixNano())

	// Liste des 16 configurations possibles de la Salle sur Demande
	sallesPossibles := []ConfigSalleDemande{
		{
			Nom:          "La Salle des Objets Cachés",
			Description:  "Des montagnes de meubles cassés, de vieux livres et d'objets perdus s'élèvent jusqu'au plafond comme une ville en ruine. L'endroit sent la poussière et les secrets oubliés.",
			LootPossible: true,
			LootNom:      "Vieux diadème en argent terni (Artefact ancien)",
			LootType:     TypeDivers,
			LootPrix:     50,
		},
		{
			Nom:          "La Salle d'Entraînement de l'A.D.",
			Description:  "Une vaste pièce éclairée par des torches, tapissée de coussins pour amortir les chutes. Des mannequins d'entraînement en bois se tiennent silencieusement dans les coins.",
			LootPossible: true,
			LootNom:      "Cape d'entraînement résistante aux sortilèges",
			LootType:     TypeEquipement,
			LootPrix:     15,
		},
		{
			Nom:          "Le Labyrinthe des Livres",
			Description:  "Des étagères titanesques remplies de grimoires s'étendent à perte de vue. Des parchemins volent doucement d'un rayonnage à l'autre comme des oiseaux de papier.",
			LootPossible: true,
			LootNom:      "Livre : Les Sortilèges Oubliés du Moyen-Âge",
			LootType:     TypeLivre,
			LootPrix:     25,
		},
		{
			Nom:          "La Serre Botanique Oubliée",
			Description:  "Une chaleur humide vous enveloppe. La salle ressemble à une jungle intérieure, envahie de lianes violettes et de fleurs chantantes.",
			LootPossible: true,
			LootNom:      "Racine de Mandragore séchée très rare",
			LootType:     TypeConsommable,
			LootPrix:     20,
		},
		{
			Nom:          "La Salle aux Pots de Chambre",
			Description:  "Une pièce absurde remplie d'une magnifique collection de pots de chambre de toutes les époques. Dumbledore en serait ravi.",
			LootPossible: false,
		},
		{
			Nom:          "La Cachette du Potionniste",
			Description:  "Des dizaines de chaudrons bouillonnent doucement, dégageant des fumées multicolores. Le mur est couvert d'étagères de fioles étiquetées.",
			LootPossible: true,
			LootNom:      "Fiole de Felix Felicis (imparfait)",
			LootType:     TypeConsommable,
			LootPrix:     40,
		},
		{
			Nom:          "L'Arsenal Gobelin",
			Description:  "Une armurerie médiévale étincelante. Des râteliers d'armes, des boucliers et des armures en fer gobelin recouvrent les murs de pierre.",
			LootPossible: true,
			LootNom:      "Gants de duel en mailles gobelines",
			LootType:     TypeEquipement,
			LootPrix:     35,
		},
		{
			Nom:          "La Salle de Repos Chaleureuse",
			Description:  "Un feu de cheminée crépite doucement, éclairant de profonds fauteuils en cuir et un gros tapis moelleux. C'est l'endroit parfait pour se détendre.",
			LootPossible: false,
		},
		{
			Nom:          "La Remise à Balais Poussiéreuse",
			Description:  "La pièce sent la cire et le bois verni. Des dizaines de vieux balais d'apprentissage sont empilés dans des tonneaux.",
			LootPossible: true,
			LootNom:      "Lunettes de vol vintage en laiton",
			LootType:     TypeEquipement,
			LootPrix:     12,
		},
		{
			Nom:          "L'Échiquier Géant Abandonné",
			Description:  "Le sol est un immense damier noir et blanc. Des morceaux de pièces d'échecs géantes en marbre brisé jonchent le sol.",
			LootPossible: true,
			LootNom:      "Pion en marbre blanc intact (Antiquité)",
			LootType:     TypeDivers,
			LootPrix:     15,
		},
		{
			Nom:          "L'Horlogerie Magique",
			Description:  "Le son de milliers de tic-tac remplit l'air. Des horloges, des sabliers et d'étranges mécanismes tournoient dans tous les sens.",
			LootPossible: true,
			LootNom:      "Rouage en or massif (Objet de valeur)",
			LootType:     TypeDivers,
			LootPrix:     45,
		},
		{
			Nom:          "Les Bains Luxueux Isolés",
			Description:  "Une piscine de marbre blanc remplie d'eau chaude parfumée, entourée d'une centaine de robinets dorés crachant des bulles multicolores.",
			LootPossible: false,
		},
		{
			Nom:          "La Salle des Miroirs",
			Description:  "Une salle circulaire couverte de miroirs déformants, de miroirs à double sens et de glaces antiques. Votre reflet semble bouger de son propre chef.",
			LootPossible: false,
		},
		{
			Nom:          "L'Observatoire Secret",
			Description:  "Le plafond n'existe pas : il ouvre directement sur un ciel étoilé immaculé. Un énorme télescope en cuivre trône au centre.",
			LootPossible: true,
			LootNom:      "Lentille de télescope en cristal pur",
			LootType:     TypeDivers,
			LootPrix:     30,
		},
		{
			Nom:          "La Grotte de Glace",
			Description:  "Il fait un froid glacial. Les murs sont recouverts de stalactites bleutées et le sol est une patinoire parfaite.",
			LootPossible: true,
			LootNom:      "Cristal de givre éternel",
			LootType:     TypeDivers,
			LootPrix:     18,
		},
		{
			Nom:          "La Salle de Musique Ensorcelée",
			Description:  "Des dizaines d'instruments flottent dans les airs, jouant doucement une symphonie mélancolique tout seuls.",
			LootPossible: true,
			LootNom:      "Flûte enchantée en argent (Instrument rare)",
			LootType:     TypeDivers,
			LootPrix:     25,
		},
	}

	// Sélection aléatoire de la salle pour cette entrée
	salleActuelle := sallesPossibles[rand.Intn(len(sallesPossibles))]
	dejaFouille := false

	ClearTerminal()
	fmt.Println(Yellow + "Vous passez trois fois devant le mur vierge du 7ème étage en pensant très fort à ce dont vous avez besoin..." + Reset)
	time.Sleep(2 * time.Second)
	fmt.Println(Cyan + "Une porte en bois verni ornée d'une poignée en laiton apparaît par magie. Vous entrez." + Reset)
	time.Sleep(2 * time.Second)

	for {
		ClearTerminal()
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Println(Cyan + "               LA SALLE SUR DEMANDE                       " + Reset)
		fmt.Println(Cyan + "==========================================================" + Reset)
		fmt.Println(Yellow + "Lieu : " + salleActuelle.Nom + Reset)
		fmt.Println("\n" + salleActuelle.Description + "\n")
		fmt.Println("----------------------------------------------------------")
		fmt.Println(" 1 - Fouiller et explorer les recoins de la pièce")
		fmt.Println(" 2 - S'imprégner de la magie du lieu (Méditer)")
		fmt.Println(" 3 - Sortir de la Salle (La porte disparaîtra derrière vous)")
		fmt.Println(" INV - Ouvrir votre sac à dos (Inventaire)")

		choix := strings.TrimSpace(strings.ToUpper(LireSaisie(reader, "\nQue voulez-vous faire ? : ")))

		switch choix {
		case "1":
			ClearTerminal()
			if dejaFouille {
				fmt.Println(Red + "Vous avez déjà fouillé cette pièce de fond en comble. Il n'y a plus rien d'intéressant." + Reset)
			} else {
				dejaFouille = true
				fmt.Println(Yellow + "Vous fouillez attentivement les moindres recoins de la salle..." + Reset)
				time.Sleep(1 * time.Second)

				if salleActuelle.LootPossible {
					fmt.Println(Green + "Incroyable ! Vous avez déniché quelque chose d'intéressant :" + Reset)
					fmt.Println(Cyan + "-> " + salleActuelle.LootNom + Reset)
					
					// Ajout à l'inventaire structuré
					p.AjouterItem(salleActuelle.LootNom, salleActuelle.LootType, salleActuelle.LootPrix, 1)
					
					if salleActuelle.LootType == TypeEquipement {
						fmt.Println(DarkGray + "(Vous pouvez l'équiper depuis votre inventaire ou le revendre plus tard !)" + Reset)
					} else {
						fmt.Println(DarkGray + "(L'objet a été glissé dans votre sac. Idéal pour être revendu à un bon prix !)" + Reset)
					}
				} else {
					fmt.Println(DarkGray + "Vous ne trouvez rien de valeur, juste de la vieille poussière magique et des curiosités sans intérêt." + Reset)
				}
			}
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "2":
			ClearTerminal()
			fmt.Println(Cyan + "Vous fermez les yeux et respirez profondément l'air chargé de magie antique..." + Reset)
			time.Sleep(1 * time.Second)
			gainExp := rand.Intn(10) + 5
			p.Exp += gainExp
			fmt.Printf(Green+"Vous vous sentez revigoré. (+%d Exp)\n"+Reset, gainExp)
			LireSaisie(reader, DarkGray+"\n[Appuyez sur Entrée pour continuer]"+Reset)

		case "3":
			ClearTerminal()
			fmt.Println(Yellow + "Vous ressortez dans le couloir. Le mur redevient parfaitement lisse derrière vous." + Reset)
			time.Sleep(1 * time.Second)
			Halldepoudlard() // Remplace par la fonction du hub des étages si tu l'as créée[cite: 1]
			return

		case "INV":
			ClearTerminal()
			AfficherInventaire(Secretssallesurdemande) //[cite: 1]
			return

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}