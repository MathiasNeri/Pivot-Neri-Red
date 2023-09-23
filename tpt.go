package red

import "fmt"

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
				fmt.Println("\nChoisissez votre attaque : ")
				fmt.Println("1. Attaque basique")
				fmt.Printf("2. Attaque spécifique (%s)\n", c.classe)
				var choix string
				fmt.Scan(&choix)
				switch choix {
				case "1":
					c.AttaqueBasique(m)
				case "2":
					c.AttaqueSpecifique(m)
				default:
					fmt.Println("Choix invalide, utilisez l'attaque basique.")
					c.AttaqueBasique(m)
				}
			}
		}
	}
}

// AttaqueBasique effectue l'attaque basique commune à toutes les classes.
func (c *Character) AttaqueBasique(m *Monstre) {
	fmt.Println("Attaque basique !")
	m.curpv -= 10
	if m.curpv <= 0 {
		fmt.Println("Le monstre est mort !")
		return
	}
	fmt.Println("Il lui reste", m.curpv, "points de vie.")
}

// AttaqueSpecifique effectue l'attaque spécifique en fonction de la classe.
func (c *Character) AttaqueSpecifique(m *Monstre) {
	fmt.Printf("Attaque spécifique de la classe %s !\n", c.classe)
	switch c.classe {
	case "Nain":
		// Utilisez l'attaque spécifique des nains
		fmt.Println("Frappe Sismique !")
		m.curpv -= 15
		if m.curpv <= 0 {
			fmt.Println("Le monstre est mort !")
			return
		}
		fmt.Println("Il lui reste", m.curpv, "points de vie.")
	case "Elfe":
		// Utilisez l'attaque spécifique des elfes
		fmt.Println("Tir de Précision !")
		m.curpv -= 15
		if m.curpv <= 0 {
			fmt.Println("Le monstre est mort !")
			return
		}
		fmt.Println("Il lui reste", m.curpv, "points de vie.")
	case "Humain":
		// Utilisez l'attaque spécifique des humains
		fmt.Println("Stratégie Tactique !")
		m.curpv -= 15
		if m.curpv <= 0 {
			fmt.Println("Le monstre est mort !")
			return
		}
		fmt.Println("Il lui reste", m.curpv, "points de vie.")
	default:
		fmt.Println("Classe invalide, utilisez l'attaque basique.")
		c.AttaqueBasique(m)
	}
}
