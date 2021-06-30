package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	for k := range inorder {
		if inorder[k] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:k+1], inorder[0:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}

//计算树高
func CountDepth(root *TreeNode,height float64) float64 {
	if root == nil {
		return height
	}
	height++
	return math.Max(CountDepth(root.Left,height),CountDepth(root.Right,height))
}

func main()  {
	// example
	pre := []int{1,2,4,7,3,5,6,8}
	in := []int{4,7,2,1,5,3,6,8}

	root := buildTree(pre, in)
	fmt.Println(CountDepth(root,0))
}