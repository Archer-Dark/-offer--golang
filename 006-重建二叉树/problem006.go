package main

import "fmt"

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

func printPreOrder(root *TreeNode){
	if root != nil {
		fmt.Printf("%d ", root.Val)
		printPreOrder(root.Left)
		printPreOrder(root.Right)
	}
}

func printInOrder(root *TreeNode){
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		printInOrder(root.Right)
	}
}

func main() {
	// example
	pre := []int{1,2,4,7,3,5,6,8}
	in := []int{4,7,2,1,5,3,6,8}

	fmt.Println("preOder: ", pre)
	fmt.Println("inOrder: ", in)

	// Reconstruct
	fmt.Println("\nReconstruct Binary Tree... \n ",)
	root := buildTree(pre, in)

	// test
	fmt.Printf("preOder from Tree reconstructed:  ")
	printPreOrder(root)
	fmt.Printf("\n")

	fmt.Printf("inOder from Tree reconstructed:  ")
	printInOrder(root)
	fmt.Printf("\n")
}