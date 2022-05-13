package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List){
	fmt.Println("倒着输出")
	for t := l.Back(); t != nil ; t = t.Prev()  {
		fmt.Println(t.Value," ")
	}
	fmt.Println()
	fmt.Println("正着输出")
	for t := l.Front(); t != nil ; t = t.Next()  {
		fmt.Println(t.Value," ")
	}
}

func main() {
	//初始化
	values := list.New()
	//基本使用
	e1 := values.PushBack("ONE")  //后面插入
	e2 := values.PushBack("TWO")
	values.PushFront("ZERO") //前面插入
	printList(values)
	values.InsertBefore("insertBefore one",e1) //指定位置插入
	values.InsertAfter("insertAfter two",e2)
	printList(values) //自定义函数打印链表
	values.PushBackList(values)//链表整个插入
	printList(values)
}
