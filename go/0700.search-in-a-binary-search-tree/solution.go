// Created by Callum Gardner at 2026/09/10 15:17
// leetgo: dev
// https://leetcode.com/problems/search-in-a-binary-search-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	val := Deserialize[int](ReadLine(stdin))
	ans := searchBST(root, val)

	fmt.Println("\noutput:", Serialize(ans))
}
