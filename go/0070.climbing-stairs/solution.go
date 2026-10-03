// Created by Callum Gardner at 2026/10/03 21:02
// leetgo: dev
// https://leetcode.com/problems/climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func climbStairs(n int) (ans int) {
	if n <= 2 {
		return n
	}
	backTwo, backOne := 1, 2
	for i := 2; i < n; i++ {
		backTwo, backOne = backOne, backTwo+backOne
	}
	return backOne
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := climbStairs(n)

	fmt.Println("\noutput:", Serialize(ans))
}
