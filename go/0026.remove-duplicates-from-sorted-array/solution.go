// Created by Callum Gardner at 2026/03/13 16:23
// leetgo: dev
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeDuplicates(nums []int) int {
	length := 1 // length of the unique elements subarray. (the first element is always unique.)
	for j := 1; j < len(nums); j++ {
		if nums[length-1] != nums[j] {
			nums[length] = nums[j]
			length++
		}
	}
	return length
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := removeDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
