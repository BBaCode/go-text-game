package player

import (
	"fmt"

	"github.com/bbacode/go-text-game/mobs"
)

var levelMap = map[int]int{
	2: 100,
	3: 250,
	4: 500,
}

type Player struct {
	Name         string // this is the name
	HP           int    // this is value
	Damage       int    // how much damage we're dealing
	Experience   int    // current expTotal
	CurrentLevel int    // current player level
	NextLevel    int    // next player level
	GameLevel    string // game level player is currently on
}

func NewPlayer(name string) Player {
	return Player{
		Name:         name,
		HP:           100,
		Damage:       8,
		Experience:   0,
		CurrentLevel: 1,
		NextLevel:    2,
		GameLevel:    "Level1",
	}
}

func AddExperience(p Player, m mobs.Monster) Player {
	fmt.Print("Current Exp: ", p.Experience)
	p.Experience += m.ExperienceRewarded
	fmt.Println(" => ", p.Experience, "/", levelMap[p.NextLevel])
	if p.Experience > levelMap[p.NextLevel] {
		fmt.Println("Congrats! You have levelled up: Level ", p.CurrentLevel, " => ", p.NextLevel)
		fmt.Println("Damage: ", p.Damage, " => ", p.Damage+4)
		p.Damage += 4
		fmt.Println("Hitpoints: ", p.HP, " => ", p.HP+12)
		p.HP += 12
		p.CurrentLevel = p.NextLevel
		p.NextLevel++
	}
	return p
}
