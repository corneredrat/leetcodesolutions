package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	repeatTracker := make(map[rune]int)
	foundDuplicate := false
	longestStringLen := 0
	currentStringLen := 0
	stringLen := len(s)
	i := 0
	for i < stringLen {
		j := i
		for j < stringLen && !foundDuplicate {
			if _, ok := repeatTracker[rune(s[j])]; ok {
				if longestStringLen < currentStringLen {
					longestStringLen = currentStringLen
				}
				foundDuplicate = true
			} else {
				repeatTracker[rune(s[j])] = 1
				currentStringLen += 1
				if longestStringLen < currentStringLen {
					longestStringLen = currentStringLen
				}
			}
			j += 1
		}
		currentStringLen = 0
		foundDuplicate = false
		repeatTracker = make(map[rune]int)
		i += 1
	}
	return longestStringLen
}

func main() {
	len := lengthOfLongestSubstring("pwwkew")
	fmt.Printf("Length: %v\n", len)
}
