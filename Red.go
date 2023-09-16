package main

import (
	"fmt"
	"os"
)

//mathias le looser

func main() {
	inventory := map[string]string{"Arme": "Épée", "Armure": "Légère"}
	skill := []string{"Coup de Poing"}
	P1 := Init("Mat", "Elfe", 1, 100, 40, inventory, skill)

	var choix string

	for choix != "4" {
		fmt.Println("Menu :")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Vérifier si je suis mort")
		fmt.Println("4. Quitter")

		fmt.Print("Choisissez une option (1/2/3/4) : ")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			fmt.Println("Affichage des informations du personnage...")
			P1.displayInfo()
		case "2":
			fmt.Println("Accès au contenu de l'inventaire...")
			P1.accessInventory()
			P1.takePot()
		case "3":
			fmt.Println("Vérifier si je suis mort...")
			P1.dead()
		case "4":
			fmt.Println("Merci d'avoir joué, a bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (1/2/3/4).")
		}

		fmt.Println()
	}
}

type Character struct { // creation de la classe Character
	nickname   string
	classe     string
	lvl        int
	hp_max     int
	current_hp int
	inventory  map[string]string
	skill      []string
}

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, inventory map[string]string, skill []string) Character { //Init du perso
	Character := Character{
		nickname:   nickname,
		classe:     classe,
		lvl:        lvl,
		hp_max:     hp_max,
		current_hp: current_hp,
		inventory:  inventory,
		skill:      skill,
	}
	return Character
}
func (c *Character) displayInfo() { //Affiche les infos du perso
	fmt.Printf("\n Nickname: %s \n Class: %s \n Level: %d \n Hp_Max : %d \n Current_Hp : %d  \n Skill : %s",
		c.nickname, c.classe, c.lvl,
		c.hp_max, c.current_hp, c.skill)

}

func (inv *Character) accessInventory() {
	for i, value := range inv.inventory {
		fmt.Printf("\n %s : %s \n", i, value)
	}
}

func (c *Character) takePot() {
	found := c.inventory["Potion"]
	if found == "" {
		fmt.Println("Vous n'avez pas de potion dans l'inventaire.")
		return
	}

	delete(c.inventory, "Potion")
	pointsDeSoins := 50

	c.current_hp += pointsDeSoins

	if c.current_hp > c.hp_max {
		c.current_hp = c.hp_max
	}

	fmt.Printf("Vous avez utilisé une potion et avez récupéré %d points de vie.\n", pointsDeSoins)
	fmt.Printf("Points de vie actuels : %d / %d\n", c.current_hp, c.hp_max)
}

func (c *Character) dead() {
	if c.current_hp == 0 {
		fmt.Println("Vous êtes mort !!")
		c.current_hp += c.hp_max / 2
		fmt.Println("Vous venez de récussité avec", c.current_hp, "PV")
	} else {
		fmt.Println("Il vous reste", c.current_hp, "PV")
	}
}
