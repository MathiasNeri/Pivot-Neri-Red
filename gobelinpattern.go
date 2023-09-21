package red

import "fmt"

var nbattaqueg int

func (m *Monstre) GobelinPattern() {
	if nbattaqueg == 3 {
		P1.PlayerDamage(m.damagept * 2)
		nbattaqueg = 0
	} else {
		P1.PlayerDamage(m.damagept)
		nbattaqueg += 1
	}
}
func (c *Character) PlayerDamage(dmg int) {
	c.current_hp -= dmg
	fmt.Println("Vous avez perdu", dmg, "HP")
}
