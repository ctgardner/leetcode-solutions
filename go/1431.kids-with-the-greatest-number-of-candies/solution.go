// Created by Callum Gardner at 2026/09/03 14:08
// leetgo: dev
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func kidsWithCandies(candies []int, extraCandies int) (ans []bool) {
	maxCandies := slices.Max(candies)

	ans = make([]bool, len(candies))
	for i, candy := range candies {
		ans[i] = candy+extraCandies >= maxCandies
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	candies := Deserialize[[]int](ReadLine(stdin))
	extraCandies := Deserialize[int](ReadLine(stdin))
	ans := kidsWithCandies(candies, extraCandies)

	fmt.Println("\noutput:", Serialize(ans))
}
