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
		}
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

func (c *Character) Attaque1(m *Monstre) {
	var choix1 string
	Defil("Appuyez sur la touche 1 de votre clavier pour lancer une attaque basique")
	fmt.Scanln(&choix1)
	switch choix1 {
	case "1":
		Defil("Bien joué !")
		c.AttaqueBasique(m)
		DefilLeft("Il lui reste ", &M1, " Point de vie\n")
	default:
		Defil("Touche Incorecte, Veuillez appuyer sur la touche 1 pour attaquer le Gobelin\n")
		c.Attaque1(&M1)
	}
}
