package main

import (
	"flag"
	"fmt"
)

func main() {
	//looking for a substring in a string (using runes array)
	var str string
	var substr string

	flag.StringVar(&str, "str", "Программирование - это просто!", "set string")
	flag.StringVar(&substr, "substr", "мир", "set substring")
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

func subSearch(substring []rune, string []rune, i, j int, isSubstring bool) bool {
	if len(substring[i:]) > len(string[j:]) {
		return false
	}
	for ; i < len(substring); i++ {
		for ; j < len(string); j++ {
			if substring[i] == string[j] {
				i, j = i+1, j+1
				isSubstring = true
				return subSearch(substring, string, i, j, isSubstring)
			} else {
				isSubstring = false
				if i == 0 {
					continue
				} else {
					return subSearch(substring, string, 0, j, isSubstring)
				}
			}
		}
	}
	return isSubstring
}
