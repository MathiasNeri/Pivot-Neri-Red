package red

import "fmt"

func (c *Character) Intro() {
	DefilIntro("Bienvenu dans notre jeu ", &P1, " ! \nNous t'attendions avec impatience ! Notre royaume a été attaqué par de terrible monstre qui nuisent au bonheur de nos habitants... \nToi seul pourra vaincre ces envahisseurs !\nTiens ! En voila un !\n")
}

func IMarchand() {
	Defil("\nAprès ce rude combat vous arrivez dans un village...\nVous appercevez au loin un Marchand. Allons vois ce qu'il a à vendre !\n")
	Defil("\nAppuyez sur la touche 3 de votre clavier pour aller au Marchand\n")
}

func AfterLoup() {
	Defil("\n\nLe village vous remercie !\nLe Marchand vous fait savoir qu'un vieux sorcier aurait laissé un grimoire dans la boutique et que celui-ci permettrai d'apprendre à faire des Boule de Feu lors des combats !\n")
}

func AllerTour() {
	Defil("\nVous vous dirigez vers la tour que le Marchand vous a indiqué quand soudain un Troll apparait devant vous !\n")
}

func Tour() {
	Defil("\nVous arrivez à la tour ou vous rencontrez le sorcier\nCelui-ci vous propose de vous apprendre à maitriser les boules de Feu\nIl vous explique que une fois maitriser cette attaque est devastatrice mais très difficile a utiliser en combat !\n")
}

func ApprendreBDF() {
	Defil("\nPour apprendre la Boule de feu, appuyez sur la touche 6 de votre clavier !")
	var choixbdf string
	fmt.Scanln(&choixbdf)
	switch choixbdf {
	case "6":
		P1.spellBook()
		Defil("\nVous venez d'apprendre la compétence Boule de Feu !")
		return
	default:
		Defil("\nTouche incorrecte.")
		ApprendreBDF()
	}
}

func AfterBDF() {
	Defil("\nAprès avoir apris la Boule de Feu vous décidez de retourner au village.\nVous vous dirigez vers le forgeron qui vous propose de fabriquer des équipements pour vous aider à combattre les monstres !")
	Defil("\nIl semblerai que vous ayez les ressources nécessaires pour fabriquer la Tunique de l'aventurier !\n\nAppuyez sur la touche 2 de votre clavier pour fabriquer la Tunique de l'aventurier.")
}

func BossFinal() {
	Defil("Vous avez a peine le temps d'équiper votre tunique qu'un immmense dragon en origami survole le village!!!\nKheir-Eddine le brave, le maire du village est désespéré de la situation...\nNous devons abbatre ce terrible dragon !\n")
}

func LeHero() {
	Defil("\nLe dragon est mort !\n\nVous etes maintenant un héros !\n\nToute nos félicitation vous avez terminé le mode histoire avec succés !\nVous allez maintenant pouvoir explorer notre monde ouvert (le terminal...) !\nMais prenez garde ! De nombreux monstres rodent toujours...\n")
}
