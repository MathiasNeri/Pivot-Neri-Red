package main

import p "RED"

func main() {
	p.P1.CharCreation()
	p.P1.Intro()
	p.M1.InitGobelin()
	p.P1.TPTtuto(&p.M1)

	p.IMarchand()
	p.MenuBase()
	p.M3.InitLoup()
	p.P1.TPTLoup(&p.M3)

	p.AfterLoup()
	p.MenuBDF(&p.P1)
	p.AllerTour()
	p.M4.InitTroll()
	p.P1.TPTTroll(&p.M4)
	p.Tour()
	p.ApprendreBDF()
	p.AfterBDF()
	p.P1.ForgeronTuto()
	p.BossFinal()
	p.M5.InitDragon()
	p.P1.TPTDragon(&p.M5)
	p.LeHero()
	p.Menu()
}

/*
p.P1.TPT(&p.M1)
	p.M2.InitGobelin()
	p.Menu()
*/
