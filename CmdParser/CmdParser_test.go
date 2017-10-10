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
