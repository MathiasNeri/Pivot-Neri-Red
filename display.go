package red

import "fmt"

func (c *Character) displayInfo() {
	fmt.Printf("\nNickname: %s\nClass: %s\nLevel: %d\nHp_Max : %d\nCurrent_Hp : %d\nMoney : %d\nMana : %d\nSkill : %s\n",
		c.nickname, c.classe, c.lvl,
		c.hp_max, c.current_hp, c.money, c.mana, c.skill)
}

func (c *Character) displayInventory() {
	fmt.Println("Inventaire actuel :")
	c.inventoryList = nil
	i := 1
	for item, count := range c.inventory {
		fmt.Printf("%d. %s : %d\n", i, item, count)
		c.inventoryList = append(c.inventoryList, item)
		i++
	}
}
