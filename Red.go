package main

import (
	"fmt"
	"os"
)

//mathias le looser

func main() {
	inventory := map[string]string{"Arme": "Épée", "Armure": "Légère"}
	P1 := Init("Mat", "Elfe", 1, 100, 40, inventory)

	var choix string

	for choix != "3" {
		fmt.Println("Menu :")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Quitter")

		fmt.Print("Choisissez une option (1/2/3) : ")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			fmt.Println("Affichage des informations du personnage...")
			P1.displayInfo()
		case "2":
			fmt.Println("Accès au contenu de l'inventaire...")
			P1.accessInventory()
		case "3":
			fmt.Println("Merci d'avoir joué, a bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (1/2/3).")
		}

		fmt.Println()
	}
}

type Character struct {
	nickname   string
	classe     string
	lvl        int
	hp_max     int
	current_hp int
	inventory  map[string]string
}

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, inventory map[string]string) Character {
	Character := Character{
		nickname:   nickname,
		classe:     classe,
		lvl:        lvl,
		hp_max:     hp_max,
		current_hp: current_hp,
		inventory:  inventory,
	}
	return Character
}
func (c *Character) displayInfo() {
	fmt.Printf("\n Nickname: %s \n Class: %s \n Level: %d \n Hp_Max : %d \n Current_Hp : %d  \n",
		c.nickname, c.classe, c.lvl,
		c.hp_max, c.current_hp)

}

func (inv *Character) accessInventory() {
	for i, value := range inv.inventory {
		fmt.Printf("\n %s : %s \n", i, value)
	}
}

func takePot() {

}

func dead(c *Character) {

}
