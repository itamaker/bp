package main

import (
	"log"
	"os"
	// "bytes"
	"database/sql"
	"fmt"
	// "os/exec"
	// "strconv"
	"bufio"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// prepare and return sqlite3 database
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

// prompt for new alias input
func promptNewAlias() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter alias name: ")
	alias, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	alias = strings.TrimSpace(alias)

	fmt.Print("Enter command literal: ")
	command, err2 := reader.ReadString('\n')
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	command = strings.TrimSpace(command)

	insertNewAlias(alias, command)
}

func insertNewAlias(alias string, command string) {
	db := prepareDB()
	insert, err1 := db.Prepare("INSERT INTO `aliases` (`alias`, `command`) VALUES (?, ?)")
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	_, err2 := insert.Exec(alias, command)
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	rows, _ := db.Query("SELECT `alias`, `command` FROM `aliases`")
	var insertedAlias string
	var insertedCmd string
	for rows.Next() {
		rows.Scan(&insertedAlias, &insertedCmd)
		fmt.Println(insertedAlias + ": " + insertedCmd)
	}
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
	operation := os.Args[1]
	if operation == "add" {
		promptNewAlias()
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
