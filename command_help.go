package main

import (
	"fmt"
)

type CommandHelp struct {
	Line1 string
	Line2 string
	Line3 string
	Line4 string
	Line5 string
	Line6 string
	Line7 string
}

func (command *CommandHelp) Init() {
	command.Line1 = "adventure\t: Set out on an adventure to find experience, loot, and random battles."
	command.Line2 = "friends\t: See who your friends are!"
	command.Line3 = "me\t: Look at your stats."
	command.Line4 = "stuff\t: See your stuff."
	command.Line5 = "shop\t: Sell your stuff, buy more stuff."
	command.Line6 = "quit or q\t: Quit the game :-("
	command.Line7 = "help\t: Get command help."

}

func (command *CommandHelp) PrintHelpCommands() {
	fmt.Println()
	fmt.Println("Game Commands")
	fmt.Println("-------------")
	fmt.Println(command.Line1)
	fmt.Println(command.Line2)
	fmt.Println(command.Line3)
	fmt.Println(command.Line4)
	fmt.Println(command.Line5)
	fmt.Println(command.Line6)
	fmt.Println(command.Line7)
	fmt.Println()
}
