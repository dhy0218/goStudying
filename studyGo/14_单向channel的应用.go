package main

import "fmt"

//只能写不能读
func producer(out chan<- int){
	for i:=0;i<10 ;i++  {
		out<- i*i
	}
	close(out)
}
//只能读不能写
func consumer(in <-chan int){
	for num := range in  {
		fmt.Println(num)
	}
}
func main() {
	ch := make(chan int)
	//生产者 写入channel
	go producer(ch)//channel传参 引用传递
	//消费者 从channel读取内容打印
	consumer(ch)
}
