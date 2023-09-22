package red

var P1 Character
var M1 Monstre
var CoupdePoing Skill
var BouledeFeu Skill

func Init(nickname string, classe string, lvl int, hp_max int, current_hp int, money int, inventory map[string]int, skill []string) Character {
	Character := Character{
		nickname:      nickname,
		classe:        classe,
		lvl:           lvl,
		hp_max:        hp_max,
		current_hp:    current_hp,
		money:         money,
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
	m.damagept = 5
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
