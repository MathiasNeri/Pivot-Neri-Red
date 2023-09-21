package red

import (
	"fmt"
)

func (c *Character) TPT(m *Monstre) {
	if c.current_hp <= 0 {
		fmt.Println("Vous êtes mort !")
		return
	} else {
		for i := 1; i <= 20; i++ {
			fmt.Println("Tour", i)
			if i%2 != 0 {
				fmt.Println("L'ennemi attaque !")
				c.current_hp -= 5
			} else {
				fmt.Println("On attaque !")
				fmt.Println("\n choisissez votre attaque : ")
				fmt.Println("1. Attaque basique\n2. Attaque puissante")
				var choix string
				fmt.Scan(&choix)
				switch choix {
				case "1":
					fmt.Println("Le gobelin a perdu 5 pv")
					m.curpv -= 5
					fmt.Println(m.curpv)
				case "2":
					fmt.Println("Le gobelin a perdu 15 pv")
					m.curpv -= 15
					fmt.Println(m.curpv)
				}
			}
		}
	}
}
