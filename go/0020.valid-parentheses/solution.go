// Created by Callum Gardner at 2026/03/13 09:35
// leetgo: dev
// https://leetcode.com/problems/valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

var complement = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)
	for _, r := range s {
		if c, ok := complement[r]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false // no corresponding opening bracket
			}
			stack = stack[:len(stack)-1] // pop stack
		} else {
			stack = append(stack, r) // push stack
		}
	}
	return len(stack) == 0 // stack is empty if every opening bracket has a closing bracket
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isValid(s)

	fmt.Println("\noutput:", Serialize(ans))
}
