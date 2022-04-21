package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	defer fmt.Println("主协程调用结束")

	go func() {
		defer fmt.Println("子协程调用结束")
		for i:=0;i<2 ;i++  {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程，工作完毕"
	}()

	str := <-ch//没有数据之前一直阻塞
	fmt.Println(str)
}
