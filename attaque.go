package red

// AttaqueBasique crée une attaque basique commune à toutes les classes.
func AttaqueBasique() Skill {
	return Skill{
		nom:         "Attaque basique",
		dmg:         10, // Mettez les dommages appropriés ici
		Description: "Attaque de base qui inflige des dommages standard.",
	}
}

// AttaqueNain crée une attaque spécifique à la classe nain.
func AttaqueNain() Skill {
	return Skill{
		nom:         "Frappe Sismique",
		dmg:         15, // Dommages spécifiques aux nains
		Description: "Les nains utilisent leur force et leur robustesse pour créer une onde de choc dévastatrice. Ils frappent le sol avec une grande puissance, provoquant des secousses sismiques qui endommagent leurs ennemis.",
	}
}

// AttaqueElfe crée une attaque spécifique à la classe elfe.
func AttaqueElfe() Skill {
	return Skill{
		nom:         "Tir de Précision",
		dmg:         15, // Dommages spécifiques aux elfes
		Description: " Les elfes excellent dans l'archerie. Ils décochent une flèche avec une précision mortelle, visant les points faibles de l'ennemi..",
	}
}

// AttaqueHumain crée une attaque spécifique à la classe humain.
func AttaqueHumain() Skill {
	return Skill{
		nom:         "Stratégie Tactique",
		dmg:         15, // Dommages spécifiques aux humains
		Description: "Les humains sont connus pour leur intelligence. Ils utilisent leur connaissance stratégique pour anticiper les mouvements de l'ennemi et attaquer avec précision.",
	}
}
