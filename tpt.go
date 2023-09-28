package red

import (
	"fmt"
	"math/rand"
)

func (c *Character) LevelUp() {
	c.lvl++
	c.hp_max += 10
	c.current_hp = c.hp_max
}

func (c *Character) TPT(m *Monstre) {
	compteur_bdf := 0
	if c.current_hp <= 0 {
		Defil("Vous êtes mort !\n")
		return
	} else {
		for i := 1; i <= 20; i++ {
			fmt.Println("Tour", i)
			fmt.Println("")
			if compteur_bdf%5 == 0 {
				// Check if the character has "Boule de Feu" skill
				if c.HasSkill("Boule de Feu") {
					fmt.Println("4. Utiliser Boule de Feu")
				}

				if i%2 == 0 {
					Defil("L'ennemi attaque !\n")
					if i%3 == 0 {
						Defil("L'ennemi envoie une attaque chargée !\n")
						c.current_hp -= 10
						Defil("Vous avez perdu 5 PV\n")
						c.Dead()
						continue
					}
					c.current_hp -= 5
					Defil("Vous avez perdu 5 PV\n")
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
							Defil("Vous avez gagné 2500 XP ! \n")
							c.EXP += 2500
							if c.EXP >= 5000 {
								Defil("Vous passez au supérieur ! \n")
								c.LevelUp()
								c.EXP -= 5000
								return

							}
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
								Defil("Vous avez gagné 2500 XP ! \n")
								c.EXP += 2500
								if c.EXP >= 5000 {
									Defil("Vous passez au supérieur ! \n")
									c.LevelUp()
									c.EXP -= 5000
									return

								}
								return
							}
						}
					case "3":
						fmt.Println("Ouvrir l'Inventaire :")
						c.HandleInventory()

					case "4":
						if c.HasSkill("Boule de Feu") {
							// Call a function to use the skill here
							c.UseBouleDeFeu(m)
							compteur_bdf = 0
							if m.curpv <= 0 {
								Defil("Le monstre est mort !\n")
								Defil("Vous avez gagné 2500 XP ! \n")
								c.EXP += 2500
								if c.EXP >= 5000 {
									Defil("Vous passez au supérieur ! \n")
									c.LevelUp()
									c.EXP -= 5000
									return

								}
								return
							}
						} else {
							fmt.Println("Vous ne connaissez pas la compétence Boule de Feu.")
						}

					default:
						Defil("Choix invalide, utilisez l'attaque basique.\n")
						c.AttaqueBasique(m)
						if m.curpv <= 0 {
							Defil("Le monstre est mort !\n")
							c.EXP += 2500
							if c.EXP >= 5000 {
								Defil("Vous passez au supérieur ! \n")
								c.LevelUp()
								c.EXP -= 5000
								return

							}
							return

						}
					}
				}
			}
			compteur_bdf++
		}
	}
}

func (c *Character) UseBouleDeFeu(m *Monstre) {
	// Implement the logic for using Boule de Feu here
	// You can deduct the required resources and apply the skill's effects
	// Example:
	fmt.Println("Vous utilisez Boule de Feu contre le monstre!")
	dmg_bouleDeFeu := 30
	m.curpv -= dmg_bouleDeFeu
	c.mana -= 20
	fmt.Printf("Vous lui infligez %d de dégâts.\n", dmg_bouleDeFeu)
}

func (c *Character) HasSkill(skillName string) bool {
	for _, s := range c.skill {
		if s == skillName {
			return true
		}
	}
	return false
}

// AttaqueBasique effectue l'attaque basique commune à toutes les classes.
func (c *Character) AttaqueBasique(m *Monstre) {
	Defil("Attaque basique !\n")
	dmg_attaquebasique := 5
	m.curpv -= dmg_attaquebasique
	fmt.Printf("Vous lui infligez %d", dmg_attaquebasique)
	fmt.Println("")
	DefilLeft("Il lui reste ", m, " Point de vie\n")
}

// AttaqueSpecifique effectue l'attaque spécifique en fonction de la classe.
func (c *Character) AttaqueSpecifique(m *Monstre) {
	fmt.Printf("Attaque spécifique de la classe des %s !\n", c.classe)
	switch c.classe {
	case "Nains":
		// Utilisez l'attaque spécifique des nains
		Defil("Frappe Sismique !\n")
		dmg_attaquespecifique := 15
		m.curpv -= dmg_attaquespecifique
		fmt.Printf("Vous lui infligez %d", dmg_attaquespecifique)
		fmt.Println("")
		DefilLeft("Il lui reste ", m, " Point de vie\n")
		c.mana -= 10
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
	compteur_bdf := 0
	for i := 1; i <= 20; i++ {
		fmt.Println("")
		fmt.Println("Tour", i)
		fmt.Println("")

		if i%2 == 0 {
			Defil("L'ennemi attaque !\n")

			c.current_hp -= m.damagept
			DefilDMG("Vous avez perdu ", &M3, " PV\n")
			c.Dead()
		} else {
			Defil("On attaque !\n")
			Defil("\nChoisissez votre attaque : \n")
			Defil("1. Attaque basique\n")
			DefilAS("2. Attaque spécifique des ", &P1)
			fmt.Println("")
			Defil("3. Ouvrir l'inventaire\n")
			Defil("4. Utiliser BDF")
			var choix string
			fmt.Scan(&choix)
			switch choix {
			case "1":
				c.AttaqueBasique(m)
				compteur_bdf++
				DefilLeft("Il lui reste ", &M3, " PV\n")
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					c.EXP += 2500
					if c.EXP >= 5000 {
						Defil("Vous passez au supérieur ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return

					}
					return

				}
			case "2":
				if Esquive() {
					Defil("Le monstre a esquivé l'attaque !\n")
					continue
				} else {
					if !P1.Mana10() {
						Defil("\nVous n'avez pas assez de mana pour utiliser cette attaque...")
					}
					c.AttaqueSpecifique(m)
					compteur_bdf++
					DefilLeft("Il lui reste ", &M3, " PV\n")
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						c.EXP += 2500
						if c.EXP >= 5000 {
							Defil("Vous passez au supérieur ! \n")
							c.LevelUp()
							c.EXP -= 5000
							return

						}
						return
					}
				}
			case "3":
				fmt.Println("Ouvrir l'Inventaire :")
				c.HandleInventory()
			case "4":
				if c.HasSkill("Boule de Feu") {
					// Call a function to use the skill here
					c.UseBouleDeFeu(m)
					compteur_bdf = 0
					DefilLeft("Il lui reste ", &M3, " PV\n")
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						c.EXP += 2500
						if c.EXP >= 5000 {
							Defil("Vous passez au supérieur ! \n")
							c.LevelUp()
							c.EXP -= 5000
							return

						}
						return
					}
				} else {
					fmt.Println("Vous ne connaissez pas la compétence Boule de Feu.")
				}

			default:
				Defil("Choix invalide, utilisez l'attaque basique.\n")
				c.AttaqueBasique(m)
				compteur_bdf++
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					c.inventory["Fourrure de Loup"]++
					Defil("\nVous avez récupéré la fourrure du loup et elle a été ajoutée à votre inventaire \n")
					c.EXP += 2500
					if c.EXP >= 5000 {
						Defil("Vous passez au supérieur ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return

					}
					return

				}
			}
		}
	}
}

func (c *Character) TPTTroll(m *Monstre) {
	compteur_bdf := 0
	for i := 1; i <= 20; i++ {
		fmt.Println("")
		fmt.Println("Tour", i)
		fmt.Println("")

		if i%2 == 0 {
			Defil("L'ennemi attaque !\n")

			c.current_hp -= m.damagept
			DefilDMG("Vous avez perdu ", &M4, " PV\n")
			c.Dead()
			continue
		} else {
			Defil("On attaque !\n")
			Defil("\nChoisissez votre attaque : \n")
			Defil("1. Attaque basique\n")
			DefilAS("2. Attaque spécifique des ", &P1)
			fmt.Println("")
			Defil("3. Ouvrir l'inventaire")
			var choix string
			fmt.Scanln(&choix)
			switch choix {
			case "1":
				c.AttaqueBasique(m)
				DefilLeft("Il lui reste ", &M4, " PV\n")
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					c.EXP += 3500
					if c.EXP >= 5000 {
						Defil("Vous passez au niveau 2 ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return
					}
					return

				}
			case "2":
				if Esquive() {
					Defil("Le monstre a esquivé l'attaque !\n")
					continue
				} else {
					c.AttaqueSpecifique(m)
					DefilLeft("Il lui reste ", &M4, " PV\n")
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						c.EXP += 3500
						if c.EXP >= 5000 {
							Defil("Vous passez au niveau 2 ! \n")
							c.LevelUp()
							c.EXP -= 5000
							return
						}
						return

					}
				}
			case "3":
				fmt.Println("Ouvrir l'Inventaire :")
				c.HandleInventory()
			case "4":
				if c.HasSkill("Boule de Feu") {
					// Call a function to use the skill here
					c.UseBouleDeFeu(m)
					compteur_bdf = 0
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						c.EXP += 3500
						if c.EXP >= 5000 {
							Defil("Vous passez au niveau 2 ! \n")
							c.LevelUp()
							c.EXP -= 5000
							return
						}
						return
					}
				} else {
					fmt.Println("Vous ne connaissez pas la compétence Boule de Feu.")
				}

			default:
				Defil("Choix invalide, utilisez l'attaque basique.\n")
				c.AttaqueBasique(m)
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					c.inventory["Peau de Troll"]++
					Defil("\nVous avez récupéré la peau du Troll et elle a été ajoutée à votre inventaire \n")
					c.EXP += 3500
					if c.EXP >= 5000 {
						Defil("Vous passez au niveau 2 ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return
					}
					return

				}
			}
			compteur_bdf += 1
		}
	}
}

func (c *Character) TPTDragon(m *Monstre) {
	compteur_bdf := 0
	for i := 1; i <= 30; i++ {
		fmt.Println("")
		fmt.Println("Tour", i)
		fmt.Println("")

		if i%2 == 0 {
			Defil("L'ennemi attaque !\n")

			c.current_hp -= m.damagept
			DefilDMG("Vous avez perdu ", &M5, " PV\n")
			c.Dead()
			continue
		} else {
			Defil("On attaque !\n")
			Defil("\nChoisissez votre attaque : \n")
			Defil("1. Attaque basique\n")
			DefilAS("2. Attaque spécifique des ", &P1)
			fmt.Println("")
			Defil("3. Ouvrir l'inventaire")
			var choix string
			fmt.Scanln(&choix)
			switch choix {
			case "1":
				c.AttaqueBasique(m)
				DefilLeft("Il lui reste ", &M5, " PV\n")
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					c.EXP += 4500
					if c.EXP >= 5000 {
						Defil("Vous passez au niveau 2 ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return
					}
					return
				}
			case "2":
				if Esquive() {
					Defil("Le monstre a esquivé l'attaque !\n")
					continue
				} else {
					c.AttaqueSpecifique(m)
					DefilLeft("Il lui reste ", &M5, " PV\n")
					if m.curpv <= 0 {
						Defil("Le monstre est mort !\n")
						c.EXP += 4500
						if c.EXP >= 5000 {
							Defil("Vous passez au niveau 2 ! \n")
							c.LevelUp()
							c.EXP -= 5000
							return
						}
						return
					}

				}
			case "3":
				fmt.Println("Ouvrir l'Inventaire :")
				c.HandleInventory()
			case "4":
				if c.HasSkill("Boule de Feu") {
					c.UseBouleDeFeu(m)
					compteur_bdf = 0
					Defil("Le monstre est mort !\n")
					c.EXP += 4500
					if c.EXP >= 5000 {
						Defil("Vous passez au niveau 2 ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return
					}
					return

				} else {
					fmt.Println("Vous ne connaissez pas la compétence Boule de Feu.")
				}

			default:
				Defil("Choix invalide, utilisez l'attaque basique.\n")
				c.AttaqueBasique(m)
				if m.curpv <= 0 {
					Defil("Le monstre est mort !\n")
					c.EXP += 4500
					if c.EXP >= 5000 {
						Defil("Vous passez au niveau 2 ! \n")
						c.LevelUp()
						c.EXP -= 5000
						return
					}
					return
				}
			}
			compteur_bdf += 1
		}
	}
}
