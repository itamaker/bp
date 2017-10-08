package main

import (
	"log"
	"os"
	// "bytes"
	"database/sql"
	"fmt"
	// "os/exec"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func prepareDB() *sql.DB {
	db, err1 := sql.Open("sqlite3", "./betaphor.sqlite3")
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	createTable, err2 := db.Prepare("CREATE TABLE IF NOT EXISTS `aliases` (`alias` TEXT PRIMARY KEY, `command` TEXT NOT NULL)")
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	_, err3 := createTable.Exec()
	if err3 != nil {
		log.Fatal(err3.Error())
		os.Exit(3)
	}
	return db
}

func init() {
	prepareDB()
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
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement1.Exec()
	statement2, _ := database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement2.Exec("Nic", "Raboy")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
