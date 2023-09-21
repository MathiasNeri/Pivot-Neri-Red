package red

import (
	"fmt"
)

// pense a les remettre a 0 quand tu les perd clampin

var Chapeauvendu int
var Tuniquevendu int
var Bottevendu int

func (c *Character) Forgeron() {
	fmt.Println("Quel item ouhaitez vous fabriquer ?")
	fmt.Println("0. Revenir au menu")
	if Chapeauvendu < 1 {
		fmt.Println("1. Chapeau de l'aventurier (X1 Plume de Corbeau + X1 Cuir de Sanglier) et 5 pièces d'or")
	}
	if Tuniquevendu < 1 {
		fmt.Println("2. Tunique de l'aventurier (X1 Fourrure de Loup + X1 Peau de Troll) et 5 pièces d'or")
	}
	if Bottevendu < 1 {
		fmt.Println("3. Botte de l'aventurier (X1 Fourrure de Loup + X1 Cuir de Sanglier) et 5 pièces d'or")
	}

	var choix string
	fmt.Scan(&choix)
	switch choix {
	case "0":
		Menu()
	case "1":
		if !c.LimitInv() {
			fmt.Println("Votre inventaire est plein !")
			return
		}
		var stockc int
		for i := range c.inventory {
			if i == "Plume de Corbeau" {
				stockc += 1
			}
			if i == "Cuir de Sanglier" {
				stockc += 1
			}
		}
		if stockc == 2 {
			c.money -= 5
			fmt.Println("Vous venez de craft le Chapeau de l'aventurier !")
			c.inventory["Chapeau de l'aventurier"]++
			stock++
			Chapeauvendu++
			delete(c.inventory, "PLume de Corbeau")
			stock--
			delete(c.inventory, "Cuir de Sanglier")
			stock--
			return
		} else {
			fmt.Println("Vous n'avez pas les objets requis pour carft l'item")

		}
	case "2":
		if !c.LimitInv() {
			fmt.Println("Votre inventaire est plein !")
			return
		}
		var stockt int
		for i := range c.inventory {
			if i == "Fourrure de Loup" {
				stockt += 1
			}
			if i == "Peau de Troll" {
				stockt += 1
			}
		}
		if stockt == 2 {
			c.money -= 5
			fmt.Println("Vous venez de fabriquer la Tunique de l'aventurier !")
			c.inventory["Tunique de l'aventurier"]++
			stock++
			Tuniquevendu++
			delete(c.inventory, "Fourrure de loup")
			stock--
			delete(c.inventory, "Peau de Troll")
			stock--
			stockt = 0

		} else {
			fmt.Println("Vous n'avez pas les objets requis pour carft l'item")
		}
	case "3":
		if !c.LimitInv() {
			fmt.Println("Votre inventaire est plein !")
			return
		}
		var stockb int
		for i := range c.inventory {
			if i == "Fourrure de Loup" {
				stockb += 1
			}
			if i == "Cuir de Sanglier" {
				stockb += 1
			}
		}
		if stockb == 2 {
			c.money -= 5
			fmt.Println("Vous venez de craft les bottes de l'aventurier !")
			c.inventory["Bottes de l'aventurier"]++
			stock++
			Bottevendu++
			delete(c.inventory, "Fourrure de Loup")
			stock--
			delete(c.inventory, "Cuir de Sanglier")
			stock--
			return
		} else {
			fmt.Println("Vous n'avez pas les objets requis pour carft l'item")

		}
	default:
		fmt.Println("Article invalide. Veuillez choisir un article valide.")

	}
}
