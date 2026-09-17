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
	memo := make([]int, len(cost))
	for i := range memo {
		memo[i] = -1
	}
	memo[0], memo[1] = cost[0], cost[1]

	var minCost func(int) int
	minCost = func(i int) int {
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = min(minCost(i-1), minCost(i-2)) + cost[i]
		return memo[i]
	}
	return min(minCost(len(cost)-1), minCost(len(cost)-2))
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	cost := Deserialize[[]int](ReadLine(stdin))
	ans := minCostClimbingStairs(cost)

	fmt.Println("\noutput:", Serialize(ans))
}
