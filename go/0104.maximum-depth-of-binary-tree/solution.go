// Created by Callum Gardner at 2026/09/09 10:14
// leetgo: dev
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type pair struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depthMax := 0
	stack := []pair{{root, 1}}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		depthMax = max(p.depth, depthMax)
		if p.node.Right != nil {
			stack = append(stack, pair{p.node.Right, p.depth + 1})
		}
		if p.node.Left != nil {
			stack = append(stack, pair{p.node.Left, p.depth + 1})
		}
	}
	return depthMax
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := maxDepth(root)

	fmt.Println("\noutput:", Serialize(ans))
}
