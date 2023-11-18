package levels

import (
	"fmt"

	"github.com/bbacode/go-text-game/battles"
	"github.com/bbacode/go-text-game/mobs"
	"github.com/bbacode/go-text-game/player"
)

func Level1(p player.Player) {
	fmt.Println("Level 1: The Cursed Abyss")
	updatedPlayer := p

	// Scene 1: The Entrance to Darkness
	fmt.Println("You stand at the entrance of the Cursed Abyss.")
	fmt.Println("a) Take the well-lit path, possibly guarded by ancient guardians.")
	fmt.Println("b) Venture into the shadows, risking encounters with unknown creatures.")

	// Decision Point 1
	var pathChoice string
	fmt.Scanln(&pathChoice)

	if pathChoice == "a" {
		fmt.Println("You choose the well-lit path.")
	} else if pathChoice == "b" {
		fmt.Println("You venture into the shadows.")
	} else {
		fmt.Println("Invalid choice. Defaulting to the well-lit path.")
	}

	// Scene 2: Trials of the Cursed Guardians
	fmt.Println("As you delve deeper, cursed guardians awaken to test your worth.")

	// Battle Encounter
	if battles.Battle(mobs.NewGoblin(), p).PlayerWon == true {
		fmt.Println("You defeat the cursed guardians.")
		updatedPlayer = battles.Battle(mobs.NewGoblin(), p).UpdatedPlayer
	} else {
		fmt.Println("The cursed guardians overwhelm you. Game over.")
		return
	}

	// Scene 3: The Enigmatic Runes
	fmt.Println("A massive chamber adorned with mysterious runes awaits you.")

	// Puzzle Challenge
	// if Puzzle() {
	fmt.Println("You solve the rune puzzle and reveal the hidden path forward.")
	// } else {
	// 	fmt.Println("Incorrect choices trigger traps or summon more cursed entities. Game over.")
	// 	return
	// }

	// Scene 4: The Abyssal Depths
	fmt.Println("Navigating through treacherous terrain, you reach the abyssal depths.")

	// Story Interaction
	fmt.Println("Uncover the history of your bloodline and make choices that shape your character's perception.")

	// Scene 5: Confrontation with the Abyssal Warden (Boss Battle)
	fmt.Println("At the heart of the Cursed Abyss stands the Abyssal Warden.")

	// Boss Battle
	if battles.Battle(mobs.NewGoblin(), updatedPlayer).PlayerWon == true {
		fmt.Println("You defeat the Abyssal Warden and claim a mystical cloak.")
	} else {
		fmt.Println("The Abyssal Warden proves too powerful. Game over.")
		return
	}

	// Reward
	fmt.Println("The radiant cloak enhances your defenses, providing a glimpse of godlike power.")
	fmt.Println("With the cloak adorning your shoulders, you are ready to ascend further into the realms of your destiny.")
}
