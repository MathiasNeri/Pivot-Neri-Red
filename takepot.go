package red

import (
	"fmt"
	"time"
)

func (c *Character) takePotHeal() {
	fmt.Println("Quelle potion de soin souhaitez-vous prendre ?")
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
}

func (c *Character) takePotPoison(adversaire *Monstre) {
	fmt.Println("Quelle potion de poison souhaitez-vous prendre ?")
	if c.inventory["Potion de Poison"] > 0 {
		c.inventory["Potion de Poison"]--
		poisonDuration := 3 * time.Second
		damageInterval := 1 * time.Second
		damagePerTick := 10

		fmt.Println("Vous avez utilisé une Potion de Poison sur l'adversaire !")

		ticker := time.NewTicker(damageInterval)
		defer ticker.Stop()

		damageTimer := time.NewTimer(poisonDuration)
		defer damageTimer.Stop()

		for {
			select {
			case <-ticker.C:
				adversaire.curpv -= damagePerTick
				if adversaire.curpv < 0 {
					adversaire.curpv = 0
				}
				fmt.Printf("Points de vie de l'adversaire : %d / %d\n", adversaire.curpv, adversaire.pvmax)
			case <-damageTimer.C:
				fmt.Println("L'effet du poison sur l'adversaire est terminé.")
				return
			}
		}
	} else {
		fmt.Println("Vous n'avez pas de Potion de Poison dans l'inventaire.")
	}
}

func (c *Character) takePotMana() {
	if c.inventory["Potion de Mana"] > 0 {
		c.inventory["Potion de mana"]--
		pointsDeMana := 60
		c.current_hp += pointsDeMana

		if c.current_hp > c.hp_max {
			c.current_hp = c.hp_max
		}

		fmt.Printf("Vous avez utilisé une Potion de Mana et avez récupéré %d points de vie.\n", pointsDeMana)
		fmt.Printf("Points de vie actuels : %d / %d\n", c.current_hp, c.hp_max)
	} else {
		fmt.Println("Vous n'avez pas de Potion de Mana dans l'inventaire.")
	}
}
