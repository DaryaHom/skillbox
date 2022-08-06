package sort

import (
	"errors"
	"strings"
)

//Separate - returns two arrays: one of even numbers, the second of odd ones
func Separate(array []int) (even, odd []int) {
	for _, v := range array {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return
}

//ParseTest - I really can't explain what it's doing on...
func ParseTest(sentences []string, chars []rune) ([][]int, error) {
	result := make([][]int, len(chars))
	//checking for the presence of values in sentences & chars
	if sentences == nil || len(sentences) == 0 || chars == nil || len(chars) == 0 {
		err := errors.New("The array is empty")
		return result, err
	}
	for i := 0; i < len(chars); i++ {
		charsElem := strings.ToLower(string(chars[i]))
		for _, sentence := range sentences {
			//checking for the presence of values in array's cell
			if sentence == "" {
				continue
			}
			//getting the last word of sentence
			arrayOfWords := strings.Split(sentence, " ")
			lastWord := strings.ToLower(arrayOfWords[len(arrayOfWords)-1])

			index := strings.Index(lastWord, charsElem)
			if index >= 0 {
				index = len(sentence) - 1 - (len(lastWord) - 1 - index)
			}
			result[i] = append(result[i], index)
		}
	}
	return result, nil
}
