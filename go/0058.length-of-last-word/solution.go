// Created by Callum Gardner at 2026/03/16 11:33
// leetgo: dev
// https://leetcode.com/problems/length-of-last-word/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s); i > 0; i-- {
		if s[i-1] == ' ' {
			if length > 0 {
				break
			}
			continue
		}
		length++
	}
	return length
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLastWord(s)

	fmt.Println("\noutput:", Serialize(ans))
}
