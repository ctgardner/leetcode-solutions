// Created by Callum Gardner at 2026/09/17 13:19
// leetgo: dev
// https://leetcode.com/problems/n-th-tribonacci-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

var memo = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}

func tribonacci(n int) (ans int) {
	if _, ok := memo[n]; !ok {
		memo[n] = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	}
	return memo[n]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := tribonacci(n)

	fmt.Println("\noutput:", Serialize(ans))
}
