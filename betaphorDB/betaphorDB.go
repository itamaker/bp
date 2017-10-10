//bp
//
//Copyright (c) 2017 Ke Yang
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package betaphorDB

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	shellwords "github.com/mattn/go-shellwords"
	_ "github.com/mattn/go-sqlite3"
)

// prepare and return sqlite3 database
func PrepareDB() *sql.DB {
	usr, err1 := user.Current()
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	db, err2 := sql.Open("sqlite3", usr.HomeDir+"/.betaphor")
	if err2 != nil {
		log.Fatal(err2.Error())
		os.Exit(2)
	}
	createTable, err3 := db.Prepare("CREATE TABLE IF NOT EXISTS `aliases` (`alias` TEXT PRIMARY KEY, `command` TEXT NOT NULL)")
	if err3 != nil {
		log.Fatal(err3.Error())
		os.Exit(3)
	}
	_, err4 := createTable.Exec()
	if err4 != nil {
		log.Fatal(err4.Error())
		os.Exit(4)
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

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdout)

	for in.Scan() {
		fmt.Println(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
}
