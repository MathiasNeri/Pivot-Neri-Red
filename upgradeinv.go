package red

import "fmt"

func (c *Character) UpgradeInventory() {
	for i := range c.inventory {
		if i == "Compartiment Supplémentaire" {
			leninv += 10
			fmt.Println("Vous avez 10 emplacements d'item supplémentaire !")
		}
	}
}
