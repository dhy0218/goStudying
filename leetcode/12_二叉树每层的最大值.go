package main

import (
	"fmt"
	"math"
)

/**
输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]
解释:
          1
         / \
        3   2
       / \   \
      5   3   9
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
	ans := largestValues(&a)
	fmt.Println(ans)
}
func largestValues(root *TreeNode) []int {
	ans := []int{}
	level := []*TreeNode{root}
	for len(level)>0 {
		temp := level
		level = []*TreeNode{}
		max := math.MinInt
		for _,data := range temp  {
			val := data.Val
			max = Max(max,val)
			if data.Left!=nil {
				level = append(level,data.Left)
			}
			if data.Right != nil {
				level = append(level,data.Right)
			}
		}
		ans = append(ans,max)
	}
	return ans
}
func Max(a int,b int)int{
	if a>b {
		return a
	}else{
		return b
	}
}