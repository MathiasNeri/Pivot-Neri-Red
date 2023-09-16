package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	inventory := map[string]int{"Arme": 1, "Armure": 1} // Utilisez un map[string]int pour compter le nombre d'objets de chaque type
	P1 := Init("Mat", "Elfe", 1, 100, 40, inventory, []string{"Coup de poing"})

	var choix string

	for choix != "0" {
		fmt.Println("Menu :")
		fmt.Println("0. Quitter")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Vérifier si je suis mort")
		fmt.Println("4. Marchand")
		fmt.Println("5. Prendre Potion")
		fmt.Println("6. Apprendre Boule de Feu ")
		fmt.Print("Choisissez une option (0/1/2/3/4/5/6) : ")
		fmt.Scanln(&choix)

		switch choix {
		case "0":
			fmt.Println("Merci d'avoir joué, a bientôt !")
			os.Exit(0)
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
		case "6":
			fmt.Println("Apprendre la compétence Boule de Feu")
			P1.spellBook()

		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (0/1/2/3/4/5/6).")

			fmt.Println()
		}
	}
}

type Character struct { // creation de la classe Character
	nickname   string
	classe     string
	lvl        int
	hp_max     int
	current_hp int
	inventory  map[string]int // Utilisez un map[string]int pour compter le nombre d'objets de chaque type
	skill      []string
}

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, inventory map[string]int, skill []string) Character { //Init du perso
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

func (c *Character) accessInventory() {
	fmt.Println("Inventaire :")
	for item, count := range c.inventory {
		fmt.Printf("%s : %d\n", item, count)
	}
}

func (c *Character) takePot() {
	fmt.Println("Quelle potion souhaitez-vous prendre ?")
	fmt.Println("1. Potion de Soin")
	fmt.Println("2. Potion de Poison")
	var choix string
	fmt.Print("Choisissez une option (1/2) : ")
	fmt.Scanln(&choix)

	switch choix {
	case "1":
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
	case "2":
		if c.inventory["Potion de Poison"] > 0 {
			c.inventory["Potion de Poison"]--
			poisonDuration := 3 * time.Second
			damageInterval := 1 * time.Second
			damagePerTick := 10

			fmt.Println("Vous avez été empoisonné !")

			ticker := time.NewTicker(damageInterval)
			defer ticker.Stop()

			damageTimer := time.NewTimer(poisonDuration)
			defer damageTimer.Stop()

			for {
				select {
				case <-ticker.C:
					c.current_hp -= damagePerTick
					if c.current_hp < 0 {
						c.current_hp = 0
					}
					fmt.Printf("Points de vie actuels : %d / %d\n", c.current_hp, c.hp_max)
				case <-damageTimer.C:
					fmt.Println("L'effet de poison est terminé.")
					return
				}
			}
		} else {
			fmt.Println("Vous n'avez pas de Potion de Poison dans l'inventaire.")
		}
	default:
		fmt.Println("Choix invalide. Veuillez choisir une option valide (1/2).")
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

var potionsDeSoinVendues = 0
var potionsDePoisonVendues = 0

func marchand(c *Character) {
	if potionsDeSoinVendues < 1 || potionsDePoisonVendues < 1 {
		fmt.Println("Articles disponibles chez le marchand :")
		if potionsDeSoinVendues < 1 {
			fmt.Println("1. Potion de Soin (gratuitement)")
		}
		if potionsDePoisonVendues < 1 {
			fmt.Println("2. Potion de Poison (gratuitement)")
		}

		var choix string
		fmt.Print("Choisissez un article : \n")
		if potionsDeSoinVendues < 1 {
			fmt.Print("1 pour la Potion de Soin\n")
		}
		if potionsDePoisonVendues < 1 {
			fmt.Print("2 pour la Potion de Poison\n")
		}
		fmt.Print(": ")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			if potionsDeSoinVendues < 1 {
				c.inventory["Potion de Soin"]++
				potionsDeSoinVendues++
				fmt.Println("Vous avez acheté une Potion de Soin et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Soin à vendre.")
			}
		case "2":
			if potionsDePoisonVendues < 1 {
				c.inventory["Potion de Poison"]++
				potionsDePoisonVendues++
				fmt.Println("Vous avez acheté une Potion de Poison et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Poison à vendre.")
			}
		default:
			fmt.Println("Article invalide. Veuillez choisir un article valide.")
		}
	} else {
		fmt.Println("Le marchand n'a plus d'articles à vendre.")
	}
}

func (c *Character) spellBook() {
	for _, i := range c.skill {
		if i == "Livre de Sort : Boule de Feu" {
			fmt.Println("Vous maitrisez deja la compétence Boule de Feu !")
			return
		}
	}
	c.skill = append(c.skill, "Livre de Sort : Boule de Feu")
	fmt.Println("Vous venez d'apprendre la compétence Boule de Feu !")
}
