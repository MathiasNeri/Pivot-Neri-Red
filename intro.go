package red

func (c *Character) Intro() {
	DefilIntro("Bienvenu dans notre jeu ", &P1, " ! \nNous t'attendions avec impatience ! Notre royaume a été attaqué par de terrible monstre qui nuisent au bonheur de nos habitants... \nToi seul pourra vaincre ces envahisseurs !\nTiens ! En voila un !")
}

func IMarchand() {
	Defil("\nAprès ce rude combat vous arrivez dans un village...\nVous appercevez au loin un Marchand. Allons vois ce qu'il a à vendre !\n")
	Defil("\nAppuyez sur la touche 3 de votre clavier pour aller au Marchand\n")
}

func AfterLoup() {
	Defil("\nVous avez récupéré une peau de loup et elle a été ajoutée à votre inventaire \nLe village vous remercie !\nLe Marchand vous fait savoir qu'un vieux sorcier aurait laissé un vieux grimoire dans la boutique qui permettrai d'apprendre a faire des Boule de Feu !\n")
}

func AllerTour() {
	Defil("\nVous vous dirigez vers la tour que le Marchand vous a indiqué quand soudain un Troll apparait devant vous !")
}
