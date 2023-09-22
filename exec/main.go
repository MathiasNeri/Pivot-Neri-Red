package main

import p "RED"

func main() {
	p.P1.CharCreation()
	p.M1.InitGobelin()
	p.P1.TPT(&p.M1)
	p.Menu()
}
