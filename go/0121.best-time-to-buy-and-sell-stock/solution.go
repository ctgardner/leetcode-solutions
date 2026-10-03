// Created by Callum Gardner at 2026/10/03 10:34
// leetgo: dev
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	maxFuturePrice := 0
	for i := len(prices) - 1; i >= 0; i-- {
		ans = max(ans, maxFuturePrice-prices[i])
		maxFuturePrice = max(maxFuturePrice, prices[i])
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prices := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfit(prices)

	fmt.Println("\noutput:", Serialize(ans))
}
