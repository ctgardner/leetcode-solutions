// Created by Callum Gardner at 2026/09/19 17:34
// leetgo: dev
// https://leetcode.com/problems/single-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func singleNumber(nums []int) (ans int) {
	for _, n := range nums {
		ans ^= n
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := singleNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
