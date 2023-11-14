package mobs

type Monster struct {
	Name               string // name of the mob
	HP                 int    // hitpoints
	Damage             int    // damage per attack
	ExperienceRewarded int    // exp on kill
}

func NewMonster(name string, hp int, damage int, exprew int) Monster {
	return Monster{
		Name:               name,
		HP:                 hp,
		Damage:             damage,
		ExperienceRewarded: exprew,
	}
}

func NewRat() Monster {
	return NewMonster("Rat", 10, 1, 3)
}

func NewGoblin() Monster {
	return NewMonster("Goblin", 20, 3, 40)
}

func NewTortelli() Monster {
	return NewMonster("Tortelli", 30, 5, 13)
}
