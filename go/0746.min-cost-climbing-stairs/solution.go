// Created by Callum Gardner at 2026/09/17 14:18
// leetgo: dev
// https://leetcode.com/problems/min-cost-climbing-stairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minCostClimbingStairs(cost []int) int {
	prev2, prev1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		prev2, prev1 = prev1, min(prev2, prev1)+cost[i]
	}
	return min(prev2, prev1)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	cost := Deserialize[[]int](ReadLine(stdin))
	ans := minCostClimbingStairs(cost)

	fmt.Println("\noutput:", Serialize(ans))
}
