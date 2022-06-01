package main

import "fmt"

/**
root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15
返回最深一层的值
 */
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main() {
	a := TreeNode{Val:1}
	b := TreeNode{Val:2}
	c := TreeNode{Val:3}
	a.Left = &b
	a.Right = &c
	ans := deepestLeavesSum(&a)
	fmt.Println(ans)
}
func deepestLeavesSum(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	level := []*TreeNode{root}
	temp := [][]int{}
	for len(level)>0 {
		t := level
		level = []*TreeNode{}
		tempLevel := []int{}
		for _,data := range t  {
			tempLevel = append(tempLevel,data.Val)
			if data.Left!=nil {
				level = append(level,data.Left)
			}
			if data.Right!=nil {
				level = append(level,data.Right)
			}
		}
		temp = append(temp,tempLevel)
	}
	for _,data := range temp[len(temp)-1] {
		ans += data
	}
	return ans
}