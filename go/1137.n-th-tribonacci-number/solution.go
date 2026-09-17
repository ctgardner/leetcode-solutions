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

func tribonacci(n int) int {
	window := [3]int{0, 1, 1}
	if n < 3 {
		return window[n]
	}
	for i := 3; i <= n; i++ {
		sum := window[0] + window[1] + window[2]
		window[0], window[1] = window[1], window[2]
		window[2] = sum
	}
	return window[2]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := tribonacci(n)

	fmt.Println("\noutput:", Serialize(ans))
}
