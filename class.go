package red

type equipement struct { // Équipement du joueur
	Tete   string
	Torse  string
	Bottes string
}

type Character struct {
	nickname      string
	classe        string
	lvl           int
	hp_max        int
	current_hp    int
	money         int
	mana          int
	inventory     map[string]int
	skill         []string
	inventoryList []string
	Equipement    equipement
}

type Monstre struct {
	name     string
	pvmax    int
	curpv    int
	damagept int
}

type Skill struct {
	nom         string // Nom de la compétence
	dmg         int    // Dommages infligés
	mana        int
	Description string //Courte description de l'attaque
}
