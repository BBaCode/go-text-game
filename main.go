package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bbacode/go-text-game/levels"
	"github.com/bbacode/go-text-game/player"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please Enter your name, Hero: ")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSuffix(playerName, "\n")

	fmt.Println(playerName, "? Huh? Don't sound like much of a hero, but it doesnt matter")

	player := player.NewPlayer(playerName)

	levels.Level1(player)

	// p1 := battles.Battle(goblin, player)
	// p2 := battles.Battle(goblin, p1.UpdatedPlayer)
	// battles.Battle(goblin, p2.UpdatedPlayer)

}
