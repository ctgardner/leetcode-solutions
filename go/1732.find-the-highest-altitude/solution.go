// Created by Callum Gardner at 2026/09/04 16:23
// leetgo: dev
// https://leetcode.com/problems/find-the-highest-altitude/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func largestAltitude(gain []int) int {
	altitude, maxAltitude := 0, 0
	for _, g := range gain {
		altitude += g
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}
	return maxAltitude
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	gain := Deserialize[[]int](ReadLine(stdin))
	ans := largestAltitude(gain)

	fmt.Println("\noutput:", Serialize(ans))
}
