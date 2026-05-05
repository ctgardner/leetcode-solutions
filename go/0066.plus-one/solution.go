// Created by Callum Gardner at 2026/03/16 12:14
// leetgo: dev
// https://leetcode.com/problems/plus-one/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func plusOne(digits []int) []int {
	for i := len(digits); i > 0; i-- {
		if digits[i-1] != 9 {
			digits[i-1] += 1
			break
		}
		digits[i-1] = 0
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	digits := Deserialize[[]int](ReadLine(stdin))
	ans := plusOne(digits)

	fmt.Println("\noutput:", Serialize(ans))
}
