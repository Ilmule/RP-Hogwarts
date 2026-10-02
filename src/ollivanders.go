package src

import (
	"fmt"
	"strconv"
	"strings"
)

func LancerMagasinOllivander() {
	p := JoueurActuel
	reader := Lecteur
	// Une baguette coûte 7 Gallions
	prixG, prixM, prixN := 7, 0, 0 

	fmt.Println("\n" + Yellow + "=== OLLIVANDERS - FABRICANT DE BAGUETTES ===" + Reset)
	fmt.Printf("Bourse : %d G | %d M | %d N\n", p.Gallions, p.Mornilles, p.Noises)

	// Menu d'accueil
	fmt.Println("\nQue souhaitez-vous faire ?")
	fmt.Println("1 - Fabriquer une baguette (7 Gallions)")
	fmt.Println("0 - Quitter le magasin")

	choixMenu := LireSaisie(reader, "Entrez votre choix : ")
	switch strings.TrimSpace(choixMenu) {
	case "1":
		// On continue la création de la baguette
	case "0":
		fmt.Println("Vous quittez le magasin Ollivanders.")
		ClearTerminal()
		Chemindetraverse()
	case "leave" :
		Leave()
	case "inv" :
		AfficherInventaire(LancerMagasinOllivander)
	case "fournitures":
		Printlistefourniture()
	default:
		fmt.Println(Red + "Choix invalide." + Reset)
		ClearTerminal()
		LancerMagasinOllivander()
	}

	totalJoueur := (p.Gallions * 493) + (p.Mornilles * 29) + p.Noises
	if totalJoueur < (prixG * 493) {
		fmt.Println(Red + "Une baguette coûte 7 Gallions. Vous n'avez pas assez d'argent !" + Reset)
		return
	}

	bois := []string{"Houx", "Saule", "Sureau", "Chêne", "If"}
	fmt.Println("\n--- 1. Choisissez le bois ---")
	for i, b := range bois {
		fmt.Printf("%d - %s\n", i+1, b)
	}
	choixBois := LireSaisie(reader, "Entrez le numéro du bois : ")
	idxBois, err1 := strconv.Atoi(strings.TrimSpace(choixBois))
	if err1 != nil || idxBois < 1 || idxBois > len(bois) {
		fmt.Println(Red + "Choix invalide." + Reset)
		return
	}

	var taille float64
	for {
		saisie := LireSaisie(reader, "Choix de la taille (entre 9.0 et 14.5 pouces) : ")
		saisie = strings.ReplaceAll(saisie, ",", ".")
		val, err := strconv.ParseFloat(strings.TrimSpace(saisie), 64)
		if err == nil && val >= 9.0 && val <= 14.5 {
			taille = val
			break
		}
		fmt.Println(Red + "Taille invalide. Réessayez." + Reset)
	}

	coeurs := []string{"Plume de Phénix", "Crin de Licorne", "Ventricule de Dragon"}
	fmt.Println("\n--- 3. Choisissez le cœur magique ---")
	for i, c := range coeurs {
		fmt.Printf("%d - %s\n", i+1, c)
	}
	choixCoeur := LireSaisie(reader, "Entrez le numéro du cœur : ")
	idxCoeur, err2 := strconv.Atoi(strings.TrimSpace(choixCoeur))
	if err2 != nil || idxCoeur < 1 || idxCoeur > len(coeurs) {
		fmt.Println(Red + "Choix invalide, recommencez la création de votre baguette." + Reset)
		return
	}

	if p.Payer(prixG, prixM, prixN) {
		nomBaguette := fmt.Sprintf("Baguette en %s, %.1f pouces, cœur en %s", bois[idxBois-1], taille, coeurs[idxCoeur-1])
		p.AjouterItem(nomBaguette, TypeEquipement, prixG, 1)
		fmt.Println("\n" + Green + "Félicitations ! Vous avez fait fabriquer votre baguette :" + Reset)
		fmt.Println(Yellow + "-> " + nomBaguette + Reset)
		LancerMagasinOllivander()
	}
}