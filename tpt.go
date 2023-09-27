package red

import (
	"fmt"
	"math/rand"
)

func (c *Character) TPT(m *Monstre) {
	if c.current_hp <= 0 {
		Defil("Vous êtes mort !\n")
		return
	} else {
		for i := 1; i <= 20; i++ {
			fmt.Println("Tour", i)
			fmt.Println("")
			if i%2 == 0 {
				Defil("L'ennemi attaque !\n")
				if i%3 == 0 {
					Defil("L'ennemi envoie une attaque chargée !\n")
					c.current_hp -= 10
					Defil("Vous avez perdu 5 HP\n")
					c.Dead()
					continue
				}
				c.current_hp -= 5
				Defil("Vous avez perdu 5 HP\n")
				c.Dead()
			} else {
				Defil("On attaque !\n")
				Defil("\nChoisissez votre attaque : \n")
				Defil("1. Attaque basique\n")
				DefilAS("2. Attaque spécifique des ", &P1)
				fmt.Println("")
				Defil("3. Ouvrir l'inventaire")
				var choix string
				fmt.Scan(&choix)
				switch choix {
				case "1":
					c.AttaqueBasique(m)
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						return
					}
				case "2":
					if Esquive() {
						Defil("Le monstre a esquivé l'attaque !\n")
						continue
					} else {
						c.AttaqueSpecifique(m)
						if m.curpv <= 0 {
							Defil("Le monstre est mort !\n")
							return
						}
					}
				case "3":
					if c.inventory == nil {
						fmt.Println("Votre inventaire est vide ")
						c.AttaqueBasique(m)
						if m.curpv <= 0 {
							Defil("Le monstre est mort !\n")
							return
						}
					}
					c.displayInventory()
					fmt.Println("")

				default:
					Defil("Choix invalide, utilisez l'attaque basique.\n")
					c.AttaqueBasique(m)
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						return
					}
				}
			}
		}
	}
}

// AttaqueBasique effectue l'attaque basique commune à toutes les classes.
func (c *Character) AttaqueBasique(m *Monstre) {
	Defil("Attaque basique !\n")
	m.curpv -= 5
}

// AttaqueSpecifique effectue l'attaque spécifique en fonction de la classe.
func (c *Character) AttaqueSpecifique(m *Monstre) {
	fmt.Printf("Attaque spécifique de la classe des %s !\n", c.classe)
	switch c.classe {
	case "Nains":
		// Utilisez l'attaque spécifique des nains
		Defil("Frappe Sismique !\n")
		m.curpv -= 15

	case "Elfes":
		// Utilisez l'attaque spécifique des elfes
		Defil("Tir de Précision !\n")
		m.curpv -= 15

	case "Humains":
		// Utilisez l'attaque spécifique des humains
		Defil("Stratégie Tactique !\n")
		m.curpv -= 15

	default:
		Defil("Classe invalide, utilisez l'attaque basique.\n")
		c.AttaqueBasique(m)

		if m.curpv <= 0 {
			Defil("Le monstre n'a plus de point de vie!\n")
			return
		}
	}
}

func Esquive() bool {
	esquive := rand.Intn(5)
	return esquive < 2
}

func (c *Character) TPTLoup(m *Monstre) {

	for i := 1; i <= 20; i++ {
		fmt.Println("")
		fmt.Println("Tour", i)
		fmt.Println("")

		if i%2 == 0 {
			Defil("L'ennemi attaque !\n")

			c.current_hp -= m.damagept
			DefilDMG("Vous avez perdu ", &M3, " HP\n")
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
				DefilLeft("Il lui reste ", &M3, " PV\n")
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					return
				}
			case "2":
				if Esquive() {
					Defil("Le monstre a esquivé l'attaque !\n")
					continue
				} else {
					c.AttaqueSpecifique(m)
					DefilLeft("Il lui reste ", &M3, " PV\n")
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						return
					}
				}

			default:
				Defil("Choix invalide, utilisez l'attaque basique.\n")
				c.AttaqueBasique(m)
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					return
				}
			}
		}
	}
}
