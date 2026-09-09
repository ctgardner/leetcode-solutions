// Created by Callum Gardner at 2026/09/09 13:15
// leetgo: dev
// https://leetcode.com/problems/leaf-similar-trees/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	stack1 := []*TreeNode{root1}
	stack2 := []*TreeNode{root2}
	var leaf1, leaf2 *TreeNode
	for len(stack1) > 0 && len(stack2) > 0 {
		leaf1, stack1 = nextLeaf(stack1)
		leaf2, stack2 = nextLeaf(stack2)
		if leaf1.Val != leaf2.Val {
			return false
		}
	}
	return len(stack1) == 0 && len(stack2) == 0
}

func nextLeaf(stack []*TreeNode) (*TreeNode, []*TreeNode) {
	var leaf *TreeNode
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			leaf = node
			break
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return leaf, stack
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root1 := Deserialize[*TreeNode](ReadLine(stdin))
	root2 := Deserialize[*TreeNode](ReadLine(stdin))
	ans := leafSimilar(root1, root2)

	fmt.Println("\noutput:", Serialize(ans))
}
