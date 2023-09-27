package red

func (c *Character) Intro() {
	DefilIntro("Bienvenu dans notre jeu ", &P1, " ! \nNous t'attendions avec impatience ! Notre royaume a été attaqué par de terrible monstre qui nuisent au bonheur de nos habitants... \nToi seul pourra vaincre ces envahisseurs !\nTiens ! En voila un !")
}

func IMarchand() {
	Defil("\nAprès ce rude combat vous arrivez dans un village...\nVous appercevez au loin un Marchand. Allons vois ce qu'il a à vendre !\n")
	Defil("\nAppuyez sur la touche 3 de votre clavier pour aller au Marchand\n")
}
