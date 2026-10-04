// Created by Callum Gardner at 2026/10/04 17:27
// leetgo: dev
// https://leetcode.com/problems/contains-duplicate/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := containsDuplicate(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
