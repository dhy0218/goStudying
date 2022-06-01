package main

import "fmt"

/**
输入: root = [2,1,3], p = 1

  2
 / \
1   3

输出: 2
找出后继节点
 */
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main() {
	a := TreeNode{Val:2}
	b := TreeNode{Val:1}
	c := TreeNode{Val:3}
	a.Left = &b
	a.Right = &c
	temp := inorderSuccessor(&a,&b)
	fmt.Println(temp)
}
var l []*TreeNode
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	inorder(root)
	len := len(l)
	for i:=0;i<len ;i++  {
		if p == l[i] {
			if len>=i+2 {
				return l[i+1]
			}else{
				return nil
			}
		}
	}
	return nil
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	l = append(l, root)
	inorder(root.Right)
}