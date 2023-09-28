package red

import "fmt"

var PotionsDeSoinVendues = 0
var PotionsDePoisonVendues = 0
var LivreDeSortBDF = 0
var FourrureVendues = 0
var PeauTrollVendue = 0
var CuirSanglierVendue = 0
var PlumeCorbeauVendue = 0
var PotionMana = 0

func Marchand(c *Character) {
	if PotionsDeSoinVendues < 1 || PotionsDePoisonVendues < 1 || LivreDeSortBDF < 1 || PotionMana < 1 || FourrureVendues < 1 || PeauTrollVendue < 1 || CuirSanglierVendue < 1 || PlumeCorbeauVendue < 1 {
		fmt.Println("Articles disponibles chez le marchand :")
		if PotionsDeSoinVendues < 1 {
			fmt.Println("1. Potion de Soin : 3 pièces d'or")
		}
		if PotionsDePoisonVendues < 1 {
			fmt.Println("2. Potion de Poison : 6 pièces d'or")
		}
		if LivreDeSortBDF < 1 {
			fmt.Println("3. Livre de Sort : Boule de Feu : 25 pièces d'or")
		}
		if PotionMana < 1 {
			fmt.Println("4. Potion de Mana : 5 pièces d'or")
		}
		if FourrureVendues < 1 {
			fmt.Println("5. Fourrure de Loup : 4 pièces d'or")
		}
		if PeauTrollVendue < 1 {
			fmt.Println("6. Peau de Troll : 7 pièces d'or")
		}
		if CuirSanglierVendue < 1 {
			fmt.Println("7. Cuir de Sanglier : 3 pièces d'or")
		}
		if PlumeCorbeauVendue < 1 {
			fmt.Println("8. Plume de Corbeau : 1 pièce d'or")
		}

		var choix string
		fmt.Println("Choisissez un article :")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionsDeSoinVendues < 1 {
				c.inventory["Potion de Soin"]++
				stock++
				PotionsDeSoinVendues++
				c.money -= 3
				fmt.Println("Vous avez acheté une Potion de Soin pour 3 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Soin à vendre.")
			}
		case "2":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionsDePoisonVendues < 1 {
				c.inventory["Potion de Poison"]++
				stock++
				PotionsDePoisonVendues++
				c.money -= 6
				fmt.Println("Vous avez acheté une Potion de Poison pour 6 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Poison à vendre.")
			}
		case "3":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if LivreDeSortBDF < 1 {
				c.inventory["Livre de Sort : Boule de Feu"]++
				stock++
				LivreDeSortBDF++
				c.money -= 25
				fmt.Println("Vous avez acheté Livre de Sort : Boule de Feu pour 25 pièces d'or et la compétence a été ajoutée à votre inventaire.")
			} else if LivreDeSortBDF >= 1 {
				fmt.Println("Le marchand n'a plus de Potion de Livre de Sort :Boule de Feu à vendre.")
			}
		case "4":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionMana < 1 {
				c.inventory["Potion de Mana"]++
				stock++
				PotionMana++
				c.money -= 5
				fmt.Println("Vous avez acheté une Potion de Mana pour 5 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Mana à vendre.")
			}
		case "5":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if FourrureVendues < 1 {
				c.inventory["Fourrure de Loup"]++
				stock++
				c.money -= 4
				fmt.Println("Vous avez acheté Fourrure de Loup pour 4 pièces d'or et l'item a été ajoutée à votre inventaire.")
			}
		case "6":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PeauTrollVendue < 1 {
				c.inventory["Peau de Troll"]++
				stock++
				c.money -= 7
				fmt.Println("Vous avez acheté Peau de Troll pour 7 pièces d'or et l'item' a été ajoutée à votre inventaire.")
			}
		case "7":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if CuirSanglierVendue < 1 {
				c.inventory["Cuir de Sanglier"]++
				stock++
				c.money -= 3
				fmt.Println("Vous avez acheté Cuir de Sanglier pour 3 pièces d'or et l'item a été ajoutée à votre inventaire.")
			}
		case "8":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PlumeCorbeauVendue < 1 {
				c.inventory["Plume de Corbeau"]++
				stock++
				c.money -= 1
				fmt.Println("Vous avez acheté Plume de Corbeau pour 1 pièce d'or et l'item a été ajoutée à votre inventaire.")
			}
		default:
			fmt.Println("Article invalide. Veuillez choisir un article valide.")
		}
	} else {
		fmt.Println("Le marchand n'a plus d'articles à vendre pour le moment.")
	}
}

func Marchand1(c *Character) {
	if PotionsDeSoinVendues < 1 {
		fmt.Println("Articles disponibles chez le marchand :")
		if PotionsDeSoinVendues < 1 {
			fmt.Println("1. Potion de Soin : gratuit")
		}

		var choix string
		fmt.Println("Choisissez un article :")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionsDeSoinVendues < 1 {
				c.inventory["Potion de Soin"]++
				stock++
				PotionsDeSoinVendues++
				fmt.Println("Vous avez récupérer une Potion de Soin gratuitement et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Soin à vendre.")
				fmt.Println("\nRevenez plus tard !")

			}

		default:
			fmt.Println("Article invalide. Veuillez choisir un article valide.")
		}
	}
}

func Marchand2(c *Character) {
	PotionsDePoisonVendues = 0
	PotionMana = 0
	if PotionsDeSoinVendues < 1 || PotionsDePoisonVendues < 1 || LivreDeSortBDF < 1 || PotionMana < 1 {
		fmt.Println("Articles disponibles chez le marchand :")
		if PotionsDeSoinVendues < 1 {
			fmt.Println("1. Potion de Soin : 3 pièces d'or")
		}
		if PotionsDePoisonVendues < 1 {
			fmt.Println("2. Potion de Poison : 6 pièces d'or")
		}
		if LivreDeSortBDF < 1 {
			fmt.Println("3. Livre de Sort : Boule de Feu : 25 pièces d'or")
		}
		if PotionMana < 1 {
			fmt.Println("4. Potion de Mana : 5 pièces d'or")
		}

		var choix string
		fmt.Println("Choisissez un article :")
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionsDeSoinVendues < 1 {
				c.inventory["Potion de Soin"]++
				stock++
				PotionsDeSoinVendues++
				c.money -= 3
				fmt.Println("Vous avez acheté une Potion de Soin pour 3 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Soin à vendre.")
			}
		case "2":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionsDePoisonVendues < 1 {
				c.inventory["Potion de Poison"]++
				stock++
				PotionsDePoisonVendues++
				c.money -= 6
				fmt.Println("Vous avez acheté une Potion de Poison pour 6 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Poison à vendre.")
			}
		case "3":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if LivreDeSortBDF < 1 {
				c.inventory["Livre de Sort : Boule de Feu"]++
				stock++
				LivreDeSortBDF++
				c.money -= 25
				fmt.Println("Vous avez acheté Livre de Sort : Boule de Feu pour 25 pièces d'or et la compétence a été ajoutée à votre inventaire.")
			} else if LivreDeSortBDF >= 1 {
				fmt.Println("Le marchand n'a plus de Potion de Livre de Sort :Boule de Feu à vendre.")
			}
		case "4":
			if !c.LimitInv() {
				fmt.Println("Votre inventaire est plein !")
				return
			}
			if PotionMana < 1 {
				c.inventory["Potion de Mana"]++
				stock++
				PotionMana++
				c.money -= 5
				fmt.Println("Vous avez acheté une Potion de Mana pour 5 pièces d'or et elle a été ajoutée à votre inventaire.")
			} else {
				fmt.Println("Le marchand n'a plus de Potion de Mana à vendre.")
			}
		default:
			fmt.Println("Article invalide. Veuillez choisir un article valide.")
		}
	} else {
		fmt.Println("Le marchand n'a plus d'articles à vendre pour le moment.")
	}
}
