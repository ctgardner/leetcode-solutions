// Created by Callum Gardner at 2026/03/13 08:33
// leetgo: dev
// https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestCommonPrefix(strs []string) string {
	lcp := strs[0]
	for _, s := range strs[1:] {
		i := 0
		for i < len(lcp) && i < len(s) {
			if lcp[i] != s[i] {
				break
			}
			i++
		}

		if i < len(lcp) {
			lcp = lcp[:i]
		}

		if len(lcp) == 0 {
			break
		}
	}
	return lcp
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := longestCommonPrefix(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
