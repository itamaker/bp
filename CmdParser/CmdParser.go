package CmdParser

import (
	"regexp"
	"strings"
)

const (
	joint = "-+-"
)

func Encode(cmdStr string) string {
	// if contain option, pick out options
	// \B-{1,2}[A-Za-z-]+(\ )?[A-Za-z]+\b
	// var cmd = "carthage update --platform iOS --no-use-binaries"
	// var cmd = "ab cd --ef gh --i-j-k"
	var cmd = cmdStr
	re := regexp.MustCompile("\\B-{1,2}[A-Za-z-]+(\\ )[A-Za-z]+")
	substrings := re.FindAllString(cmd, -1)
	for i := 0; i < len(substrings); i++ {
		origin := substrings[i]
		origin = strings.Replace(origin, " ", joint, -1)
		// fmt.Println(substrings[i])
		cmd = strings.Replace(cmd, substrings[i], origin, -1)
	}
	return cmd
}

func Decode(cmdStr string) string {
	var cmd = cmdStr
	cmd = strings.Replace(cmd, joint, " ", -1)
	return cmd
}
