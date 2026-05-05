// Created by Callum Gardner at 2026/03/15 15:46
// leetgo: dev
// https://leetcode.com/problems/search-insert-position/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		k := (i+j)/2
		if nums[k] == target {
			return k
		}

		if target < nums[k] {
			j = k
		} else {
			i = k+1
		}
	}

	return i
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchInsert(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
