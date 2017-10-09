package main

import (
	"bp/betaphorDB"
	"fmt"
	"os"
)

func init() {
	betaphorDB.PrepareDB()
}

// func main() {
// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: no operation specified")
		fmt.Println("Usage: betaphor <operation> [arguments]")
		os.Exit(1)
	}

	operation := os.Args[1]
	switch operation {
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
	// database, _ := sql.Open("sqlite3", "./nraboy.db")
	// statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	// statement1.Exec()
	// statement2, _ := database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	// statement2.Exec("Nic", "Raboy")
	// rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	// var id int
	// var firstname string
	// var lastname string
	// for rows.Next() {
	// 	rows.Scan(&id, &firstname, &lastname)
	// 	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	// }
}
