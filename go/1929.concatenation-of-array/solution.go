// Created by Callum Gardner at 2026/10/04 13:15
// leetgo: dev
// https://leetcode.com/problems/concatenation-of-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getConcatenation(nums []int) (ans []int) {
	n := len(nums)
	ans = make([]int, n*2)
	for i, v := range nums {
		ans[i], ans[i+n] = v, v
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := getConcatenation(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
