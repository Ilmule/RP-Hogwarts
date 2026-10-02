package src

type TypeItem string

const (
	TypeEquipement  TypeItem = "Équipement"
	TypeConsommable TypeItem = "Consommable"
	TypeLivre       TypeItem = "Livre"
	TypeAnimal      TypeItem = "Animal"
	TypeDivers      TypeItem = "Divers"
)

type Item struct {
	Nom         string
	Description string
	Type        TypeItem
	PrixValeur  int
	Quantite    int
}

// ItemBoutique représente un article vendu dans un magasin
type ItemBoutique struct {
	Nom           string
	PrixGallions  int
	PrixMornilles int
	PrixNoises    int
}