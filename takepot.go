package red

import (
	"fmt"
	"time"
)

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
