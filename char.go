package red

import (
	"fmt"
	"time"
)

func (c *Character) CharCreation() {
	var nickname string
	var class int
	pseudo := "\nQuel est votre pseudo ? \n "
	for _, letter := range pseudo {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
	for {
		fmt.Println("")
		fmt.Scanln(&nickname)
		if isAlpha(nickname) {
			break
		} else {
			fmt.Println("Le pseudo doit contenir uniquement des lettres. Réessayez :")
		}
	}
	c.nickname = Capitalize(&nickname)

	lenom := ("\nVous vous appelerez donc ")
	for _, letter := range lenom {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Print(c.nickname)
	fmt.Println("")

	choixc := "\nA quelle classe souhaitez-vous appartenir ?"
	for _, letter := range choixc {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}

	fmt.Println("\n\n1. Elfes : Majestueux, mais avec un vitalité faible...\n\n2. Humains : Ingénieux mais parfois très bêtes...\n\n3. Nains : Pas très grand mais très robuste !\n\nChoisissez un numéro de classe (1/2/3) : ")
	fmt.Scanln(&class)

	switch class {
	case 1:
		c.classe = "Elfes"
		fmt.Println("Vous êtes donc un Elfe !")
		inventory := map[string]int{}
		P1 = Init(c.nickname, c.classe, 1, 80, 40, 100, inventory, []string{"Coup de poing"})
	case 2:
		c.classe = "Humains"
		fmt.Println("Vous êtes donc un Humain !")
		inventory := map[string]int{}
		P1 = Init(c.nickname, c.classe, 1, 100, 50, 100, inventory, []string{"Coup de poing"})
	case 3:
		c.classe = "Nains"
		fmt.Println("Vous êtes donc un Nain !")
		inventory := map[string]int{}
		P1 = Init(c.nickname, c.classe, 1, 120, 60, 100, inventory, []string{"Coup de poing"})
	default:
		fmt.Println("Classe invalide. Vous serez un Humain par défaut.")
		c.classe = "Humains"
		inventory := map[string]int{}
		P1 = Init(c.nickname, c.classe, 1, 100, 50, 100, inventory, []string{"Coup de poing"})
	}

	Ctipar := "\n\nNous allons pouvoir commencé !\n\n"
	for _, letter := range Ctipar {
		fmt.Print(string(letter))
		time.Sleep(5 * time.Millisecond)
	}
}
