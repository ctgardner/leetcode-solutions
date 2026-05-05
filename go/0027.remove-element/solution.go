// Created by Callum Gardner at 2026/03/13 17:40
// leetgo: dev
// https://leetcode.com/problems/remove-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeElement(nums []int, val int) int {
	length := len(nums)

	i := 0
	for i < length {
		if nums[i] != val {
			i++
			continue
		}
		nums[i] = nums[length-1]
		length--
	}

	return length
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := removeElement(nums, val)

	fmt.Println("\noutput:", Serialize(ans))
}
