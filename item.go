package red

const (
	chapeauAventurier     = "Chapeau de l'aventurier"
	tuniqueAventurier     = "Tunique de l'aventurier"
	bottesAventurier      = "Bottes de l'aventurier"
	potionDeVie           = "Potion de vie"
	potionDePoison        = "Potion de poison"
	bouleDeFeu            = "Boule de feu"
	livreDeSortBouleDeFeu = "Livre de sort : Boule de feu"
)

func (char1 *Character) AddItem(item string) {
	// Vérifiez si l'élément est déjà présent dans l'inventaire
	if _, exists := char1.inventory[item]; exists {
		// L'élément existe déjà, vous pouvez ajouter votre logique ici si nécessaire.
	} else {
		// L'élément n'existe pas, vous pouvez l'ajouter
		char1.inventory[item] = 1 // Vous pouvez ajuster la quantité si nécessaire
	}

	// Gérez l'effet de l'ajout de l'item sur les caractéristiques du personnage ici
}

func (char1 *Character) RemoveItem(item string) {
	// Vérifiez si l'élément est présent dans l'inventaire
	if _, exists := char1.inventory[item]; exists {
		// L'élément existe, vous pouvez retirer votre logique ici si nécessaire.
		// Gérez l'effet du retrait de l'item sur les caractéristiques du personnage ici

		// Supprimez l'élément de l'inventaire
		delete(char1.inventory, item)
	} else {
		// L'élément n'existe pas dans l'inventaire, vous pouvez gérer ce cas si nécessaire.
	}
}
