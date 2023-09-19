package red

import (
	"fmt"
	"time"
)

var P1 Character
var stock int = 2 // a modifier quand on pourra craft les equipements

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
	if stock == 10 {
		return false
	} else {
		return true
	}
}

var chapeauVendu = 0
var tuniqueVendu = 0
var botteVendu = 0

func (c *Character) Forgeron() {

	fmt.Println("Choisissez un item à fabriquer:")
	if chapeauVendu < 1 {
		fmt.Println("1. le Chapeau de l'aventurier")
	}
	if tuniqueVendu < 1 {
		fmt.Println("2 Tunique de l'aventurier")
	}
	if botteVendu < 1 {
		fmt.Println("3 pour Bottes de l'aventurier")
	}
	var choix string
	fmt.Scan(&choix)
	switch choix {
	case "1":
		fmt.Println("Pour fabriqur cet item il vous faut :")
		fmt.Println("X1 Plume de Corbeau")
		fmt.Println("X1 Cuir de Sanglier")
		fmt.Println("")
		fmt.Println("Voulez-vous créer cet item ?")
		fmt.Println("1. OUI")
		fmt.Println("2. NON")
		fmt.Println("")
		var choixc string
		fmt.Scan(&choixc)
		switch choixc {
		case "1":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			for i := range c.inventory {
				if i == "Plume de Corbeau" {
					for j := range c.inventory {
						if j == "Cuir de Sanglier" {
							c.inventory["Chapeau de l'aventurier"]++
							stock++
							chapeauVendu++
						}
					}
				}
			}
		}
	case "2":
		fmt.Println("Pour fabriqur cet item il vous faut :")
		fmt.Println("X2 Fourrure de Loup")
		fmt.Println("X1 Peau de Troll")
		fmt.Println("")
		fmt.Println("Voulez-vous créer cet item ?")
		fmt.Println("1. OUI")
		fmt.Println("2. NON")
		fmt.Println("")
		var choixt string
		fmt.Scan(&choixt)
		switch choixt {
		case "1":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			for i := range c.inventory {
				if i == "Fourrure de Loup" {
					for j := range c.inventory {
						if j == "Peau de Troll" {
							c.inventory["Tunique de l'aventurier"]++
							stock++
							tuniqueVendu++
						}
					}
				}
			}
		}
	}
}
