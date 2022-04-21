package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer1(s string){
	for _,data := range s{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
}
func person11(){
	printer1("hello")
	ch <- 666//给管道写数据，发送
}
func person21(){
	<-ch //从管道取数据，接收如果通道没有数据它就会阻塞
	printer1("world")
}
func main(){

	go person11()
	go person21()
	for {

	}

}
