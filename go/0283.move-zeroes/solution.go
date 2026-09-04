// Created by Callum Gardner at 2026/09/04 12:13
// leetgo: dev
// https://leetcode.com/problems/move-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	moveZeroes(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
