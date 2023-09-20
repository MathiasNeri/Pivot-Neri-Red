package red

import (
	"fmt"
)

var leninv = 10

var P1 Character
var stock int = 2 // a modifier quand on pourra craft les equipements

type Equipements struct {
	Nom              string
	PointsDeVieBonus int
	Equipe           bool
}

type Equipement struct {
	Helmet     Equipements
	Chestplate Equipements
	Boots      Equipements
}

type Character struct {
	nickname      string
	classe        string
	lvl           int
	hp_max        int
	current_hp    int
	money         int
	inventory     map[string]int
	skill         []string
	inventoryList []string
	Equipements   Equipement
}

type Monstre struct {
	name     string
	pvmax    int
	curpv    int
	damagept int
}

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, money int, inventory map[string]int, skill []string) Character {
	Character := Character{
		nickname:      nickname,
		classe:        classe,
		lvl:           lvl,
		hp_max:        hp_max,
		current_hp:    current_hp,
		money:         money,
		inventory:     inventory,
		skill:         skill,
		inventoryList: []string{}, // Initialise la liste vide
	}
	return Character
}

func (m *Monstre) InitGobelin() {
	m.name = "Gobelin d'entrainement"
	m.pvmax = 40
	m.curpv = 40
	m.damagept = 5
}

func (c *Character) dead() {
	if c.current_hp == 0 {
		fmt.Println("Vous êtes mort !!")
		c.current_hp += c.hp_max / 2
		fmt.Println("Vous venez de réssusciter avec", c.current_hp, "PV")
	} else {
		fmt.Println("Il vous reste", c.current_hp, "PV")
	}
}

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
			fmt.Println("Vous venez d'apprendre la compétence Boule de Feu !")
			c.inventory["Livre de Sort : Boule de Feu"]--
			stock--
			LivreDeSortBDF++
			return
		}
	}
	fmt.Println("Pour apprendre la compétence Boule de Feu, il faut d'abord l'acheter chez le marchand.")
}

func (c *Character) LimitInv() bool {
	if stock == leninv {
		return false
	} else {
		return true
	}
}
