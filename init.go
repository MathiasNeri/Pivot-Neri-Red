package red

import "math/rand"

var P1 Character // moi
var M1 Monstre   // gobelin d'entrainement
var M2 Monstre   // Monstre d'essai
var M3 Monstre   // Loup
var M4 Monstre   // Troll
var M5 Monstre   // Dragon
var CoupdePoing Skill
var BouledeFeu Skill

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, money int, mana int, inventory map[string]int, skill []string) Character {
	Character := Character{
		nickname:      nickname,
		classe:        classe,
		lvl:           lvl,
		hp_max:        hp_max,
		current_hp:    current_hp,
		money:         money,
		mana:          mana,
		inventory:     inventory,
		skill:         skill,
		inventoryList: []string{}, // Initialise la liste vide
	}
	return Character
}

func (m *Monstre) InitGobelin() {
	m.name = "Gobelin d'entrainement"
	m.pvmax = 40
	m.curpv = 40
	m.damagept = 3
}

func (m *Monstre) InitLoup() {
	m.name = "Loup très méchant"
	m.pvmax = 60
	m.curpv = 40 + rand.Intn(10)
	m.damagept = 4
}

func (m *Monstre) InitTroll() {
	m.name = "Troll tout pas beau"
	m.pvmax = 80
	m.curpv = 50 + rand.Intn(29)
	m.damagept = 2
}

func (m *Monstre) InitDragon() {
	m.name = "Dragon en Origami"
	m.pvmax = 120
	m.curpv = 120
	m.damagept = 10
}

func (char *equipement) InitEquipement(Tete string, Torse string, Bottes string) { // Équipement du personnage
	char.Tete = Tete
	char.Torse = Torse
	char.Bottes = Bottes
}

func (s *Skill) InitSkill(nom string, dmg int) {
	s.nom = nom
	s.dmg = dmg
}
