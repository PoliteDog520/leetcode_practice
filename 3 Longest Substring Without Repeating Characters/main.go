package main

import (
	"fmt"
	"log"
)

func lengthOfLongestSubstring(s string) int {
	var charMap = make(map[string]struct{})
	var longestLength int

	for i := range s {
		startChar := fmt.Sprintf("%c", s[i])
		charMap[startChar] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			var contentChar = fmt.Sprintf("%c", s[j])
			if _, ok := charMap[contentChar]; ok {
				if len(charMap) > longestLength {
					longestLength = len(charMap)
				}
				charMap = make(map[string]struct{})
				break
			} else {
				charMap[contentChar] = struct{}{}
			}
		}
	}

	// 最大長度為整組字串
	if len(charMap) > 0 && len(charMap) > longestLength {
		longestLength = len(charMap)
	}

	return longestLength
}

func main() {
	log.Println(lengthOfLongestSubstring("abcabcbb"))
	log.Println(lengthOfLongestSubstring("bbbbb"))
	log.Println(lengthOfLongestSubstring("pwwkew"))
	log.Println(lengthOfLongestSubstring("dvdf"))
	log.Println(lengthOfLongestSubstring(" "))
	log.Println(lengthOfLongestSubstring("au"))
}
