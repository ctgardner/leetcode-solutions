// Created by Callum Gardner at 2026/09/18 15:32
// leetgo: dev
// https://leetcode.com/problems/counting-bits/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countBits(n int) (ans []int) {
	ans = make([]int, 0, n+1)
	ans = append(ans, 0)

	for i := 1; i <= n; i++ {
		ans = append(ans, ans[i>>1]+i&1)
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := countBits(n)

	fmt.Println("\noutput:", Serialize(ans))
}
