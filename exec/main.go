package main

import p "RED"

func main() {
	p.P1.CharCreation()
	//p.P1.Intro()
	//p.M1.InitGobelin()
	//p.P1.TPTtuto(&p.M1)
<<<<<<< HEAD
	//p.IMarchand()
	//p.MenuBase()
	//p.M3.InitLoup()
	//p.P1.TPTLoup(&p.M3)
=======
	p.IMarchand()
	p.MenuBase()
	p.M3.InitLoup()
	p.P1.TPTLoup(&p.M3)
>>>>>>> 55b6fee0f469466a7421811ded1534c72904826d
	p.AfterLoup()
	p.MenuBDF(&p.P1)
	p.AllerTour()
	p.M4.InitTroll()
	p.P1.TPTTroll(&p.M4)

}

/*
p.P1.TPT(&p.M1)
	p.M2.InitGobelin()
	p.Menu()
*/
