// Created by Callum Gardner at 2026/09/08 10:09
// leetgo: dev
// https://leetcode.com/problems/find-the-difference-of-two-arrays/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findDifference(nums1 []int, nums2 []int) [][]int {
	const arrayOffset = 1000
	const arraySize = 2*arrayOffset + 1

	m1 := [arraySize]bool{}
	for _, v := range nums1 {
		m1[v+arrayOffset] = true
	}
	m2 := [arraySize]bool{}
	for _, v := range nums2 {
		m2[v+arrayOffset] = true
	}

	s1 := []int{}
	s2 := []int{}
	for i, v := range m1 {
		if v && !m2[i] {
			s1 = append(s1, i-arrayOffset)
		}
	}
	for i, v := range m2 {
		if v && !m1[i] {
			s2 = append(s2, i-arrayOffset)
		}
	}

	return [][]int{s1, s2}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := findDifference(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
