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

package rm

import (
	"bp/db"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Remove(alias string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure to REMOVE `%s`? (y/N): ", alias)
	input, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	input = strings.TrimSpace(input)
	switch input {
	case "y", "Y":
		if db.Delete(alias) {
			fmt.Printf("Alias `%s` removed\n", alias)
		} else {
			log.Printf("Error occured when removing `%s`.\n", alias)
		}
	}
}

func RemoveAll() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ALL alias(es) will be removed, this operation can't be undone.\nAre you sure? (y/N): ")
	input, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1.Error())
		os.Exit(1)
	}
	input = strings.TrimSpace(input)
	switch input {
	case "y", "Y":
		if db.DeleteAll() {
			fmt.Println("All aliases removed\n")
		} else {
			log.Println("Error occured when removing all alias.\n")
		}
	}
}
