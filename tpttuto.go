package red

import (
	"fmt"
)

func (c *Character) TPTtuto(m *Monstre) {
	Defil("\nUn Gobelin apparait soudainement devant vous !\n")
	Defil("\nVous allez devoir le combattre pour survivre !\n")
	fmt.Println("")
	Defil("Attaquons le !")
	fmt.Println("")

	for i := 1; i <= 20; i++ {
		fmt.Println("")
		fmt.Println("Tour", i)
		fmt.Println("")
		if i == 1 {
			c.Attaque1(&M1)
			i++
			fmt.Println("")
			fmt.Println("Tour", i)
		} else if i == 3 {
			Defil("Vous avez une autre attaque que vous pouvez utiliser pour vaincre ce vilain garnement... Elle est plus puissante mais le monstre a des chances de l'esquiver\n")
			c.Attaque3(&M1)
			i++
			Defil("Maintenant que vous connaissez vos attaque nous allons pouvoir l'éliminer !\n")
			fmt.Println("")
			fmt.Println("Tour", i)

		}

		if i%2 == 0 {
			Defil("L'ennemi attaque !\n")
			if i%3 == 0 {
				Defil("L'ennemi envoie une attaque chargée !\n")
				c.current_hp -= 4
				Defil("Vous avez perdu 4 HP\n")
				c.Dead()
				continue
			}
			c.current_hp -= 2
			Defil("Vous avez perdu 2 HP\n")
			c.Dead()
		} else {
			Defil("On attaque !\n")
			Defil("\nChoisissez votre attaque : \n")
			Defil("1. Attaque basique\n")
			DefilAS("2. Attaque spécifique des ", &P1)
			fmt.Println("")
			var choix string
			fmt.Scan(&choix)
			switch choix {
			case "1":
				c.AttaqueBasique(m)
				DefilLeft("Il lui reste ", &M1, " PV\n")
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					Defil("Vous avez gagné 2500 XP ! Vous montez de niveau tous les 5000 point d'EXP, battez plus de créatures pour grimper en niveaux\n")
					c.EXP += 2500
					return
				}
			case "2":
				if Esquive() {
					Defil("Le monstre a esquivé l'attaque !\n")
					continue
				} else {
					c.AttaqueSpecifique(m)
					DefilLeft("Il lui reste ", &M1, " PV\n")
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						Defil("Vous avez gagné 2500 XP ! Vous montez de niveau tous les 5000 point d'EXP, battez plus de créatures pour grimper en niveaux\n")
						c.EXP += 2500

						return
					}
				}
			case "3":
				c.HandleInventory()
			case "SLAY":
				Slay(&M1)

			default:
				Defil("Choix invalide, utilisez l'attaque basique.\n")
				c.AttaqueBasique(m)
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					Defil("Vous avez gagné 2500 XP ! Vous montez de niveau tous les 5000 point d'EXP, battez plus de créatures pour grimper en niveaux\n")
					c.EXP += 2500
					return
				}
			}
		}
	}
}

func (c *Character) Attaque1(m *Monstre) {
	var choix1 string
	Defil("Appuyez sur la touche 1 de votre clavier pour lancer une attaque basique")
	fmt.Scanln(&choix1)
	switch choix1 {
	case "1":
		Defil("Bien joué !")
		c.AttaqueBasique(m)
	default:
		Defil("Touche Incorecte, Veuillez appuyer sur la touche 1 pour attaquer le Gobelin\n")
		c.Attaque1(&M1)
	}
}

func (c *Character) Attaque3(m *Monstre) {
	var choix1 string
	Defil("Appuyez sur la touche 2 de votre clavier pour lancer une attaque spécifique de classe")
	fmt.Scanln(&choix1)
	switch choix1 {
	case "2":
		Defil("Bien joué !")
		c.AttaqueSpecifique(m)
	default:
		Defil("Touche Incorecte, Veuillez appuyer sur la touche 2 pour attaquer le Gobelin\n")
		c.Attaque3(&M1)
	}
}

func (c *Character) HandleInventory() {
	fmt.Println("Inventaire:")
	availableItems := []string{"Potion de Soin", "Potion de Poison", "Potion de Mana"}

	// Vérifiez si l'inventaire est vide
	inventoryEmpty := true
	for _, item := range availableItems {
		quantity, exists := c.inventory[item]
		if exists && quantity > 0 {
			inventoryEmpty = false
			fmt.Printf("%s %d disponible(s)\n", item, quantity)
		}
	}

	if inventoryEmpty {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	fmt.Printf("Laquelle voulez-vous utiliser (entrez le numéro) ?")
	var test int
	n, err := fmt.Scan(&test)
	if err != nil {
		fmt.Println(err)
	}
	switch n {
	case 1:
		if c.inventory["Potion de Soin"] > 0 {
			itemToUse := "Potion de Soin"
			UseItem(c, itemToUse, &Monstre{})
		} else {
			fmt.Println("Vous n'avez pas de Potion de Soin dans l'inventaire.")
		}
	case 2:
		if c.inventory["Potion de Poison"] > 0 {
			itemToUse := "Potion de Poison"
			UseItem(c, itemToUse, &Monstre{})
		} else {
			fmt.Println("Vous n'avez pas de Potion de Poison dans l'inventaire.")
		}
	case 3:
		if c.inventory["Potion de Mana"] > 0 {
			itemToUse := "Potion de Mana"
			UseItem(c, itemToUse, &Monstre{})
		} else {
			fmt.Println("Vous n'avez pas de Potion de Mana dans l'inventaire.")
		}
	default:
		fmt.Println("Choix invalide.")
	}
}
