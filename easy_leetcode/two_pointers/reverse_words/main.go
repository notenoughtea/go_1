package main

import (
	"fmt"
	"strings"
)

// Given a string s,
// reverse the order of characters in each word
// within a sentence while still preserving whitespace and initial word order.

// Example 1:

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Example 2:

// Input: s = "Mr Ding"
// Output: "rM gniD"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		runes := []rune(w)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
