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

/// decorate ls output
func FormatAliasMap(aliases map[string]string) {
	fmt.Println("┌─────")
	for alias, command := range aliases {
		formattedAlias := appendSpaceFromLeft(alias)
		fmt.Println("│" + formattedAlias + "➜ " + command)
	}
	fmt.Println("└─────")
}
