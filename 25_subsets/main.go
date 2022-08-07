package main

import (
	"flag"
	"fmt"
)

func main() {
	//looking for a substring in a string using runes array
	var str string
	var substr string

	flag.StringVar(&str, "str", "empty string", "set string")
	flag.StringVar(&substr, "substr", "empty substring", "set substring")
	flag.Parse()
	fmt.Println("First string:", str)
	fmt.Println("Second string:", substr)

	//translation str into runes array
	counter := 0
	strRunes := make([]rune, len(str))
	for _, r := range str {
		if r != 0 {
			strRunes[counter] = r
			counter++
		}
	}
	strRunes = strRunes[:counter]

	//translation substr into runes array
	counter = 0
	substrRunes := make([]rune, len(substr))
	for _, r := range substr {
		if r != 0 {
			substrRunes[counter] = r
			counter++
		}
	}
	substrRunes = substrRunes[:counter]

	//looking for a substring
	result := subSearch(substrRunes, strRunes, 0, 0, false)
	fmt.Println("Second string is a substring of first:", result)
}

func subSearch(substrRunes []rune, strRunes []rune, i, j int, isSubstr bool) bool {
	for ; i < len(substrRunes); i++ {
		for ; j < len(strRunes); j++ {
			if substrRunes[i] == strRunes[j] {
				i, j = i+1, j+1
				isSubstr = true
				return subSearch(substrRunes, strRunes, i, j, isSubstr)
			} else {
				isSubstr = false
				continue
			}
		}
	}
	return isSubstr
}
