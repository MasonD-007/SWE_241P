package main

import "strings"

func main() {

	s1 := "ab"
	s2 := "eidbaooo"

	result := containsPermutation(s1, s2)
	if result {
		println("s2 contains a permutation of s1")
	} else {
		println("s2 does not contain a permutation of s1")
	}
}

func containsPermutation(s1 string, s2 string) bool {
	return permute([]byte(s1), 0, s2)
}

func permute(chars []byte, k int, s2 string) bool {
	if k == len(chars) {
		return strings.Contains(s2, string(chars))
	}
	for i := k; i < len(chars); i++ {
		chars[k], chars[i] = chars[i], chars[k]
		if permute(chars, k+1, s2) {
			return true
		}
		chars[k], chars[i] = chars[i], chars[k]
	}
	return false
}
