package main

import (
	"fmt"
	"sort"
	"strings"
)

// This function checks if two strings are permutuations of one another. We first sort the strings, then compare element by element
func perm(str1 string, str2 string) bool {
	// If not same size, impossible to be permutation
	if len(str1) != len(str2) {
		return false
	}

	// Sort the strings
	s1 := strings.Split(str1, "")
	sort.Strings(s1)
	s2 := strings.Split(str2, "")
	sort.Strings(s2)

	// Iterate over str1 or str2, comparing values
	for i := range str1 {
		// Didn't match
		if s1[i] != s2[i] {
			return false
		}
	}
	// Permutations of each other
	return true
}

func main() {
	string1 := "tes  t"
	string2 := "test"
	fmt.Println(perm(string1, string2))

}
