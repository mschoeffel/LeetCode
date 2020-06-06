package LongestSubstringWithoutRepeatingCharacters

import (
	"fmt"
	"strconv"
)

func Main() {

	s := "abcabcbb"

	fmt.Println("Result: " + strconv.Itoa(lengthOfLongestSubstring(s)))

}

func lengthOfLongestSubstring(s string) int {

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
