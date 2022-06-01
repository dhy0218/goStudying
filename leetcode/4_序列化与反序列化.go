package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	a := TreeNode{Val:1}
	b := TreeNode{Val:2}
	c := TreeNode{Val:3}
	a.Left = &b
	a.Right = &c
	temp := Codec{}.serialize(&a)
	fmt.Println(temp)
}

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
}

type Codec struct{}

func (Codec) serialize(root *TreeNode) string {
	arr := []string{}
	postOrder(root,&arr)
	return strings.Join(arr, " ")
}

func (Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	return construct(math.MinInt32, math.MaxInt32,&arr)
}
func postOrder(node *TreeNode,arr *[]string) {
	if node == nil {
		return
	}
	postOrder(node.Left,arr)
	postOrder(node.Right,arr)
	*arr = append(*arr, strconv.Itoa(node.Val))
}
func construct(lower, upper int,arr *[]string) *TreeNode {
	if len(*arr) == 0 {
		return nil
	}
	lens := len(*arr)
	val, _ := strconv.Atoi((*arr)[lens-1])
	if val < lower || val > upper {
		return nil
	}
	*arr = (*arr)[:len(*arr)-1]
	return &TreeNode{Val: val, Right: construct(val, upper,arr), Left: construct(lower, val,arr)}
}

