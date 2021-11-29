package common

import (
	"strings"
)

// EveryChar returns a slice with every character in string
func EveryChar(s string) []string {
	return strings.Split(s, "")
}

// GetDiff returns number of differences between two strings of equal length
func GetDiff(s1, s2 string) int {
	res := 0
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			res++
		}
	}
	return res
}

// GetSimilarity returns a string with the characters that is equal in both strings
func GetSimilarity(s1, s2 string) string {
	res := ""
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] == s2[i] {
			res += string(s1[i])
		}
	}
	return res
}

// CountChar counts how many times a character exists in a string slice
func CountChar(data []string, c rune) int {
	sum := 0
	for _, r := range data {
		for _, v := range r {
			if v == c {
				sum++
			}
		}
	}
	return sum
}
