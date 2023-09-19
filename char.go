package red

import "fmt"

func (c *Character) CharCreation() {
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
