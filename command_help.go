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
}

func (command *CommandHelp) Init() {
	command.Line1 = "adventure		:	Set out on an adventure to find experience, loot, and random battles."
	command.Line2 = "friends		:	See who your friends are!"
	command.Line3 = "me		: 	Look at your stats."
	command.Line4 = "stuff		: 	See your stuff."
	command.Line5 = "shop		:	Sell your stuff, buy more stuff."
	command.Line6 = "quit		:	Quit the game :-("

}

func (command *CommandHelp) PrintHelpCommands() {
	fmt.Println()
	fmt.Println(command.Line1)
	fmt.Println(command.Line2)
	fmt.Println(command.Line3)
	fmt.Println(command.Line4)
	fmt.Println(command.Line5)
	fmt.Println(command.Line6)
	fmt.Println()
}
