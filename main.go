package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bbacode/go-text-game/battles"
	"github.com/bbacode/go-text-game/mobs"
	"github.com/bbacode/go-text-game/player"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please Enter your name, Hero: ")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSuffix(playerName, "\n")

	fmt.Println(playerName, "? Huh? Don't sound like much of a hero, but it doesnt matter")

	player := player.NewPlayer(playerName)
	goblin := mobs.NewGoblin()

	p1 := battles.Battle(goblin, player)
	p2 := battles.Battle(goblin, p1)
	battles.Battle(goblin, p2)

}
