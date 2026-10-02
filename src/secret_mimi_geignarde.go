package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Secretsmimigeignarde gère l'exploration des toilettes des filles au 2e étage
func Secretsmimigeignarde() {
	reader := Lecteur

	for {
		fmt.Println("\n" + Blue + "=============== Toilettes du 2ème étage ===============" + Reset)
		fmt.Println("L'endroit est lugubre, les miroirs sont ébréchés et le sol est inondé.")
		fmt.Println("Des gémissements résonnent depuis l'une des cabines fermées...")
		fmt.Println(Blue + "-------------------------------------------------------" + Reset)
		fmt.Println(" 1 - Parler à Mimi Geignarde")
		fmt.Println(" 2 - Inspecter les lavabos en forme de serpent")
		fmt.Println(" 3 - Retourner au Hall")
		fmt.Println(" INV - Ouvrir votre inventaire")
		fmt.Println(Blue + "-------------------------------------------------------" + Reset)

		Option := strings.TrimSpace(strings.ToLower(LireSaisie(reader, DarkGray+"Que voulez-vous faire ? : "+Reset)))

		switch Option {
		case "1":
			ClearTerminal()
			DialoguesMimiGeignarde()
		case "2":
			ClearTerminal()
			fmt.Println(Yellow + "Vous inspectez les robinets rouillés. Rien ne se passe... pour le moment." + Reset)
			time.Sleep(2 * time.Second)
			ClearTerminal()
		case "3":
			ClearTerminal()
			Halldepoudlard()
			return // Permet de ne pas empiler les fonctions
		case "inv":
			ClearTerminal()
			AfficherInventaire(Secretsmimigeignarde)
			return
		default:
			ClearTerminal()
			fmt.Println(Red + "Choix invalide." + Reset)
		}
	}
}

// DialoguesMimiGeignarde génère une conversation aléatoire avec le fantôme
func DialoguesMimiGeignarde() {
	reader := Lecteur
	rand.Seed(time.Now().UnixNano())

	dialogues := []string{
		"« Oh ! Tu viens te moquer de moi ? TOUT LE MONDE SE MOQUE DE MIMI ! Ouuuh... *snif* »",
		"« Si tu meurs un jour, tu pourras partager mes toilettes ! Ce serait chouette... »",
		"« J'étais juste assise là, à pleurer, et puis j'ai vu de grands yeux jaunes... et c'était fini. »",
		"« Regarde-moi ce carnage ! Quelqu'un a encore jeté un livre dans mes toilettes ! C'est intolérable ! »",
		"« Tu sais, Harry vient parfois me voir... Lui au moins, il est poli avec moi ! »",
		"« Ouuuuuh ! *Mimi plonge la tête la première dans la cuvette en vous arrosant d'eau sale* »",
	}

	choixDialogue := dialogues[rand.Intn(len(dialogues))]

	fmt.Println(Cyan + "=== RENCONTRE AVEC MIMI GEIGNARDE ===" + Reset)
	fmt.Println("Le fantôme translucide d'une jeune fille à lunettes surgit d'une des cabines.")
	fmt.Println()
	fmt.Println(Blue + "Mimi Geignarde : " + choixDialogue + Reset)
	fmt.Println()
	
	LireSaisie(reader, DarkGray+"[Appuyez sur Entrée pour continuer]"+Reset)
	ClearTerminal()
}