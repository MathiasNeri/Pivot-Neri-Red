package main

import (
	"fmt"
	"os"
	"time"
)

var P1 Character

func main() {
	var choix string

	for choix != "0" {
		fmt.Println("----------------------------------------------")
		fmt.Println("Menu :")
		fmt.Println("0. Quitter")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder au contenu de l'inventaire")
		fmt.Println("3. Vérifier si je suis mort")
		fmt.Println("4. Marchand")
		fmt.Println("5. Prendre Potion")
		fmt.Println("6. Apprendre Boule de Feu")
		fmt.Println("7. Créer son personnage")
		fmt.Println("8. Retirer un objet de l'inventaire")
		fmt.Println("----------------------------------------------")
		fmt.Print("Choisissez une option (0/1/2/3/4/5/6/7/8) : ")

		fmt.Scanln(&choix)

		switch choix {
		case "0":
			fmt.Println("Merci d'avoir joué, à bientôt !")
			os.Exit(0)
		case "1":
			fmt.Println("Affichage des informations du personnage...")
			P1.displayInfo()
		case "2":
			fmt.Println("Accès au contenu de l'inventaire...")
			P1.displayInventory()
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
		case "7":
			fmt.Println("Création du personnage :")
			P1.charCreation()
		case "8":
			fmt.Println("Retirer un objet de l'inventaire :")
			P1.displayInventory()
			P1.retirerObjet()
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (0/1/2/3/4/5/6/7/8).")
		}
	}
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

func (c *Character) displayInfo() {
	fmt.Printf("\nNickname: %s\nClass: %s\nLevel: %d\nHp_Max : %d\nCurrent_Hp : %d\nMoney : %d\nSkill : %s\n",
		c.nickname, c.classe, c.lvl,
		c.hp_max, c.current_hp, c.money, c.skill)
}

func (c *Character) displayInventory() {
	fmt.Println("Inventaire actuel :")
	c.inventoryList = nil
	i := 1
	for item, count := range c.inventory {
		fmt.Printf("%d. %s : %d\n", i, item, count)
		c.inventoryList = append(c.inventoryList, item)
		i++
	}
}

func (c *Character) retirerObjet() {
	if len(c.inventoryList) == 0 {
		fmt.Println("L'inventaire est vide, vous ne pouvez pas retirer d'objet.")
		return
	}

	var numObjet int
	fmt.Print("Numéro de l'objet à retirer : ")
	fmt.Scanln(&numObjet)

	numObjet--

	if numObjet >= 0 && numObjet < len(c.inventoryList) {
		objetARetirer := c.inventoryList[numObjet]
		count := c.inventory[objetARetirer]
		if count > 0 {
			c.inventory[objetARetirer]--
			fmt.Printf("Vous avez retiré un(e) %s de l'inventaire.\n", objetARetirer)
			if c.inventory[objetARetirer] == 0 {
				delete(c.inventory, objetARetirer)
			}
			c.displayInventory()
		} else {
			fmt.Printf("Vous n'avez pas de %s dans l'inventaire.\n", objetARetirer)
		}
	} else {
		fmt.Println("Numéro d'objet invalide. Veuillez choisir un numéro valide.")
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
var LivreDeSortBDF = 0
var FourrureVendues = 0
var PeauTrollVendue = 0
var CuirSanglierVendue = 0
var PlumeCorbeauVendue = 0

func marchand(c *Character) {
	if potionsDeSoinVendues < 1 || potionsDePoisonVendues < 1 || LivreDeSortBDF < 1 || FourrureVendues < 1 || PeauTrollVendue < 1 || CuirSanglierVendue < 1 || PlumeCorbeauVendue < 1 {
		fmt.Println("Articles disponibles chez le marchand :")
		if potionsDeSoinVendues < 1 {
			fmt.Println("1. Potion de Soin : 3 pièces d'or")
		}
		if potionsDePoisonVendues < 1 {
			fmt.Println("2. Potion de Poison : 6 pièces d'or")
		}
		if LivreDeSortBDF < 1 {
			fmt.Println("3. Livre de Sort : Boule de Feu : 25 pièces d'or")
		}
		if FourrureVendues < 1 {
			fmt.Println("4. Fourrure de Loup : 4 pièces d'or")
		}
		if PeauTrollVendue < 1 {
			fmt.Println("5. Peau de Troll : 7 pièces d'or")
		}
		if CuirSanglierVendue < 1 {
			fmt.Println("6. Cuir de Sanglier : 3 pièces d'or")
		}
		if PlumeCorbeauVendue < 1 {
			fmt.Println("7. Plume de Corbeau : 1 pièce d'or")
		}

		var choix string
		fmt.Println("Choisissez un article :")
		if potionsDeSoinVendues < 1 {
			fmt.Println("1 pour la Potion de Soin")
		}
		if potionsDePoisonVendues < 1 {
			fmt.Println("2 pour la Potion de Poison")
		}
		if LivreDeSortBDF < 1 {
			fmt.Println("3 pour Livre de Sort : Boule de Feu")
		}
		if FourrureVendues < 1 {
			fmt.Println("4. Fourrure de Loup")
		}
		if PeauTrollVendue < 1 {
			fmt.Println("5. Peau de Troll")
		}
		if CuirSanglierVendue < 1 {
			fmt.Println("6. Cuir de Sanglier")
		}
		if PlumeCorbeauVendue < 1 {
			fmt.Println("7. Plume de Corbeau :")
		}

		fmt.Print(": ")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			if potionsDeSoinVendues < 1 {
				c.inventory["Potion de Soin"]++
				potionsDeSoinVendues++
				c.money -= 3
				fmt.Println("Vous avez acheté une Potion de Soin pour 3 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Soin à vendre.")
			}
		case "2":
			if potionsDePoisonVendues < 1 {
				c.inventory["Potion de Poison"]++
				potionsDePoisonVendues++
				c.money -= 6
				fmt.Println("Vous avez acheté une Potion de Poison pour 6 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Poison à vendre.")
			}
		case "3":
			if LivreDeSortBDF < 1 {
				c.inventory["Livre de Sort : Boule de Feu"]++
				LivreDeSortBDF++
				c.money -= 25
				fmt.Println("Vous avez acheté Livre de Sort : Boule de Feu pour 25 pièces d'or et la compétence a été ajoutée à votre inventaire.")
			} else if LivreDeSortBDF >= 1 {
				fmt.Println("Le marchand n'a plus de Potion de Livre de Sort :Boule de Feu à vendre.")
			}
		case "4":
			if FourrureVendues < 1 {
				c.inventory["Fourrure de Loup"]++
				c.money -= 4
				fmt.Println("Vous avez acheté Fourrure de Loup pour 4 pièces d'or et l'item a été ajoutée à votre inventaire.")
			}
		case "5":
			if PeauTrollVendue < 1 {
				c.inventory["Peau de Troll"]++
				c.money -= 7
				fmt.Println("Vous avez acheté Peau de Troll pour 7 pièces d'or et l'item' a été ajoutée à votre inventaire.")
			}
		case "6":
			if CuirSanglierVendue < 1 {
				c.inventory["Cuir de Sanglier"]++
				c.money -= 3
				fmt.Println("Vous avez acheté Cuir de Sanglier pour 3 pièces d'or et l'item a été ajoutée à votre inventaire.")
			}
		case "7":
			if PlumeCorbeauVendue < 1 {
				c.inventory["Plume de Corbeau"]++
				c.money -= 1
				fmt.Println("Vous avez acheté Plume de Corbeau pour 1 pièce d'or et la compétence a été ajoutée à votre inventaire.")
			}
		default:
			fmt.Println("Article invalide. Veuillez choisir un article valide.")
		}
	} else {
		fmt.Println("Le marchand n'a plus d'articles à vendre pour le moment.")
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
			LivreDeSortBDF++
			return
		}
	}
	fmt.Println("Pour apprendre la compétence Boule de Feu, il faut d'abord l'acheter chez le marchand.")
}

func Capitalize(s *string) string {
	result := ""
	isFirstChar := true
	for _, char := range *s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if isFirstChar {
				if char >= 'a' && char <= 'z' {
					char -= 32
				}
				isFirstChar = false
			} else {
				if char >= 'A' && char <= 'Z' {
					char += 32
				}
			}
			result += string(char)
		} else {
			isFirstChar = true
			result += string(char)
		}
	}
	return result
}

func (c *Character) charCreation() {
	var nickname string
	var class int

	fmt.Println("Quel est votre pseudo ?")
	for {
		fmt.Scanln(&nickname)
		if isAlpha(nickname) {
			break
		} else {
			fmt.Println("Le pseudo doit contenir uniquement des lettres. Réessayez :")
		}
	}
	c.nickname = Capitalize(&nickname)

	fmt.Println("Vous vous appelerez donc", c.nickname)

	fmt.Println("Quelle classe souhaitez-vous être ?")
	fmt.Println("1. Elfe")
	fmt.Println("2. Humain")
	fmt.Println("3. Nain")
	fmt.Print("Choisissez un numéro de classe (1/2/3) : ")
	fmt.Scanln(&class)

	switch class {
	case 1:
		c.classe = "Elfe"
		fmt.Println("Vous êtes donc un elfe !")
		inventory := map[string]int{"Arme": 1, "Armure": 1}
		P1 = Init(c.nickname, c.classe, 1, 80, 40, 100, inventory, []string{"Coup de poing"})
	case 2:
		c.classe = "Humain"
		fmt.Println("Vous êtes donc un Humain !")
		inventory := map[string]int{"Arme": 1, "Armure": 1}
		P1 = Init(c.nickname, c.classe, 1, 100, 50, 100, inventory, []string{"Coup de poing"})
	case 3:
		c.classe = "Nain"
		fmt.Println("Vous êtes donc un Nain !")
		inventory := map[string]int{"Arme": 1, "Armure": 1}
		P1 = Init(c.nickname, c.classe, 1, 120, 60, 100, inventory, []string{"Coup de poing"})
	default:
		fmt.Println("Classe invalide. Vous serez un Humain par défaut.")
		c.classe = "Humain"
		inventory := map[string]int{"Arme": 1, "Armure": 1}
		P1 = Init(c.nickname, c.classe, 1, 100, 50, 100, inventory, []string{"Coup de poing"})
	}
}

func isAlpha(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
}

func (c *Character) LimitInv() bool {
	if len(c.inventory) > 10 {
		fmt.Println("Votre inventaire est plein ! Vous ne pouvez plus ajouter d'objets !")
		return false
	} else {
		return true
	}
}
