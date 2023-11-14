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

func Battle(mob mobs.Monster, p player.Player) player.Player {
	reader := bufio.NewReader(os.Stdin)
	turn := 1
	fmt.Println("Oh no you encountered a ", mob.Name, "!")
	fmt.Println("BATTLE START")
	for p.HP > 0 && mob.HP > 0 {

		fmt.Println("---------- TURN", turn, "----------")
		fmt.Println(p.Name, " HP: ", p.HP, " ", mob.Name, " HP: ", mob.HP)
		fmt.Print("Attack (-", p.Damage, "):[a] | Pass (0):[s]")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
		if userInput == "a" {
			fmt.Println("You attack! ", mob.Name, "'s health: ", mob.HP, " => ", mob.HP-p.Damage)
			mob.HP = mob.HP - p.Damage
		}
		if mob.HP <= 0 {
			fmt.Println("You won!")
			winner := player.AddExperience(p, mob)
			return winner
		} else {
			fmt.Println(mob.Name, " attacks!", p.Name, "'s health: ", p.HP, " => ", p.HP-mob.Damage)
			p.HP = p.HP - mob.Damage
		}
		if p.HP <= 0 {
			fmt.Println("Oh dear! You have died.")
		}

		turn++
	}
	return p
}
