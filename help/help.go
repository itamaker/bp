package help

import (
	"fmt"
)

func PrintHelpInfo() {
	fmt.Println("Usage:")
	fmt.Println("bp add:\t\tAdd new alias by following prompt.")
	fmt.Println("bp ls:\t\tList all aliases")
	fmt.Println("bp rm <alias>:\tRemove a alias.")
	fmt.Println("bp rmAll:\tRemove all aliases")
	fmt.Println("bp <alias>:\tExecute a alias.")
}
