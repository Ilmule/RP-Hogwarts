package src

import "fmt"

func Debutjeu() {
	reader := Lecteur
	fmt.Println("Vous allez commencer votre aventure au sein de Poudlard")
	fmt.Println()
	fmt.Println("Avant cela vous allez devoir faire vos emplettes dans le Chemin de traverse")
	fmt.Println()
	fmt.Println("Pour ce faire vous pourrez à tout moment consulter votre liste de fourniture scolaires en tapant")
	fmt.Println("'fournitures' quel que soit l'endroit ou vous vous trouvez au sein du Chemin de traverse")
	fmt.Println()
	fmt.Println(LightGray + "\nAppuyez sur Entrée pour vous rendre au chaudron baveur et commencer votre aventure" + Reset)
	fmt.Println()
	reader.ReadString('\n')
	ClearTerminal()
	Chaudronbaveur()
}

func Printlistefourniture() {
	fmt.Println(`
	Collège Poudlard - Ecole de sorcellerie


	Liste des vêtements dont les élèves devront obligatoirement être équipés

	1) 3 robes de travail (noires), modèle normal
	2) Un chapeau pointu (noir)
	3) Une paire de gants protecteurs (en cuir de dragon ou autre matière semblable)
	4) Une cape d'hiver (noire vec attaches d'argent)



	Livres et manuels

	Chaque élève devra se procurer un exemplaire des ouvrages suivants : 


	Le livre des sorts et enchantements, niveau 1
	par Miranda Fauconnette

	Histoire de la magie
	par Bathilda Tourdesac
	
	Magie théorique
	par Adalbert Lasornette

	Manuel de métamorphose à l'usage des débutants
	par Emeric G.Changé

	Mille herbes et champignons magiques
	par Phyllida Augirolle

	Potions magiques
	par Arsenius Beaulitron

	Vie et habitat des animaux fantastiques
	par Quentin Jentremble

	Forces obscures : comment s'en protéger
	par quentin Jentremble




	Fournitures

	- 1 baguette magique
	- 1 chaudron (modèle standard en étain, taille 2)
	- 1 boite de fioles en verre ou en crystal
	- 1 téléscope
	- 1 balance en cuivre




	Les élèves peuvent également emporter

	un hibou
	OU un chat
	OU un crapaud
	
	
	
	`)
}
