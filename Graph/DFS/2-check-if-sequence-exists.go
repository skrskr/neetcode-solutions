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
 

func existPathUtil(root *TreeNode, arr []int, lenght int, index int)  bool{

	if (root == nil || index == lenght) {
		return false
	}

	if (root.Left == nil && root.Right == nil) {
		if root.Val == arr[index] && index == lenght - 1 {
			return true
		}
		return false
	}

	return ( (lenght > index) && (arr[index] == root.Val) && (existPathUtil(root.Left, arr, lenght, index+1) || existPathUtil(root.Right, arr, lenght, index+1)))
}

func exists(root *TreeNode, arr []int, length int) bool {
	if root == nil {
		return (length==0);
	}
	return existPathUtil(root, arr, len(arr), 0)
}



func main() {
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 8}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node5.Left = node6
	node5.Right = node7

	arr := []int{5, 2,4}
	fmt.Println(exists(node1, arr, len(arr)))
}