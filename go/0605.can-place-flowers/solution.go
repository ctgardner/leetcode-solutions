// Created by Callum Gardner at 2026/09/03 14:35
// leetgo: dev
// https://leetcode.com/problems/can-place-flowers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canPlaceFlowers(flowerbed []int, n int) bool {
	suitablePlots := 0
	for i := 0; i < len(flowerbed) && suitablePlots < n; i++ {
		if flowerbed[i] == 1 {
			i++ // don't check adj plot
		} else if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			suitablePlots++
			i++ // don't check adj plot
		}
	}
	return n <= suitablePlots
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	flowerbed := Deserialize[[]int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := canPlaceFlowers(flowerbed, n)

	fmt.Println("\noutput:", Serialize(ans))
}
