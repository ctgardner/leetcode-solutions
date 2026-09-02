// Created by Callum Gardner at 2026/09/03 07:13
// leetgo: dev
// https://leetcode.com/problems/merge-strings-alternately/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeAlternately(word1 string, word2 string) string {
	rs := make([]byte, len(word1)+len(word2))
	for i, j := 0, 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			rs[j] = word1[i]
			j++
		}
		if i < len(word2) {
			rs[j] = word2[i]
			j++
		}
	}
	return string(rs)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	word1 := Deserialize[string](ReadLine(stdin))
	word2 := Deserialize[string](ReadLine(stdin))
	ans := mergeAlternately(word1, word2)

	fmt.Println("\noutput:", Serialize(ans))
}
