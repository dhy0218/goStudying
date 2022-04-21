package main

import (
	"fmt"
	"time"
)
//主协程退出 子协程也要退出
func main() {
	i:=0
	go func(){
		i++
		fmt.Println("子协程",i)
		time.Sleep(time.Second)

	}()
	for   {
		i++
		fmt.Println("主协程",i)
		time.Sleep(time.Second)
		if i==2 {
			break
		}
	}
}
