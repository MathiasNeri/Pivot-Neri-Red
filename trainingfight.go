package red

import "fmt"

func TrainingFight() {
	for i := 1; i <= 20; i++ {
		fmt.Println("Tour", i)
		if i%2 != 0 {
			fmt.Println("L'ennemi attaque !")
		} else {
			fmt.Println("On attaque !")
		}
	}
}
