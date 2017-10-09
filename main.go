package main

import (
	"bp/betaphorDB"
	"bp/help"
	"fmt"
	"os"
)

func init() {
	betaphorDB.PrepareDB()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: no operation specified")
		fmt.Println("Usage: betaphor <operation> [arguments]")
		os.Exit(1)
	}

	operation := os.Args[1]
	switch operation {
	case "help":
		help.PrintHelpInfo()
	case "add":
		betaphorDB.PromptNewAlias()
	case "ls":
		betaphorDB.ListAliases()
	case "rm":
		if len(os.Args) != 3 {
			fmt.Println("ERROR: alias name not specified")
			fmt.Println("Usage: betaphor rm <alias>")
			os.Exit(1)
		}
		alias := os.Args[2]
		betaphorDB.RemoveAlias(alias)
	case "reset":
		betaphorDB.RemoveAllAliases()
	default:
		alias := operation
		betaphorDB.ExecAlias(alias)
	}
}
