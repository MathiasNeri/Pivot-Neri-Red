package red

import "fmt"

func (c *Character) spellBook() {
	for _, i := range c.skill {
		if i == "Livre de Sort : Boule de Feu" {
			fmt.Println("Vous maitrisez déjà la compétence Boule de Feu !")
			return
		}
	}
	for k := range c.inventory {
		if k == "Livre de Sort : Boule de Feu" {
			c.skill = append(c.skill, "Livre de Sort : Boule de Feu")
			c.inventory["Livre de Sort : Boule de Feu"]--
			stock--
			LivreDeSortBDF++
			return
		}
	}
	fmt.Println("Pour apprendre la compétence Boule de Feu, il faut d'abord l'acheter chez le marchand.")
}
