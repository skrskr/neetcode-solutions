package main

import "fmt"

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 
 func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	cur_path := fmt.Sprintf("%v",root.Val);
	
	if root.Left == nil && root.Right == nil {
		res = append(res, cur_path)
	}

	if root.Left != nil {
		dfs(root.Left, cur_path, &res)
		fmt.Println(res)
	}
	if root.Right != nil {
		dfs(root.Right, cur_path, &res)
	}
	return res
 }

 func dfs(root *TreeNode, current_path string, result *[]string) {
	current_path += ("->" + fmt.Sprintf("%v",root.Val))

	if root.Left == nil && root.Right == nil {
		*result = append(*result, current_path)
		return
	}
	if root.Left != nil  {
		dfs(root.Left, current_path, result)
	}
	if root.Right != nil {
		dfs(root.Right, current_path, result)
	}

 }

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 5}
	fmt.Println(node1.Val)

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4

	fmt.Println(binaryTreePaths(node1))
}