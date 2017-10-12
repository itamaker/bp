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

package CmdParser

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestEncode(t *testing.T) {
	var cmdStr = "a b --c d --e-f --g h"
	t.Log("Given the need to test raw command encoding.")
	{
		t.Logf("When encoding command: %s", cmdStr)
		{
			encoded := Encode(cmdStr)
			if encoded == "a b --c"+joint+"d --e-f --g"+joint+"h" {
				t.Logf("\tThe result shall be: %s %v", encoded, checkMark)
			} else {
				t.Fatalf("\tIncorrect format: %s %v", encoded, ballotX)
			}
		}
	}
}

func TestDecode(t *testing.T) {
	var cmdStr = "a b --c" + joint + "d --e-f --g" + joint + "h"
	t.Log("Given the need to test raw command decoding.")
	{
		t.Logf("When decoding command: %s", cmdStr)
		{
			decoded := Decode(cmdStr)
			if decoded == "a b --c d --e-f --g h" {
				t.Logf("\tThe result shall be: %s %v", decoded, checkMark)
			} else {
				t.Fatalf("\tIncorrect format: %s %v", decoded, ballotX)
			}
		}
	}
}
