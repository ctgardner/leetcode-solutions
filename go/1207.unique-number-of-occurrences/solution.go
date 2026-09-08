// Created by Callum Gardner at 2026/09/08 12:11
// leetgo: dev
// https://leetcode.com/problems/unique-number-of-occurrences/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func uniqueOccurrences(arr []int) bool {
	m1 := map[int]int{}
	for _, v := range arr {
		m1[v] += 1
	}

	m2 := map[int]struct{}{}
	for _, v := range m1 {
		if _, ok := m2[v]; ok {
			return false
		}
		m2[v] = struct{}{}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	ans := uniqueOccurrences(arr)

	fmt.Println("\noutput:", Serialize(ans))
}
