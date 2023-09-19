package red

import (
	"fmt"
	"os"
)

func Menu() {
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
		fmt.Println("9. Forgeron")
		fmt.Println("10. Voir mes équipements et en équiper un")
		fmt.Println("----------------------------------------------")
		fmt.Print("Choisissez une option (0/1/2/3/4/5/6/7/8/9/10) : ")

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
			Marchand(&P1)
		case "5":
			fmt.Println("Gloup, Gloup, Gloup, Gloup,...")
			P1.takePot()
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
			fmt.Println("Gérer les équipements :")
			GestionEquipements(&P1)
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide (0/1/2/3/4/5/6/7/8/9/10)")
		}
	}
}
