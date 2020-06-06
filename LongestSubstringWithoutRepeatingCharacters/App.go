package LongestSubstringWithoutRepeatingCharacters

import (
	"fmt"
	"strconv"
)

func Main() {

	s := "abcabcbb"

	fmt.Println("ResultBruteForce: " + strconv.Itoa(lengthOfLongestSubstringBruteForce(s)))
	fmt.Println("ResultSaveStart: " + strconv.Itoa(lengthOfLongestSubstringSaveStart(s)))

}

func lengthOfLongestSubstringBruteForce(s string) int {

	runes := []rune(s)
	set := make(map[rune]bool)
	length := 0
	maxlength := 0
	start := 0
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if set[c] == false {
			set[c] = true
			length++
		} else {
			if length > maxlength {
				maxlength = length
			}
			start++
			i = start - 1
			set = make(map[rune]bool)
			length = 0
		}
	}
	if length > maxlength {
		maxlength = length
	}
	return maxlength
}

func lengthOfLongestSubstringSaveStart(s string) int {
	if len(s) == 0 {
		return 0
	}
	runes := []rune(s)
	set := make(map[rune]int)
	length := 0
	maxlength := 0
	start := 0
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if v, ok := set[c]; !ok || v < start {
			length++
		} else {
			if length > maxlength {
				maxlength = length
			}
			start = set[c] + 1
			length = i - set[c]
		}
		set[c] = i
	}
	if length > maxlength {
		maxlength = length
	}
	return maxlength
}
