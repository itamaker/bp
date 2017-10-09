package betaphorDB

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	// "strconv"
	"bufio"
	// "regexp"
	"strings"

	shellwords "github.com/mattn/go-shellwords"
	_ "github.com/mattn/go-sqlite3"
)

// prepare and return sqlite3 database
func PrepareDB() *sql.DB {
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
func PromptNewAlias() {
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

	InsertNewAlias(alias, command)
}

// insert new alias
func InsertNewAlias(alias string, command string) {
	db := PrepareDB()
	insert, err1 := db.Prepare(
		"INSERT OR REPLACE INTO `aliases` (`alias`, `command`) " +
			"VALUES (?, ?)")
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	_, err2 := insert.Exec(alias, command)
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	fmt.Println("New alias: command added")
	fmt.Println(alias + ": " + command)
}

// list all alias
func ListAliases() {
	db := PrepareDB()
	rows, _ := db.Query("SELECT `alias`, `command` FROM `aliases`")
	var alias string
	var command string
	for rows.Next() {
		rows.Scan(&alias, &command)
		fmt.Println(alias + ": " + command)
	}
}

// remove alias
func RemoveAlias(alias string) {
	db := PrepareDB()
	delete, err1 := db.Prepare("DELETE FROM `aliases` WHERE `alias` = ?")
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	_, err2 := delete.Exec(alias)
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	fmt.Printf("Alias `%s` removed\n", alias)
}

func RemoveAllAliases() {
	db := PrepareDB()
	deleteAll, err1 := db.Prepare("DELETE FROM `aliases`")
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	_, err2 := deleteAll.Exec()
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	fmt.Println("All aliases removed\n")
}

func commandWithAlias(alias string) string {
	db := PrepareDB()
	row, err1 := db.Query("SELECT `command` FROM `aliases` WHERE `alias` = ?", alias)
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	var command string
	for row.Next() {
		row.Scan(&command)
	}
	return command
}

// execute command
func ExecAlias(alias string) {
	command := commandWithAlias(alias)
	args, _ := shellwords.Parse(command)
	cmd := exec.Command(args[0], args[1:]...)
	// if isUrl, _ := regexp.MatchString("^(https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$",
	// 	command); isUrl {
	// 	cmd = exec.Command("open", command)
	// } else if isAppPath, _ := regexp.MatchString("^/.+\\.app$", command); isAppPath {
	// 	cmd = exec.Command("open", command)
	// } else {
	// 	cmd = exec.Command("osascript", "-e", "quit app \"ShadowsocksX\"")
	// 	// cmd = exec.Command(command)
	// }

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}
