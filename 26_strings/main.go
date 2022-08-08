package main

import (
	"fmt"
	"os"
	"skillbox/26_strings/str"
)

func main() {
	args := os.Args
	var firstFile string
	var secondFile string
	var resultFile string

	switch {
	case len(args) == 1:
		fmt.Println("Need a filename in args")
		return
	case len(args) == 2:
		firstFile = args[1]
		fmt.Println(str.ReadOneFile(firstFile))
	case len(args) == 3:
		firstFile, secondFile = args[1], args[2]
		fmt.Println(str.ReadTwoFiles(firstFile, secondFile))
	case len(args) > 3:
		firstFile, secondFile, resultFile = args[1], args[2], args[3]
		str.ConcatTwoFiles(firstFile, secondFile, resultFile)
	}
}
