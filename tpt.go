package red

import (
	"fmt"
	"math/rand"
)

func (c *Character) TPT(m *Monstre) {

	if c.current_hp <= 0 {
		fmt.Println("Vous êtes mort !")
		return
	} else {
		for i := 1; i <= 20; i++ {
			fmt.Println("Tour", i)
			if i%2 == 0 {
				fmt.Println("L'ennemi attaque !")
				c.current_hp -= 5
				fmt.Println("Vous avez perdu", 5, "HP")
				fmt.Println("Il vous reste", c.current_hp, "HP")
			} else {
				fmt.Println("On attaque !")
				fmt.Println("\n Choisissez votre attaque : ")
				fmt.Println("1. Attaque basique (ca passe a chaque fois !)\n2. Attaque puissante (L'attaque est puissante mais le monstre peut l'esquiver !)")
				var choix string
				fmt.Scan(&choix)
				switch choix {
				case "1":

					fmt.Println("Le gobelin a perdu 5 pv")
					m.curpv -= 5
					if m.curpv <= 0 {
						fmt.Println("Le gobelin est mort!")
						return
					}
					fmt.Println("Il lui reste", m.curpv)
				case "2":
					proba := rand.Intn(5)
					fmt.Println(proba)
					if proba < 2 {
						fmt.Println("Le gobelin a esquivé l'attaque !")
					} else {
						fmt.Println("Le gobelin a perdu 15 pv")
						m.curpv -= 15
						if m.curpv <= 0 {
							fmt.Println("Le gobelin est mort!")
							return
						}
						fmt.Println("Il lui reste", m.curpv)
					}
				}
			}
		}
	}
}
