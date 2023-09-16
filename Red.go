package main

import (
	"fmt"
	"os"
)

func main() {
<<<<<<< HEAD
	inventory := map[string]int{"Arme": 1, "Armure": 1} // Utilisez un map[string]int pour compter le nombre d'objets de chaque type
	P1 := Init("Mat", "Elfe", 1, 100, 40, inventory)
=======
	inventory := map[string]string{"Arme": "Épée", "Armure": "Légère"}
	skill := []string{"Coup de Poing"}
	P1 := Init("Mat", "Elfe", 1, 100, 40, inventory, skill)
>>>>>>> 47a1d08f801fa7826e496a7dc3f02d8e2338bd64

	var choix string

	for choix != "0" {
		fmt.Println("Menu :")
		fmt.Println("0. Quitter")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Vérifier si je suis mort")
		fmt.Println("4. Marchand")
		fmt.Println("5. Prendre Potion")

		fmt.Print("Choisissez une option (1/2/3/4/5) : ")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			fmt.Println("Affichage des informations du personnage...")
			P1.displayInfo()
		case "2":
			fmt.Println("Accès au contenu de l'inventaire...")
			P1.accessInventory()
		case "3":
			fmt.Println("Vérifier si je suis mort...")
			P1.dead()
		case "4":
			fmt.Println("Bienvenue chez le marchand !")
			marchand(&P1)
		case "5":
			fmt.Println("Gloup, Gloup, Gloup, Gloup,...")
			P1.takePot()
		case "0":
			fmt.Println("Merci d'avoir joué, à bientôt !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (1/2/3/4/5).")
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
<<<<<<< HEAD
	inventory  map[string]int // Utilisez un map[string]int pour compter le nombre d'objets de chaque type
}

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, inventory map[string]int) Character {
=======
	inventory  map[string]string
	skill      []string
}

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, inventory map[string]string, skill []string) Character { //Init du perso
>>>>>>> 47a1d08f801fa7826e496a7dc3f02d8e2338bd64
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
<<<<<<< HEAD

func (c *Character) displayInfo() {
	fmt.Printf("\n Nickname: %s \n Class: %s \n Level: %d \n Hp_Max : %d \n Current_Hp : %d  \n",
		c.nickname, c.classe, c.lvl,
		c.hp_max, c.current_hp)
=======
func (c *Character) displayInfo() { //Affiche les infos du perso
	fmt.Printf("\n Nickname: %s \n Class: %s \n Level: %d \n Hp_Max : %d \n Current_Hp : %d  \n Skill : %s",
		c.nickname, c.classe, c.lvl,
		c.hp_max, c.current_hp, c.skill)

>>>>>>> 47a1d08f801fa7826e496a7dc3f02d8e2338bd64
}

func (c *Character) accessInventory() {
	fmt.Println("Inventaire :")
	for item, count := range c.inventory {
		fmt.Printf("%s : %d\n", item, count)
	}
}

func (c *Character) takePot() {
	if c.inventory["Potion de Soin"] > 0 {
		c.inventory["Potion de Soin"]--
		pointsDeSoins := 50
		c.current_hp += pointsDeSoins

		if c.current_hp > c.hp_max {
			c.current_hp = c.hp_max
		}

		fmt.Printf("Vous avez utilisé une Potion de Soin et avez récupéré %d points de vie.\n", pointsDeSoins)
		fmt.Printf("Points de vie actuels : %d / %d\n", c.current_hp, c.hp_max)
	} else {
		fmt.Println("Vous n'avez pas de Potion de Soin dans l'inventaire.")
	}
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

func marchand(c *Character) {
	if c.inventory["Potion de Soin"] == 0 {
		fmt.Println("Articles disponibles chez le marchand :")
		fmt.Println("1. Potion de Soin (gratuitement)")

		var choix string
		fmt.Print("Choisissez un article (1 pour la Potion de Soin) : ")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			c.inventory["Potion de Soin"]++ // Ajoutez la Potion de Soin à l'inventaire du personnage
			fmt.Println("Vous avez acheté une Potion de Soin et elle a été ajoutée à votre inventaire.")
		default:
			fmt.Println("Article invalide. Veuillez choisir un article valide.")
		}
	} else {
		fmt.Println("Le marchand n'a plus d'articles à vendre.")
	}
}
