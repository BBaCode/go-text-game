package battles

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "github.com/bbacode/go-text-game/levels"
	"github.com/bbacode/go-text-game/mobs"
	"github.com/bbacode/go-text-game/player"
)

type BattleResult struct {
	PlayerWon     bool
	UpdatedPlayer player.Player // If the player won, this contains the updated player information
}

func Battle(mob mobs.Monster, p player.Player) BattleResult {
	reader := bufio.NewReader(os.Stdin)
	turn := 1
	updatedPlayer := p
	fmt.Println("Oh no you encountered a", mob.Name, "!")
	fmt.Println("BATTLE START")
	for p.CurrentHP > 0 && mob.HP > 0 {

		fmt.Println("---------- TURN", turn, "----------")
		fmt.Println(p.Name, " HP: ", p.CurrentHP)
		fmt.Println(mob.Name, " HP: ", mob.HP)
		fmt.Println("Attack (-", p.Damage, "):[a]")
		fmt.Println("Pass (0):[s]")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
		if userInput == "a" {
			fmt.Println("You attack! ", mob.Name, "'s health: ", mob.HP, " => ", mob.HP-p.Damage)
			mob.HP = mob.HP - p.Damage
		}
		if mob.HP <= 0 {
			fmt.Println("You won!")
			updatedPlayer = player.AddExperience(p, mob)
			return BattleResult{PlayerWon: true, UpdatedPlayer: updatedPlayer}
		} else {
			fmt.Println(mob.Name, " attacks!", p.Name, "'s health: ", p.CurrentHP, " => ", p.CurrentHP-mob.Damage)
			p.CurrentHP = p.CurrentHP - mob.Damage
		}
		if p.CurrentHP <= 0 {
			fmt.Println("Oh dear! You have died.")
			return BattleResult{PlayerWon: false, UpdatedPlayer: p}
		}

		turn++
	}
	return BattleResult{PlayerWon: true, UpdatedPlayer: updatedPlayer}
}
