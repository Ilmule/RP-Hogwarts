package src

import (
	"fmt"
	"time"
)

// AnimerPoudlardExpress affiche une animation de train en mouvement dans le terminal
func AnimerPoudlardExpress() {
	// Frames pour animer la fumée au-dessus de la cheminée
	fumees := []string{
		"     (  )   (   )  ",
		"    (   )  (  )    ",
		"   (  )   (   )    ",
	}

	frames := 25 // Nombre de pas pour l'animation

	for i := 0; i < frames; i++ {
		ClearTerminal()

		// Décalage progressif vers la droite
		espaces := ""
		for s := 0; s < (i % 15); s++ {
			espaces += " "
		}

		fumeesFrame := fumees[i%len(fumees)]

		fmt.Println("" + Red + "=== EN ROUTE POUR POUDLARD ===" + Reset + "")

		// Affichage de la fumée et du train
		fmt.Println(Cyan + espaces + fumeesFrame + Reset)
		fmt.Println(Red + espaces + "  ___   " + Yellow + "_______ " + Red + "______" + Reset)
		fmt.Println(Red + espaces + " |  |  " + Red + "|  P.E. || |__| |" + Reset)
		fmt.Println(Red + espaces + " |  |  " + Red + "| 9 3/4 || |  | |" + Reset)
		fmt.Println(DarkGray + espaces + "=|==|==|=======||=|==|=|=" + Reset)
		fmt.Println(Yellow + espaces + " (O) (O)       (O) (O)" + Reset)
		fmt.Println(Green + "______________________________________________________" + Reset)

		if i < frames-1 {
			fmt.Println("\n" + Yellow + "Tchou-tchou ! Le train traverse la campagne écossaise..." + Reset)
		} else {
			fmt.Println("\n" + Green + "Le Poudlard Express entre en gare de Pré-au-Lard !" + Reset)
		}

		time.Sleep(180 * time.Millisecond)
	}

	reader := Lecteur
	fmt.Println()
	fmt.Println(LightGray + "Appuyez sur Entree pour descendre du Poudlard Express" + Reset)
	fmt.Println()
	reader.ReadString('\n')
	ClearTerminal()
	Preaulard()
}