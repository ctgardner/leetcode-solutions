// Created by Callum Gardner at 2026/03/12 12:09
// leetgo: dev
// https://leetcode.com/problems/longest-palindromic-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestPalindrome(s string) string {
	longest := longestPalindromeFromCenter(s)

	for i, j := 1, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		if len(longest) >= len(s)-i {
			break
		}

		p := longestPalindromeFromCenter(s[i:])
		if len(p) > len(longest) {
			longest = p
			if len(p) == len(s[i:]) {
				continue
			}
		}

		p = longestPalindromeFromCenter(s[:j])
		if len(p) > len(longest) {
			longest = p
		}
	}
	return longest
}

func longestPalindromeFromCenter(s string) string {
	left, right := (len(s)-1)/2, len(s)/2
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return s[left+1 : right]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
