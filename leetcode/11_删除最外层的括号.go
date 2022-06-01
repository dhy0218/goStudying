package main

import "fmt"

func main() {
	//x := "我的"
	//for _,data := range x  {
	//	fmt.Printf("%c",data)
	//}
	//fmt.Println(x)
	s := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s))
}
func removeOuterParentheses(s string) string {
	ans := []rune{}
	level := 0
	for _,data := range s  {
		if data == ')' {
			level--
		}
		if level>0 {
			ans = append(ans,data)
		}
		if data == '(' {
			level++
		}
	}
	return string(ans)
}
