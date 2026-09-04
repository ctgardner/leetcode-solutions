// Created by Callum Gardner at 2026/09/04 16:35
// leetgo: dev
// https://leetcode.com/problems/find-pivot-index/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	left := 0
	for i, n := range nums {
		right := sum - left - n
		if left == right {
			return i
		}
		left += n
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := pivotIndex(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
