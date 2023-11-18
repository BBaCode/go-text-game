package player

import (
	"fmt"

	"github.com/bbacode/go-text-game/mobs"
)

type PlayerLevel struct {
	ExpReq    int // exp required to get to this level
	DamageInc int // damage increase
	HPInc     int // HP increase
}

type Player struct {
	Name         string // this is the name
	TotalHP      int    // HP when full
	CurrentHP    int    // Current HP (after taking damage or not)
	Damage       int    // how much damage we're dealing
	Experience   int    // current expTotal
	CurrentLevel int    // current player level
	NextLevel    int    // next player level
	GameLevel    string // game level player is currently on
}

var levelMap = map[int]PlayerLevel{

	2: {
		ExpReq:    50,
		DamageInc: 4,
		HPInc:     12,
	},
	3: {
		ExpReq:    121,
		DamageInc: 5,
		HPInc:     13,
	},
	4: {
		ExpReq:    205,
		DamageInc: 6,
		HPInc:     14,
	},
	5: {
		ExpReq:    298,
		DamageInc: 7,
		HPInc:     15,
	},
	6: {
		ExpReq:    398,
		DamageInc: 8,
		HPInc:     16,
	},
	7: {
		ExpReq:    505,
		DamageInc: 9,
		HPInc:     17,
	},
	8: {
		ExpReq:    616,
		DamageInc: 10,
		HPInc:     18,
	},
	9: {
		ExpReq:    733,
		DamageInc: 11,
		HPInc:     19,
	},
	10: {
		ExpReq:    855,
		DamageInc: 12,
		HPInc:     20,
	},
}

func NewPlayer(name string) Player {
	return Player{
		Name:         name,
		TotalHP:      100,
		CurrentHP:    100,
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
	fmt.Println(" => ", p.Experience, "/", levelMap[p.NextLevel].ExpReq)
	if p.Experience > levelMap[p.NextLevel].ExpReq {
		fmt.Println("Congrats! You have levelled up: Level ", p.CurrentLevel, " => ", p.NextLevel)
		fmt.Println("Damage:    ", p.Damage, " => ", p.Damage+levelMap[p.NextLevel].DamageInc)
		p.Damage += levelMap[p.NextLevel].DamageInc
		fmt.Println("Hitpoints: ", p.TotalHP, " => ", p.TotalHP+levelMap[p.NextLevel].HPInc)
		p.TotalHP += levelMap[p.NextLevel].HPInc
		p.CurrentLevel = p.NextLevel
		p.NextLevel++
	}
	return p
}
