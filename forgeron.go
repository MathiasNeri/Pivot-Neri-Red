package red

import (
	"fmt"
)

// pense a les remettre a 0 quand tu les perd clampin

var Chapeauvendu int
var Tuniquevendu int
var Bottevendu int

func (c *Character) Forgeron() {
	fmt.Println("Quel item ouhaitez vous fabriquer ?")
	fmt.Println("0. Revenir au menu")
	if Chapeauvendu < 1 {
		fmt.Println("1. Chapeau de l'aventurier (X1 Plume de Corbeau + X1 Cuir de Sanglier) et 5 pièces d'or")
	}
	if Tuniquevendu < 1 {
		fmt.Println("2. Tunique de l'aventurier (X1 Fourrure de Loup + X1 Peau de Troll) et 5 pièces d'or")
	}
	if Bottevendu < 1 {
		fmt.Println("3. Botte de l'aventurier (X1 Fourrure de Loup + X1 Cuir de Sanglier) et 5 pièces d'or")
	}

	var choix string
	fmt.Scan(&choix)
	switch choix {
	case "0":
		Menu()
	case "1":
		if !c.LimitInv() {
			fmt.Println("Votre inventaire est plein !")
			return
		}
		var stockc int
		for i := range c.inventory {
			if i == "Plume de Corbeau" {
				stockc += 1
			}
			if i == "Cuir de Sanglier" {
				stockc += 1
			}
		}
		if stockc == 2 {
			c.money -= 5
			fmt.Println("Vous venez de craft le Chapeau de l'aventurier !")
			c.inventory["Chapeau de l'aventurier"]++
			stock++
			Chapeauvendu++
			delete(c.inventory, "PLume de Corbeau")
			stock--
			delete(c.inventory, "Cuir de Sanglier")
			stock--
			return
		} else {
			fmt.Println("Vous n'avez pas les objets requis pour carft l'item")

		}
	case "2":
		if !c.LimitInv() {
			fmt.Println("Votre inventaire est plein !")
			return
		}
		var stockt int
		for i := range c.inventory {
			if i == "Fourrure de Loup" {
				stockt += 1
			}
			if i == "Peau de Troll" {
				stockt += 1
			}
		}
		if stockt == 2 {
			c.money -= 5
			fmt.Println("Vous venez de fabriquer la Tunique de l'aventurier !")
			c.inventory["Tunique de l'aventurier"]++
			stock++
			Tuniquevendu++
			delete(c.inventory, "Fourrure de loup")
			stock--
			delete(c.inventory, "Peau de Troll")
			stock--
			stockt = 0

		} else {
			fmt.Println("Vous n'avez pas les objets requis pour carft l'item")
		}
	case "3":
		if !c.LimitInv() {
			fmt.Println("Votre inventaire est plein !")
			return
		}
		var stockb int
		for i := range c.inventory {
			if i == "Fourrure de Loup" {
				stockb += 1
			}
			if i == "Cuir de Sanglier" {
				stockb += 1
			}
		}
		if stockb == 2 {
			c.money -= 5
			fmt.Println("Vous venez de craft les bottes de l'aventurier !")
			c.inventory["Bottes de l'aventurier"]++
			stock++
			Bottevendu++
			delete(c.inventory, "Fourrure de Loup")
			stock--
			delete(c.inventory, "Cuir de Sanglier")
			stock--
			return
		} else {
			fmt.Println("Vous n'avez pas les objets requis pour carft l'item")

		}
	default:
		fmt.Println("Article invalide. Veuillez choisir un article valide.")

	}
}
func InitEquipements() Equipement {
	return Equipement{
		Helmet: Equipements{
			Nom:              "Chapeau de l'aventurier",
			PointsDeVieBonus: 10,
			Equipe:           false,
		},
		Chestplate: Equipements{
			Nom:              "Tunique de l'aventurier",
			PointsDeVieBonus: 25,
			Equipe:           false,
		},
		Boots: Equipements{
			Nom:              "Bottes de l'aventurier",
			PointsDeVieBonus: 15,
			Equipe:           false,
		},
	}
}

func main() {
	P1 := Character{
		nickname:      "Joueur 1",
		classe:        "Guerrier",
		lvl:           1,
		hp_max:        100,
		current_hp:    100,
		money:         50,
		inventory:     make(map[string]int),
		skill:         []string{"Coup d'épée"},
		inventoryList: []string{},
		Equipements:   InitEquipements(),
	}

	AfficherEquipements(&P1)

	EquipementEquipement(&P1)

	AfficherEquipements(&P1)
}

func AfficherEquipements(c *Character) {
	fmt.Println("Équipements actuels :")
	if c.Equipements.Helmet.Equipe {
		fmt.Println("Chapeau de l'aventurier")
	}
	if c.Equipements.Chestplate.Equipe {
		fmt.Println("Tunique de l'aventurier")
	}
	if c.Equipements.Boots.Equipe {
		fmt.Println("Bottes de l'aventurier")
	}
	// Vous pouvez ajouter d'autres équipements ici en fonction de votre structure Equipment
}

// Fonction pour équiper un équipement
func EquipementEquipement(c *Character) {
	fmt.Println("Équipements disponibles :")
	fmt.Println("1. Chapeau de l'aventurier")
	fmt.Println("2. Tunique de l'aventurier")
	fmt.Println("3. Bottes de l'aventurier")
	fmt.Print("Choisissez un équipement à équiper (1/2/3) : ")

	var choix string
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		if c.inventory["Chapeau de l'aventurier"] > 0 {
			c.Equipements.Helmet.Equipe = true
			c.inventory["Chapeau de l'aventurier"]--
			c.UpdateMaxHP() // Mettez à jour les points de vie maximum
			fmt.Println("Vous avez équipé le Chapeau de l'aventurier.")
		} else {
			fmt.Println("Vous n'avez pas de Chapeau de l'aventurier dans l'inventaire.")
		}
	case "2":
		if c.inventory["Tunique de l'aventurier"] > 0 {
			c.Equipements.Helmet.Equipe = true
			c.inventory["Tunique de l'aventurier"]--
			c.UpdateMaxHP() // Mettez à jour les points de vie maximum
			fmt.Println("Vous avez équipé la Tunique de l'aventurier.")
		} else {
			fmt.Println("Vous n'avez pas de Tunique de l'aventurier dans l'inventaire.")
		}
	case "3":
		if c.inventory["Bottes de l'aventurier"] > 0 {
			c.Equipements.Boots.Equipe = true
			c.inventory["Bottes de l'aventurier"]--
			c.UpdateMaxHP() // Mettez à jour les points de vie maximum
			fmt.Println("Vous avez équipé les Bottes de l'aventurier.")
		} else {
			fmt.Println("Vous n'avez pas de Bottes de l'aventurier dans l'inventaire.")
		}
	default:
		fmt.Println("Choix invalide. Veuillez choisir un équipement valide (1/2/3).")
	}
}
func DesequiperEquipement(c *Character) {
	fmt.Println("Équipements actuels :")
	fmt.Println("1. Chapeau de l'aventurier")
	fmt.Println("2. Tunique de l'aventurier")
	fmt.Println("3. Bottes de l'aventurier")
	fmt.Print("Choisissez un équipement à déséquiper (1/2/3) : ")

	var choix string
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		if c.Equipements.Helmet.Equipe {
			c.Equipements.Helmet.Equipe = false
			c.inventory["Chapeau de l'aventurier"]++
			c.UpdateMaxHP() // Mettez à jour les points de vie maximum
			fmt.Println("Vous avez déséquipé le Chapeau de l'aventurier.")
		} else {
			fmt.Println("Le Chapeau de l'aventurier n'est pas équipé.")
		}
	case "2":
		if c.Equipements.Chestplate.Equipe {
			c.Equipements.Chestplate.Equipe = false
			c.inventory["Tunique de l'aventurier"]++
			c.UpdateMaxHP() // Mettez à jour les points de vie maximum
			fmt.Println("Vous avez déséquipé la Tunique de l'aventurier.")
		} else {
			fmt.Println("La Tunique de l'aventurier n'est pas équipée.")
		}
	case "3":
		if c.Equipements.Boots.Equipe {
			c.Equipements.Boots.Equipe = false
			c.inventory["Bottes de l'aventurier"]++
			c.UpdateMaxHP() // Mettez à jour les points de vie maximum
			fmt.Println("Vous avez déséquipé les Bottes de l'aventurier.")
		} else {
			fmt.Println("Les Bottes de l'aventurier ne sont pas équipées.")
		}
	default:
		fmt.Println("Choix invalide. Veuillez choisir un équipement valide (1/2/3).")
	}
}

func (c *Character) UpdateMaxHP() {
	c.hp_max = 100

	if c.Equipements.Helmet.Equipe {
		c.hp_max += c.Equipements.Helmet.PointsDeVieBonus
	}
	if c.Equipements.Chestplate.Equipe {
		c.hp_max += c.Equipements.Chestplate.PointsDeVieBonus
	}
	if c.Equipements.Boots.Equipe {
		c.hp_max += c.Equipements.Boots.PointsDeVieBonus
	}
}

func GestionEquipements(c *Character) {
	for {
		fmt.Println("Vos équipements actuels :")
		AfficherEquipements(c)

		fmt.Println("Équipements disponibles :")
		fmt.Println("1. Chapeau de l'aventurier")
		fmt.Println("2. Tunique de l'aventurier")
		fmt.Println("3. Bottes de l'aventurier")
		fmt.Println("0. Revenir au menu principal")
		fmt.Print("Choisissez un équipement à équiper (1/2/3) ou 0 pour revenir au menu : ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "0":
			return // Revenir au menu principal
		case "1":
			EquipementEquipement(c)
		// Ajoutez les cas pour les autres équipements si nécessaire
		default:
			fmt.Println("Choix invalide. Veuillez choisir un équipement valide (1/2/3/0).")
		}
	}
}
