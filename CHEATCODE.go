package red

func (c *Character) GiveMatTunique() {
	c.inventory["Fourrure de Loup"]++
	c.inventory["Peau de Troll"]++
}

func Slay(m *Monstre) {
	m.curpv -= 100000000000
}
