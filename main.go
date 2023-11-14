package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name   string // this is the name
	HP     int    // this is value
	Damage int    // how much damage we're dealing
}

type Monster struct {
	Name   string // this is the name
	HP     int    // this is value
	Damage int    // how much damage we're dealing
}

func NewPlayer() Player {
	return Player{
		Name:   "Unknown",
		HP:     100,
		Damage: 8,
	}
}

func NewMonster() Monster {
	return Monster{
		Name:   "Unknown",
		HP:     30,
		Damage: 4,
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter some text: ")

	// Remove newline character from input

	player1 := NewPlayer()
	player1.Name = "Brian"

	mob1 := NewMonster()
	mob1.Name = "Piglet"
	turn := 1
	fmt.Println("Oh no you encountered a ", mob1.Name, "!")
	fmt.Println("BATTLE START")
	for player1.HP > 0 && mob1.HP > 0 {

		fmt.Println("---------- TURN", turn, "----------")
		fmt.Println(player1.Name, " HP: ", player1.HP, " ", mob1.Name, " HP: ", mob1.HP)
		fmt.Print("Attack (-", player1.Damage, "): Press: [a] | Pass (0): Press [s]")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")
		if userInput == "a" {
			fmt.Println("You attack! ", mob1.Name, "'s health: ", mob1.HP, " => ", mob1.HP-player1.Damage)
			mob1.HP = mob1.HP - player1.Damage
		}
		fmt.Println(mob1.Name, " attacks!", player1.Name, "'s health: ", player1.HP, " => ", player1.HP-mob1.Damage)
		player1.HP = player1.HP - mob1.Damage
		if player1.HP <= 0 {
			fmt.Println("Oh dear! You have died.")
		}
		if mob1.HP <= 0 {
			fmt.Println("You won!")
		}
		turn++
	}

}
