// Created by Callum Gardner at 2026/03/12 09:43
// leetgo: dev
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) int {
	maxLength := 0

	lastIndex := make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		next := s[j]
		if l, ok := lastIndex[next]; ok && l >= i {
			i = l + 1
		}
		lastIndex[next] = j

		if j-i+1 > maxLength {
			maxLength = j - i + 1
		}
	}

	return maxLength
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
