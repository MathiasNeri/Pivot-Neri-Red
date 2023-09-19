package red

import "fmt"

func (c *Character) retirerObjet() {
	if len(c.inventoryList) == 0 {
		fmt.Println("L'inventaire est vide, vous ne pouvez pas retirer d'objet.")
		return
	}

	var numObjet int
	fmt.Print("Numéro de l'objet à retirer : ")
	fmt.Scanln(&numObjet)

	numObjet--

	if numObjet >= 0 && numObjet < len(c.inventoryList) {
		objetARetirer := c.inventoryList[numObjet]
		count := c.inventory[objetARetirer]
		if count > 0 {
			c.inventory[objetARetirer]--
			fmt.Printf("Vous avez retiré un(e) %s de l'inventaire.\n", objetARetirer)
			if c.inventory[objetARetirer] == 0 {
				delete(c.inventory, objetARetirer)
				stock--
			}
			c.displayInventory()
		} else {
			fmt.Printf("Vous n'avez pas de %s dans l'inventaire.\n", objetARetirer)
		}
	} else {
		fmt.Println("Numéro d'objet invalide. Veuillez choisir un numéro valide.")
	}
}
