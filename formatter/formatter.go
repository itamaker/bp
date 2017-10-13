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

package formatter

import (
	"fmt"
	// "log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/// append whitespace to string from left
func appendSpaceFromLeft(str string) string {
	maxLen := 10
	curLen := len(str)
	if curLen > maxLen {
		return str
	}
	var newStr string
	for i := 0; i < maxLen-curLen; i++ {
		newStr = newStr + " "
	}
	return newStr + str
}

func getTerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, _ := cmd.Output()
	comp := strings.Split(string(output), " ")
	widthStr := strings.TrimSpace(comp[1])
	w, _ := strconv.Atoi(widthStr)
	return w
}

/// decorate ls output
func FormatAliasMap(aliases [][]string) {
	terminalW := getTerminalSize()
	var sep string
	for i := 0; i < terminalW; i++ {
		sep = sep + "─"
	}
	fmt.Println(sep)
	for _, aliasAndCmd := range aliases {
		formattedAlias := appendSpaceFromLeft(aliasAndCmd[0])
		fmt.Println(formattedAlias + " ➜  " + aliasAndCmd[1])
	}
	fmt.Println(sep)
}
