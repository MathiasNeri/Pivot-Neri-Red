package red

import (
	"fmt"
)

var leninv = 10
var stock int = 2 // À modifier lorsque vous pourrez créer des équipements

func (c *Character) Dead() {
	if c.current_hp <= 0 {
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

func (char1 *Character) ShowEquipement() { // Permet de print la liste de l'équipement équipé
	fmt.Println("Tête :", char1.Equipement.Tete)
	fmt.Println("Torse :", char1.Equipement.Torse)
	fmt.Println("Bottes :", char1.Equipement.Bottes)
}

func ContainsKey(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

func (char1 *Character) EquipHead(s string) bool { // Permet d'équiper la tête
	if ContainsKey(char1.inventory, s) {
		char1.Equipement.Tete = s
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous venez d'équiper : ", s)
		char1.RemoveItem(s)
		return true
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Vous n'avez rien à équiper")
	return false
}

func (char1 *Character) EquipChest(s string) bool { // Permet d'équiper le torse
	if ContainsKey(char1.inventory, s) {
		char1.Equipement.Torse = s
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous venez d'équiper : ", s)
		char1.RemoveItem(s)
		return true
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Vous n'avez rien à équiper")
	return false
}

var chapeauAventurier_mis = 0
var tuniqueAventurier_mis = 0
var bottesAventurier_mis = 0

func (char1 *Character) EquipBoots(s string) bool { // Permet d'équiper les bottes
	if ContainsKey(char1.inventory, s) {
		char1.Equipement.Bottes = s
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		fmt.Println("Vous venez d'équiper : ", s)
		char1.RemoveItem(s)
		return true
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Vous n'avez rien à équiper")
	return false
}

func menuEquipement(char1 *Character) {
	char1.ShowEquipement()
	fmt.Println("-------------------EQUIPEMENT--------------------")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Que souhaitez-vous équiper")
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	if chapeauAventurier_mis < 1 {
		fmt.Println("1 = Équiper un chapeau de l'aventurier (+10 PV max)")
	}
	fmt.Println()
	if tuniqueAventurier_mis < 1 {
		fmt.Println("2 = Équiper une tunique de l'aventurier (+25 PV max)")
	}
	fmt.Println()
	if bottesAventurier_mis < 1 {
		fmt.Println("3 = Équiper des bottes de l'aventurier (+15 PV max)")
	}
	fmt.Println()
	fmt.Println("0 = Quitter le menu de l'équipement")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		if char1.EquipHead(chapeauAventurier) {
			fmt.Println("Vous avez équipé un chapeau de l'aventurier (+10 PV max)")
			char1.Update_Pv_Max()
			chapeauAventurier_mis += 1
		}
	case "2":
		if char1.EquipChest(tuniqueAventurier) {
			fmt.Println("Vous avez équipé une tunique de l'aventurier (+25 PV max)")
			char1.Update_Pv_Max()
			tuniqueAventurier_mis += 1
		}
	case "3":
		if char1.EquipBoots(bottesAventurier) {
			fmt.Println("Vous avez équipé des bottes de l'aventurier (+15 PV max)")
			char1.Update_Pv_Max()
			bottesAventurier_mis += 1
		}
	}
}

func (char1 *Character) Update_Pv_Max() {
	if char1.Equipement.Tete == chapeauAventurier {
		char1.hp_max += 10
	}
	if char1.Equipement.Torse == tuniqueAventurier {
		char1.hp_max += 25
	}
	if char1.Equipement.Bottes == bottesAventurier {
		char1.hp_max += 15
	}
}
