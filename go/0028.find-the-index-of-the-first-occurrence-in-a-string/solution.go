// Created by Callum Gardner at 2026/03/13 17:58
// leetgo: dev
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func strStr(haystack string, needle string) int {
	for i := range haystack {
		if i+len(needle) > len(haystack) {
			break
		}

		found := true
		for j := range needle {
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}

		if found {
			return i
		}
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	haystack := Deserialize[string](ReadLine(stdin))
	needle := Deserialize[string](ReadLine(stdin))
	ans := strStr(haystack, needle)

	fmt.Println("\noutput:", Serialize(ans))
}
