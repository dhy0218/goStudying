package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int,3)

	fmt.Println(len(ch),cap(ch))
	//新建协程
	go func() {
		for i:=0;i<10 ;i++  {
			fmt.Println("子协程",i)
			ch<- i
		}
		close(ch)
	}()

	time.Sleep(2*time.Second)
	//for i:=0;i<10 ;i++  {
	//	num := <-ch
	//	fmt.Println("父协程",num)
	//}
	for num := range ch{
		fmt.Println(num)
	}
}
