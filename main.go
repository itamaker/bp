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

package main

import (
	"bp/add"
	"bp/db"
	"bp/help"
	"bp/ls"
	"fmt"
	"log"
	"os"
)

func init() {
	db.PrepareDB()
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Lshortfile)
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
		add.PromptNewAlias()
	case "ls":
		ls.Output()
	case "rm":
		if len(os.Args) != 3 {
			fmt.Println("ERROR: alias name not specified")
			fmt.Println("Usage: betaphor rm <alias>")
			os.Exit(1)
		}
		alias := os.Args[2]
		db.RemoveAlias(alias)
	case "rmAll":
		db.RemoveAllAliases()
	default:
		alias := operation
		db.ExecAlias(alias)
	}
}
