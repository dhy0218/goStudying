package main

import "fmt"

func main() {
	//传入多个参数
	fmt.Println(silceFunc("sum :%d",1,2,3))
	//slice对象做参数，必须展开
	list := []int{1,2,3}
	fmt.Println(silceFunc("test",list...))
}

func silceFunc(s string , n ...int) string {
	var x int
	for _,data := range n  {
		x += data
	}
	return fmt.Sprint(s,x)
}