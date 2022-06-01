package main

import "fmt"

type Node struct {
	Val int
	children []*Node
}

func main() {
	a := Node{Val:1}
	b := Node{Val:2}
	c := Node{Val:3}
	d := Node{Val:4}
	e := Node{Val:5}
	f := Node{Val:6}
	a.children = []*Node{&b,&c,&d}
	c.children = []*Node{&e,&f}
	fmt.Println(levelOrder(&a))
}
func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	nodes := []*Node{root}
	for ; len(nodes)>0;  {
		temp := nodes
		nodes = []*Node{}
		x := []int{}
		for _,data := range temp   {
			nodes = append(nodes,data.children...)
			x = append(x,data.Val)
		}
		ans = append(ans,x)
	}
	return ans
}