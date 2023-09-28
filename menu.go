package red

import (
	"fmt"
	"os"
)

func Menu() {
	var choix string

	for choix != "0" {
		fmt.Println("\n----------------------------------------------")
		fmt.Println("Menu :")
		fmt.Println(" 0. Quitter \n 1. Afficher les informations du personnage \n 2. Accéder au contenu de l'inventaire \n 3. Vérifier si je suis mort \n 4. Marchand\n 5. Prendre Potion\n 6. Apprendre Boule de Feu\n 7. Créer son personnage\n 8. Retirer un objet de l'inventaire\n 9. Forgeron\n 10. Voir mes équipements et en équiper un")
		fmt.Print("\n----------------------------------------------\n")
		fmt.Print("Choisissez une option (0/1/2/3/4/5/6/7/8/9/10) : \n")
		fmt.Print("----------------------------------------------\n")
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
			P1.Dead()
		case "4":
			fmt.Println("Bienvenue chez le marchand !")
			Marchand(&P1)
		case "5":
			fmt.Println("Gloup, Gloup, Gloup, Gloup,...")
			P1.takePotHeal()
		case "6":
			fmt.Println("Apprendre la compétence Boule de Feu")
			P1.spellBook()
		case "7":
			fmt.Println("Création du personnage :")
			P1.CharCreation()
		case "8":
			fmt.Println("Retirer un objet de l'inventaire :")
			P1.displayInventory()
			P1.retirerObjet()
		case "9":
			fmt.Println("Forgeron")
			P1.Forgeron()
		case "10":
			fmt.Println("Equipement")
			menuEquipement(&P1)
		case "11":
			fmt.Println("tpt")
			P1.TPT(&M2)

		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (0/1/2/3/4/5/6/7/8/9/10)")
		}
	}
}

func MenuBase() {
	var choix string

	for choix != "0" {
		fmt.Println("----------------------------------------------")
		fmt.Println("Menu :")
		DefilSpeed("\n0. Quitter \n1. Afficher les informations du personnage \n2. Accéder au contenu de l'inventaire\n3. Marchand\n")
		fmt.Println("----------------------------------------------")
		fmt.Println("Choisissez une option (0/1/2/3) : ")
		fmt.Println("----------------------------------------------")
		fmt.Scanln(&choix)

		switch choix {
		case "0":
			Defil("\nEtes vous sure de vouloir quitter le jeu ?\n1. NON\n0. OUI\n")
			var sure string
			fmt.Scanln(&sure)
			switch sure {
			case "1":
				MenuBase()
			case "0":
				fmt.Println("Merci d'avoir joué, à bientôt !")
				os.Exit(0)
			}
		case "1":
			fmt.Println("Affichage des informations du personnage...")
			P1.displayInfo()
		case "2":
			fmt.Println("Accès au contenu de l'inventaire...")
			P1.displayInventory()
		case "3":
			fmt.Println("Bienvenue chez le marchand !")
			Defil("\nLe Marchand a l'air de vendre des potions de soin ! Nous devrions en prendre pour nos prochains combats !")
			Marchand(&P1)
			if PotionsDeSoinVendues == 1 {
				Defil("La potion de vie a été ajoutée à votre inventaire !\nElle nous sera utile si nous croisons des monstre ")
				Defil("\nUn Loup apparait soudainement devant le Marchand ! \nNous devons le combattre pour empecher qu'il ne fasses des ravages dans le village !")
				return
			}
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (0/1/2/3)")
		}
	}
}

func MenuBDF(c *Character) {
	var choix string

	for choix != "0" {
		fmt.Println("----------------------------------------------")
		fmt.Println("Menu :")
		DefilSpeed("\n0. Quitter \n1. Marchand\n")
		fmt.Println("----------------------------------------------")
		fmt.Println("Choisissez une option (0/1/2/3) : ")
		fmt.Println("----------------------------------------------")
		fmt.Scanln(&choix)

		switch choix {
		case "0":
			Defil("\nEtes vous sure de vouloir quitter le jeu ?\n1. NON\n0. OUI\n")
			var sure string
			fmt.Scanln(&sure)
			switch sure {
			case "1":
				MenuBase()
			case "0":
				fmt.Println("Merci d'avoir joué, à bientôt !")
				os.Exit(0)
			}
		case "1":
			Defil("\n\nAchetez le \"Livre des Sorts : Boule de Feu !\"")
			Marchand(&P1)
			for _, i := range c.skill {
				if i == "Livre des Sorts : Boule de Feu" {
					Defil("\nLe Marchand vous indique une vieille tour dans laquelle vit un grand sorcier qui pourrait vous apporter des renseignements sur ce grimoire...\n")
					return
				}
			}

		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (0/1)")
		}
	}
}
