// Created by Callum Gardner at 2026/10/03 11:15
// leetgo: dev
// https://leetcode.com/problems/majority-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func majorityElement(nums []int) int {
	element, tally := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if tally == 0 {
			element, tally = nums[i], 1
			continue
		}
		if element == nums[i] {
			tally++
		} else {
			tally--
		}
	}
	return element
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := majorityElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
