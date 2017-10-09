package main

import (
	"betaphorDB"
	"os"
)

func init() {
	betaphorDB.PrepareDB()
}

// func main() {
// 	cmd := exec.Command("pwd")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", out.String())
// }

func main() {
	operation := os.Args[1]
	switch operation {
	case "add":
		betaphorDB.PromptNewAlias()
	case "ls":
		betaphorDB.ListAliases()
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
